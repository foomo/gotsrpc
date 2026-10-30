package gotsrpc

import (
	"bytes"
	"context"
	"net/http"

	"github.com/foomo/gotsrpc/v3/semconv/httpconv"
	"github.com/pkg/errors"
)

var _ Client = &bufferedClient{}

// Client calls methods of a remote gotsrpc service over HTTP. Generated Go
// clients wrap a Client. Use [NewClient] or [NewClientWithHttpClient] to
// create one.
//
// The setters are not safe for concurrent use with Call; configure the client
// before sharing it.
type Client interface {
	// Call posts args to url+endpoint+"/"+method and decodes the response into
	// reply, whose elements must be pointers. Error values in reply are
	// restored from their [Error] wire form. Transport and encoding failures are
	// returned as [ClientError].
	Call(ctx context.Context, url string, endpoint string, method string, args []any, reply []any) (err error)
	// SetClientEncoding selects the request encoding. The default is
	// [EncodingMsgpack].
	SetClientEncoding(encoding ClientEncoding)
	// SetTransportHttpClient replaces the underlying HTTP client.
	SetTransportHttpClient(client *http.Client)
	// SetDefaultHeaders sets headers that are sent with every call. Call clones
	// them per request.
	SetDefaultHeaders(headers http.Header)
}

// NewClient returns a [Client] that uses msgpack encoding and an HTTP client
// from the factory set with [SetDefaultHttpClientFactory].
func NewClient(opts ...Option) Client {
	return &bufferedClient{
		client: defaultHttpFactory(),
		handle: getHandleForEncoding(EncodingMsgpack),
		instr:  httpconv.NewClient(opts...),
	}
}

// NewClientWithHttpClient returns a [Client] that uses msgpack encoding and
// sends requests with client. A nil client falls back to the default factory.
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
