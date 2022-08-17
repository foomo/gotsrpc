package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/config"
)

var (
	version string
)

func usage() {
	fmt.Println("Usage")
	fmt.Println(os.Args[0], " path/to/build-config.yml")
	flag.PrintDefaults()
}

func main() {
	flagDebug := flag.Bool("debug", false, "debug")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}
	if flag.Arg(0) == "version" {
		fmt.Println(version)
		os.Exit(0)
	}
	gotsrpc.ReaderTrace = *flagDebug

	// check if GOPATH has been set as env variable
	// if not use the default from the build pkg
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}

	conf, err := config.LoadConfigFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "config load error, could not load config from", args[0], ":", err)
		os.Exit(2)
	}
	gotsrpc.Build(conf, goPath)
}
