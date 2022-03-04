package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/foomo/gotsrpc/v2"
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
		call()
	}()

	panic(http.ListenAndServe("localhost:3000", mux))
}

func call() {
	ctx := context.Background()
	c := backendsvs.NewDefaultServiceGoTSRPCClient("http://localhost:3000")

	{
		fmt.Println("-------------------------")
		var gotsrpcErr *gotsrpc.Error
		serviceErr, err := c.Error(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Println("OK")
		}
		if errors.As(serviceErr, &gotsrpcErr) {
			fmt.Printf("%s\n", gotsrpcErr)
			fmt.Printf("%q\n", gotsrpcErr)
			fmt.Printf("%+v\n", gotsrpcErr)
		}
	}

	{
		fmt.Println("-------------------------")
		scalar, err := c.Scalar(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if scalar == nil {
			panic("service error should not be nil")
		} else if scalar != nil {
			fmt.Println("OK")
		}
	}

	{
		fmt.Println("-------------------------")
		scalar, err := c.MultiScalar(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if scalar == nil {
			panic("service error should not be nil")
		} else if scalar != nil {
			fmt.Println("OK")
		}
	}

	{
		fmt.Println("-------------------------")
		var gotsrpcErr *gotsrpc.Error
		serviceErr, err := c.WrappedError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Println("OK")
		}
		if errors.As(serviceErr, &gotsrpcErr) {
			fmt.Println(gotsrpcErr.Error())
			fmt.Printf("%+v\n", gotsrpcErr)
		}
		if errors.As(errors.Unwrap(serviceErr), &gotsrpcErr) {
			fmt.Println(gotsrpcErr.Error())
		}
	}

	{
		fmt.Println("-------------------------")
		var scalarErr *backend.ScalarError
		var gotsrpcErr *gotsrpc.Error
		serviceErr, err := c.ScalarError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Printf("%s\n", serviceErr)
			fmt.Printf("%q\n", serviceErr)
			fmt.Printf("%+v\n", serviceErr)
		}
		if errors.As(serviceErr, &gotsrpcErr) {
			fmt.Println(gotsrpcErr)
		}
		if errors.As(serviceErr, &scalarErr) {
			fmt.Println(scalarErr)
		}
	}

	{
		fmt.Println("-------------------------")
		var customErr *backend.CustomError
		var gotsrpcErr *gotsrpc.Error
		serviceErr, err := c.CustomError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Printf("%s\n", serviceErr)
			fmt.Printf("%q\n", serviceErr)
			fmt.Printf("%+v\n", serviceErr)
		}
		if errors.As(serviceErr, &gotsrpcErr) {
			fmt.Println(gotsrpcErr)
		}
		if errors.As(serviceErr, &customErr) {
			fmt.Println(customErr)
		}
	}

	{
		fmt.Println("-------------------------")
		serviceErr, err := c.TypedError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Println("OK")
		}
		if errors.Is(serviceErr, backend.ErrTyped) {
			fmt.Println("OK")
		}
	}

	{
		fmt.Println("-------------------------")
		serviceErr, err := c.TypedWrappedError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Println("OK")
		}
		if errors.Is(serviceErr, backend.ErrTyped) {
			fmt.Println("OK")
		}
	}

	{
		fmt.Println("-------------------------")
		serviceErr, err := c.TypedCustomError(ctx)
		if err != nil {
			panic("client error should be nil")
		} else if serviceErr == nil {
			panic("service error should not be nil")
		} else if serviceErr != nil {
			fmt.Println("OK")
		}
		if errors.Is(serviceErr, backend.ErrCustom) {
			fmt.Println("OK")
		}
	}
}
