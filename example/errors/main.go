package main

import (
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/foomo/gotsrpc/v2/example/errors/handler/backend"
	"github.com/foomo/gotsrpc/v2/example/errors/handler/frontend"
	backendsvs "github.com/foomo/gotsrpc/v2/example/errors/service/backend"
	frontendsvs "github.com/foomo/gotsrpc/v2/example/errors/service/frontend"
)

func main() {
	fs := http.FileServer(http.Dir("./client"))
	fh := frontendsvs.NewDefaultServiceGoTSRPCProxy(frontend.New(backendsvs.NewDefaultServiceGoTSRPCClient("http://localhost:3000")))
	bh := backendsvs.NewDefaultServiceGoTSRPCProxy(backend.New())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/service/frontend"):
			fh.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, "/service/backend"):
			bh.ServeHTTP(w, r)
		default:
			fs.ServeHTTP(w, r)
		}
	})

	go func() {
		time.Sleep(time.Second)
		_ = exec.Command("open", "http://127.0.0.1:3000").Run()
	}()

	go func() {
		if err := http.ListenAndServe("localhost:3000", mux); err != nil {
			panic(err)
		}
	}()
}
