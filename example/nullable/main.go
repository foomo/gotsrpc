package main

import (
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/foomo/gotsrpc/v2/example/nullable/service"
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

	go func() {
		time.Sleep(time.Second)
		_ = exec.Command("open", "http://127.0.0.1:3000").Run()
	}()

	_ = http.ListenAndServe("localhost:3000", mux)
}
