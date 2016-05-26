package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/foomo/gotsrpc"
)

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}
func main() {
	//fmt.Println("hello", os.Args[1])
	gotsrpc.ReaderTrace = true
	//gotsrpc.Read(os.Args[1], os.Args[2:])
	//gotsrpc.ReadFile("/Users/jan/go/src/github.com/foomo/gotsrpc/demo/demo.go", []string{"Service"})
	goPath := os.Getenv("GOPATH")

	if len(goPath) == 0 {
		fmt.Println("GOPATH not set")
		os.Exit(1)
	}
	jsonDump(os.Args[2:])
	services, err := gotsrpc.ReadServicesInPackage(goPath, os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Println("an error occured", err)
		os.Exit(2)
	}
	jsonDump(services)

	//gotsrpc.ReadFile("/Users/jan/go/src/github.com/foomo/gotsrpc/demo/demo.go", []string{"Service"})
}
