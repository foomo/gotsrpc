package gotsrpc_test

import (
	"fmt"
	"net/http"

	"github.com/foomo/gotsrpc/v3"
)

func ExampleNewHTTPError() {
	err := gotsrpc.NewHTTPError("method not found", http.StatusNotFound)

	fmt.Println(err)
	fmt.Println(err.StatusCode)
	// Output:
	// [404] method not found
	// 404
}
