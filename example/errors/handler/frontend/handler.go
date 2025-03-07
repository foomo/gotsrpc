package frontend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/foomo/gotsrpc/v2/example/errors/service/backend"
	"github.com/foomo/gotsrpc/v2/example/errors/service/frontend"
)

type Handler struct {
	client backend.ServiceGoTSRPCClient
}

func New(client backend.ServiceGoTSRPCClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) Simple(w http.ResponseWriter, r *http.Request) (e *frontend.ErrSimple) {
	fmt.Println("==========> incoming")
	time.Sleep(time.Second)
	if r.Context().Err() != nil {
		fmt.Println("==========>" + r.Context().Err().Error())
	}
	fmt.Println("<========== outgoing")
	return
}

func (h *Handler) Multiple(w http.ResponseWriter, r *http.Request) (e *frontend.ErrMulti) {
	return &frontend.ErrMulti{
		A: frontend.ErrMultiAOne,
		B: "",
	}
}
