// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc/v2"
)

const (
	ServiceGoTSRPCProxyHello = "Hello"
)

type ServiceGoTSRPCProxy struct {
	EndPoint string
	service  Service
}

func NewDefaultServiceGoTSRPCProxy(service Service) *ServiceGoTSRPCProxy {
	return NewServiceGoTSRPCProxy(service, "/service")
}

func NewServiceGoTSRPCProxy(service Service, endpoint string) *ServiceGoTSRPCProxy {
	return &ServiceGoTSRPCProxy{
		EndPoint: endpoint,
		service:  service,
	}
}

// ServeHTTP exposes your service
func (p *ServiceGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	} else if r.Method != http.MethodPost {
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body

	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats, _ := gotsrpc.GetStatsForRequest(r)
	callStats.Func = funcName
	callStats.Package = "github.com/foomo/gotsrpc/v2/example/monitor/service"
	callStats.Service = "Service"
	switch funcName {
	case ServiceGoTSRPCProxyHello:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_v string
		)
		args = []interface{}{&arg_v}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		helloRet := p.service.Hello(arg_v)
		callStats.Execution = time.Since(executionStart)
		rets = []interface{}{helloRet}
		if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
			gotsrpc.ErrorCouldNotReply(w)
			return
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	default:
		gotsrpc.ClearStats(r)
		gotsrpc.ErrorFuncNotFound(w)
	}
}
