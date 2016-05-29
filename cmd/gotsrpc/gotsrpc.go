package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/foomo/gotsrpc"
)

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}
func usage() {
	fmt.Println("Usage")
	fmt.Println(os.Args[0], " --ts-module MyTS.Module.Name my.server/my/package ServiceA [ ServiceB, ... ]")
	flag.PrintDefaults()
}
func main() {
	flagTsModule := flag.String("ts-module", "", "TypeScript target module")

	flag.Parse()
	if len(*flagTsModule) == 0 {
		fmt.Println("missing ts module")
	}

	args := flag.Args()
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}
	gotsrpc.ReaderTrace = true
	goPath := os.Getenv("GOPATH")

	if len(goPath) == 0 {
		fmt.Println("GOPATH not set")
		os.Exit(1)
	}
	longPackageName := args[0]
	longPackageNameParts := strings.Split(longPackageName, "/")
	goFilename := path.Join(goPath, "src", longPackageName, "gotsrpc.go")

	packageName := longPackageNameParts[len(longPackageNameParts)-1]
	services, structs, err := gotsrpc.Read(goPath, longPackageName, args[1:])

	if err != nil {
		fmt.Println("an error occured", err)
		os.Exit(2)
	}
	jsonDump(services)
	jsonDump(structs)

	ts, err := gotsrpc.RenderTypeScript(services, structs, *flagTsModule)
	if err != nil {
		fmt.Println("could not generate ts code", err)
		os.Exit(3)
	}
	fmt.Println(ts)

	gocode, goerr := gotsrpc.RenderGo(services, packageName)
	if goerr != nil {
		fmt.Println("could not generate go code", goerr)
		os.Exit(4)
	}

	formattedGoBytes, formattingError := format.Source([]byte(gocode))
	if formattingError == nil {
		gocode = string(formattedGoBytes)
	} else {
		fmt.Println("could not format go code", formattingError)
	}

	writeErr := ioutil.WriteFile(goFilename, []byte(gocode), 0644)
	if writeErr != nil {
		fmt.Println("could not write go source to file", writeErr)
		os.Exit(5)

	}
	fmt.Println(goFilename, gocode)
	//gotsrpc.ReadFile("/Users/jan/go/src/github.com/foomo/gotsrpc/demo/demo.go", []string{"Service"})
}
