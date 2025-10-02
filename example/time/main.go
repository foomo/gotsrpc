package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/foomo/gotsrpc/v2/example/time/service"
)

func main() {
	ctx := context.Background()
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
		t := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local)
		fmt.Printf("%d\n", t.UnixMilli())
		fmt.Println(t.String()) // NOTE: 2022-01-01 00:00:00 +0100 CET
		res, _ := c.Time(ctx, t)
		fmt.Println(res.String()) // NOTE: 2021-12-31 23:00:00 +0000 UTC
		spew.Dump(t.Equal(res))
	}
}
