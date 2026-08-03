package gotsrpc

import (
	"bytes"
	"context"
	"net/http"

	"github.com/foomo/gotsrpc/v3/semconv/httpconv"
	"github.com/pkg/errors"
)

var _ Client = &bufferedClient{}

type Client interface {
	Call(ctx context.Context, url string, endpoint string, method string, args []any, reply []any) (err error)
	SetClientEncoding(encoding ClientEncoding)
	SetTransportHttpClient(client *http.Client)
	SetDefaultHeaders(headers http.Header)
}

func NewClient(opts ...Option) Client {
	return &bufferedClient{
		client: defaultHttpFactory(),
		handle: getHandleForEncoding(EncodingMsgpack),
		instr:  httpconv.NewClient(opts...),
	}
}

func NewClientWithHttpClient(client *http.Client, opts ...Option) Client {
	if client == nil {
		client = defaultHttpFactory()
	}

	return &bufferedClient{
		client: client,
		handle: getHandleForEncoding(EncodingMsgpack),
		instr:  httpconv.NewClient(opts...),
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
