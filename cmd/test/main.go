package main

import (
	"os"

	"github.com/foomo/gotsrpc"
)

func main() {
	//fmt.Println("hello", os.Args[1])
	gotsrpc.Read(os.Args[1], os.Args[2:])
	//gotsrpc.ReadFile("/Users/jan/go/src/github.com/foomo/gotsrpc/demo/demo.go")
}
