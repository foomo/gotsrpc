package gotsrpc

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type bufferedClient struct {
	client  *http.Client
	handle  *transportHandle
	headers http.Header
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
func (c *bufferedClient) Call(ctx context.Context, url string, endpoint string, method string, args []any, reply []any) error {
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
		err := enc.Encode(args)
		c.handle.putEncoder(enc)

		if err != nil {
			return NewClientError(errors.Wrap(err, "failed to encode arguments"))
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

	resp, errDo := c.client.Do(request)
	if errDo != nil {
		return NewClientError(errors.Wrap(errDo, "failed to send request"))
	}
	defer resp.Body.Close()

	buf := getBuffer()
	defer putBuffer(buf)

	if _, err := io.Copy(buf, resp.Body); err != nil {
		return NewClientError(errors.Wrap(err, "failed to read response body"))
	}

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
		if value, err := clientHandle.beforeDecodeReply(reply, errorIndices); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call beforeDecodeReply hook"))
		} else {
			wrappedReply = value
		}
	}

	dec := clientHandle.getDecoder(buf)
	err := dec.Decode(wrappedReply)
	clientHandle.putDecoder(dec)

	if err != nil {
		return NewClientError(errors.Wrap(err, "failed to decode response"))
	}

	// replace error
	if clientHandle.afterDecodeReply != nil {
		if err := clientHandle.afterDecodeReply(&reply, wrappedReply, errorIndices); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call afterDecodeReply hook"))
		}
	}

	return nil
}
