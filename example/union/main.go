package main

import (
	"context"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/foomo/gotsrpc/v2/example/union/service"
)

func main() {
	ctx := context.Background()
	fs := http.FileServer(http.Dir("./union/client"))
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

	go func() {
		time.Sleep(time.Second)
		_ = exec.CommandContext(ctx, "open", "http://127.0.0.1:3000").Run()
		call(ctx)
	}()

	panic(http.ListenAndServe("localhost:3000", mux)) //nolint:gosec
}

func call(ctx context.Context) {
	c := service.NewDefaultServiceGoTSRPCClient("http://127.0.0.1:3000")

	{
		res, _ := c.InlineStruct(ctx)
		spew.Dump(res)
	}

	{
		res, _ := c.InlineStructPtr(ctx)
		spew.Dump(res) // TODO this should have nil for InlineStructB as for Bug
	}

	{
		res, _ := c.UnionString(ctx)
		spew.Dump(res)
	}

	{
		res, _ := c.UnionStruct(ctx)
		spew.Dump(res)
	}
}
