package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/example/errors/handler/backend"
	backendsvs "github.com/foomo/gotsrpc/v2/example/errors/service/backend"
)

func main() {
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
