package gotsrpc

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/foomo/gotsrpc/v3/semconv/httpconv"
	"github.com/pkg/errors"
)

type bufferedClient struct {
	client  *http.Client
	handle  *transportHandle
	headers http.Header
	instr   *httpconv.Client
}

func (c *bufferedClient) SetDefaultHeaders(headers http.Header) {
	c.headers = headers
}

func (c *bufferedClient) SetClientEncoding(encoding ClientEncoding) {
	c.handle = getHandleForEncoding(encoding)
}

func (c *bufferedClient) SetTransportHttpClient(client *http.Client) {
	c.client = client
}

// Call calls a method on the remote service
func (c *bufferedClient) Call(ctx context.Context, url string, endpoint string, method string, args []any, reply []any) (err error) {
	ctx, call := c.instr.Start(ctx, method)

	defer func() {
		if err != nil {
			call.RecordError(err, 0)
		}

		call.End()
	}()

	var errorIndices []int

	for i, v := range reply {
		if isErrorPtr(v) {
			errorIndices = append(errorIndices, i)
		}
	}
	// Marshal args
	var b *bytes.Buffer
	if len(args) > 0 {
		b = getBuffer()
		defer putBuffer(b)

		enc := c.handle.getEncoder(b)
		encErr := enc.Encode(args)
		c.handle.putEncoder(enc)

		if encErr != nil {
			return NewClientError(errors.Wrap(encErr, "failed to encode arguments"))
		}
	}

	// Create post url
	postURL := url + endpoint + "/" + method

	// Create request
	var headers http.Header
	if c.headers != nil {
		headers = c.headers.Clone()
	}

	request, errRequest := newRequest(ctx, postURL, c.handle.contentType, b, headers)
	if errRequest != nil {
		return NewClientError(errors.Wrap(errRequest, "failed to create request"))
	}

	// Propagate trace context (no-op with the default propagator; if an
	// otelhttp transport wraps the client it re-injects afterwards).
	call.Inject(request.Header)

	if b != nil {
		call.RecordRequestSize(b.Len())
	}

	resp, errDo := c.client.Do(request)
	if errDo != nil {
		return NewClientError(errors.Wrap(errDo, "failed to send request"))
	}
	defer resp.Body.Close()

	buf := getBuffer()
	defer putBuffer(buf)

	if _, copyErr := io.Copy(buf, resp.Body); copyErr != nil {
		return NewClientError(errors.Wrap(copyErr, "failed to read response body"))
	}

	call.SetStatus(resp.StatusCode)
	call.RecordResponseSize(buf.Len())

	// Check status
	if resp.StatusCode != http.StatusOK {
		return NewClientError(NewHTTPError(buf.String(), resp.StatusCode))
	}

	clientHandle := c.handle
	if ct := resp.Header.Get("Content-Type"); ct != "" && ct != c.handle.contentType {
		clientHandle = getHandlerForContentType(ct)
	}

	wrappedReply := reply
	if clientHandle.beforeDecodeReply != nil {
		if value, hookErr := clientHandle.beforeDecodeReply(reply, errorIndices); hookErr != nil {
			return NewClientError(errors.Wrap(hookErr, "failed to call beforeDecodeReply hook"))
		} else {
			wrappedReply = value
		}
	}

	dec := clientHandle.getDecoder(buf)
	decErr := dec.Decode(wrappedReply)
	clientHandle.putDecoder(dec)

	if decErr != nil {
		return NewClientError(errors.Wrap(decErr, "failed to decode response"))
	}

	// replace error
	if clientHandle.afterDecodeReply != nil {
		if hookErr := clientHandle.afterDecodeReply(&reply, wrappedReply, errorIndices); hookErr != nil {
			return NewClientError(errors.Wrap(hookErr, "failed to call afterDecodeReply hook"))
		}
	}

	return nil
}
