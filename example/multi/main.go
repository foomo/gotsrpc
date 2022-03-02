package main

import (
	"context"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/foomo/gotsrpc/v2/example/multi/service"
)

func main() {
	fs := http.FileServer(http.Dir("./client"))
	ws := service.NewDefaultServiceGoTSRPCProxy(&service.Handler{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/service/"):
			ws.ServeHTTP(w, r)
		default:
			fs.ServeHTTP(w, r)
		}
	})

	go exec.Command("open", "http://127.0.0.1:3000").Run()
	go call()

	http.ListenAndServe("localhost:3000", mux)
}

func call() {
	time.Sleep(time.Second)

	c := service.NewDefaultServiceGoTSRPCClient("http://127.0.0.1:3000")

	{
		res, _ := c.InlineStruct(context.Background())
		spew.Dump(res)
	}

	{
		res, _ := c.InlineStructPtr(context.Background())
		spew.Dump(res) // TODO this should have nil for InlineStructB as for Bug
	}

	{
		res, _ := c.UnionString(context.Background())
		spew.Dump(res)
	}

	{
		res, _ := c.UnionStruct(context.Background())
		spew.Dump(res)
	}
}
