package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/example/monitor/service"
)

func init() {
	// set custom monitor handler
	gotsrpc.Monitor = func(w http.ResponseWriter, r *http.Request, args, rets []interface{}, stats *gotsrpc.CallStats) {
		// you might want to use channels or routines here as to not block the actual call
		go func(stats *gotsrpc.CallStats) {
			fmt.Printf(
				"Monitor: %s.%s [%dμs|%dμs|%dμs]\n",
				stats.Service, stats.Func, stats.Unmarshalling.Microseconds(), stats.Execution.Microseconds(), stats.Marshalling.Microseconds(),
			)
		}(stats)
	}
}

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
		res, _ := c.Hello(ctx, "Hello World")
		fmt.Println(res)
	}
}
