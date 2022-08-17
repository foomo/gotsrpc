package gotsrpc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

const (
	HeaderServiceToService = "X-Foomo-S2S"
)

// ClientTransport to use for calls
// var ClientTransport = &http.Transport{}

var _ Client = &bufferedClient{}

type Client interface {
	Call(ctx context.Context, url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error)
	SetClientEncoding(encoding ClientEncoding)
	SetTransportHttpClient(client *http.Client)
	SetDefaultHeaders(headers http.Header)
}

func NewClient() Client {
	return &bufferedClient{client: defaultHttpFactory(), handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
}

func NewClientWithHttpClient(client *http.Client) Client { //nolint:stylecheck
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

func (c *bufferedClient) SetTransportHttpClient(client *http.Client) { //nolint:stylecheck
	c.client = client
}

// Call calls a method on the remove service
func (c *bufferedClient) Call(ctx context.Context, url string, endpoint string, method string, args []interface{}, reply []interface{}) error {
	// Marshall args
	b := new(bytes.Buffer)

	// If no arguments are set, remove
	if len(args) > 0 {
		if err := codec.NewEncoder(b, c.handle.handle).Encode(args); err != nil {
			return NewClientError(errors.Wrap(err, "failed to encode arguments"))
		}
	}

	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)

	// Create request
	request, errRequest := newRequest(ctx, postURL, c.handle.contentType, b, c.headers.Clone())
	if errRequest != nil {
		return NewClientError(errors.Wrap(errRequest, "failed to create request"))
	}

	resp, errDo := c.client.Do(request)
	if errDo != nil {
		return NewClientError(errors.Wrap(errDo, "failed to send request"))
	}
	defer resp.Body.Close()

	// Check status
	if resp.StatusCode != http.StatusOK {
		var msg string
		if value, err := io.ReadAll(resp.Body); err != nil {
			msg = "failed to read response body: " + err.Error()
		} else {
			msg = string(value)
		}
		return NewClientError(NewHTTPError(msg, resp.StatusCode))
	}

	clientHandle := getHandlerForContentType(resp.Header.Get("Content-Type"))

	wrappedReply := reply
	if clientHandle.beforeDecodeReply != nil {
		if value, err := clientHandle.beforeDecodeReply(reply); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call beforeDecodeReply hook"))
		} else {
			wrappedReply = value
		}
	}

	if err := codec.NewDecoder(resp.Body, clientHandle.handle).Decode(wrappedReply); err != nil {
		return NewClientError(errors.Wrap(err, "failed to decode response"))
	}

	// replace error
	if clientHandle.afterDecodeReply != nil {
		if err := clientHandle.afterDecodeReply(&reply, wrappedReply); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call afterDecodeReply hook"))
		}
	}

	return nil
}
