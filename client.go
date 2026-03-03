package gotsrpc

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	HeaderServiceToService = "X-Foomo-S2s"
)

// ClientTransport to use for calls
// var ClientTransport = &http.Transport{}

var _ Client = &bufferedClient{}

type Client interface {
	Call(ctx context.Context, url string, endpoint string, method string, args []interface{}, reply []interface{}, lastIsError bool) (err error)
	SetClientEncoding(encoding ClientEncoding)
	SetTransportHttpClient(client *http.Client)
	SetDefaultHeaders(headers http.Header)
}

func NewClient() Client {
	return &bufferedClient{client: defaultHttpFactory(), handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
}

func NewClientWithHttpClient(client *http.Client) Client { //nolint:staticcheck
	if client != nil {
		return &bufferedClient{client: client, handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
	} else {
		return &bufferedClient{client: defaultHttpFactory(), handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
	}
}

func newRequest(ctx context.Context, url string, contentType string, buffer *bytes.Buffer, headers http.Header) (r *http.Request, err error) {
	if buffer == nil {
		buffer = &bytes.Buffer{}
	}
	request, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, url, buffer)
	if errRequest != nil {
		return nil, errors.Wrap(errRequest, "could not create a request")
	}
	if len(headers) > 0 {
		request.Header = headers
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set(HeaderServiceToService, "true")

	return request, nil
}

type bufferedClient struct {
	client  *http.Client
	handle  *clientHandle
	headers http.Header
}

func (c *bufferedClient) SetDefaultHeaders(headers http.Header) {
	c.headers = headers
}

func (c *bufferedClient) SetClientEncoding(encoding ClientEncoding) {
	c.handle = getHandleForEncoding(encoding)
}

func (c *bufferedClient) SetTransportHttpClient(client *http.Client) { //nolint:staticcheck
	c.client = client
}

// Call calls a method on the remote service
func (c *bufferedClient) Call(ctx context.Context, url string, endpoint string, method string, args []any, reply []any, lastIsError bool) error {
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
	request, errRequest := newRequest(ctx, postURL, c.handle.contentType, b, c.headers.Clone())
	if errRequest != nil {
		return NewClientError(errors.Wrap(errRequest, "failed to create request"))
	}

	resp, errDo := c.client.Do(request) //nolint:gosec // G704 - URL is constructed from trusted service configuration, not user input
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

	clientHandle := getHandlerForContentType(resp.Header.Get("Content-Type"))

	wrappedReply := reply
	if clientHandle.beforeDecodeReply != nil {
		if value, err := clientHandle.beforeDecodeReply(reply, lastIsError); err != nil {
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
		if err := clientHandle.afterDecodeReply(&reply, wrappedReply, lastIsError); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call afterDecodeReply hook"))
		}
	}

	return nil
}
