package demo_test

import (
	"net/http"

	"github.com/foomo/gotsrpc"
)

type ServiceGoTSRPCProxy struct {
	EndPoint string
	service  *Service
}

func NewServiceGoTSRPCProxy(service *Service, endpoint string) *ServiceGoTSRPCProxy {
	return &ServiceGoTSRPCProxy{
		EndPoint: endpoint,
		service:  service,
	}
}

func (sp *ServiceGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	switch gotsrpc.GetCalledFunc(r, sp.EndPoint) {
	case "Hello":
		args := []interface{}{""}
		err := gotsrpc.LoadArgs(args, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		helloReply, helloErr := sp.service.Hello(args[0].(string))
		gotsrpc.Reply([]interface{}{helloReply, helloErr}, w)
		return
	default:
		gotsrpc.ErrorFuncNotFound(w)
	}
}
