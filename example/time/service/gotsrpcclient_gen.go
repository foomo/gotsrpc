// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	go_context "context"
	go_net_http "net/http"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	pkg_errors "github.com/pkg/errors"
)

type ServiceGoTSRPCClient interface {
	Time(ctx go_context.Context, v time.Time) (retTime_0 time.Time, clientErr error)
	TimeStruct(ctx go_context.Context, v TimeStruct) (retTimeStruct_0 TimeStruct, clientErr error)
}

type HTTPServiceGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultServiceGoTSRPCClient(url string) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClient(url, "/service")
}

func NewServiceGoTSRPCClient(url string, endpoint string) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewServiceGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPServiceGoTSRPCClient {
	return &HTTPServiceGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPServiceGoTSRPCClient) Time(ctx go_context.Context, v time.Time) (retTime_0 time.Time, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retTime_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Time", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Time")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) TimeStruct(ctx go_context.Context, v TimeStruct) (retTimeStruct_0 TimeStruct, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retTimeStruct_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "TimeStruct", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy TimeStruct")
	}
	return
}
