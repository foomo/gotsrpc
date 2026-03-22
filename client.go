package gotsrpc

import (
	"bytes"
	"context"
	"net/http"

	"github.com/pkg/errors"
)

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

func isErrorPtr(v any) bool {
	_, ok := v.(*error)
	return ok
}
