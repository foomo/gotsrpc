// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	go_context "context"
	go_net_http "net/http"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	pkg_errors "github.com/pkg/errors"
)

type ServiceGoTSRPCClient interface {
	Hello(ctx go_context.Context, v string) (retHello_0 string, clientErr error)
}

type HTTPServiceGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

const DefaultServiceGoTSRPCClientPath = "/service"

func NewDefaultServiceGoTSRPCClient(url string, options ...gotsrpc.ClientOption) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClientWithOptions(url, DefaultServiceGoTSRPCClientPath, options...)
}

// Deprecated: Use NewServiceGoTSRPCClientWithOptions instead
func NewServiceGoTSRPCClient(url string, endpoint string) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClientWithOptions(url, endpoint)
}

// Deprecated: Use NewServiceGoTSRPCClientWithOptions instead
func NewServiceGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClientWithOptions(url, endpoint, gotsrpc.WithHTTPClient(client))
}

func NewServiceGoTSRPCClientWithOptions(url string, endpoint string, options ...gotsrpc.ClientOption) *HTTPServiceGoTSRPCClient {
	return &HTTPServiceGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewBufferedClient(options...),
	}
}

func (tsc *HTTPServiceGoTSRPCClient) Hello(ctx go_context.Context, v string) (retHello_0 string, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retHello_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Hello", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Hello")
	}
	return
}
