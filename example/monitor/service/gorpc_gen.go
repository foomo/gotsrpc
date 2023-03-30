// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	tls "crypto/tls"
	gob "encoding/gob"
	fmt "fmt"
	reflect "reflect"
	strings "strings"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	gorpc "github.com/valyala/gorpc"
)

type (
	ServiceGoRPCProxy struct {
		server           *gorpc.Server
		service          Service
		callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
	}

	ServiceHelloRequest struct {
		V string
	}
	ServiceHelloResponse struct {
		RetHello_0 string
	}
)

func init() {
	gob.Register(ServiceHelloRequest{})
	gob.Register(ServiceHelloResponse{})
}

func NewServiceGoRPCProxy(addr string, service Service, tlsConfig *tls.Config) *ServiceGoRPCProxy {
	proxy := &ServiceGoRPCProxy{
		service: service,
	}

	if tlsConfig != nil {
		proxy.server = gorpc.NewTLSServer(addr, proxy.handler, tlsConfig)
	} else {
		proxy.server = gorpc.NewTCPServer(addr, proxy.handler)
	}

	return proxy
}

func (p *ServiceGoRPCProxy) Start() error {
	return p.server.Start()
}

func (p *ServiceGoRPCProxy) Serve() error {
	return p.server.Serve()
}

func (p *ServiceGoRPCProxy) Stop() {
	p.server.Stop()
}

func (p *ServiceGoRPCProxy) SetCallStatsHandler(handler gotsrpc.GoRPCCallStatsHandlerFun) {
	p.callStatsHandler = handler
}

func (p *ServiceGoRPCProxy) handler(clientAddr string, request interface{}) (response interface{}) {
	start := time.Now()

	reqType := reflect.TypeOf(request).String()
	funcNameParts := strings.Split(reqType, ".")
	funcName := funcNameParts[len(funcNameParts)-1]

	switch funcName {
	case "ServiceHelloRequest":
		req := request.(ServiceHelloRequest)
		retHello_0 := p.service.Hello(req.V)
		response = ServiceHelloResponse{RetHello_0: retHello_0}
	default:
		fmt.Println("Unknown request type", reflect.TypeOf(request).String())
	}

	if p.callStatsHandler != nil {
		p.callStatsHandler(&gotsrpc.CallStats{
			Func:      funcName,
			Package:   "github.com/foomo/gotsrpc/v2/example/monitor/service",
			Service:   "Service",
			Execution: time.Since(start),
		})
	}

	return
}