package main

import (
	"fmt"
	"os"

	"github.com/foomo/gotsrpc"
)

func main() {
	fmt.Println("hello", os.Args[1])
	gotsrpc.Read(os.Args[1], os.Args[2:])
}
