package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

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

	var goRoot string
	var goPath string
	if out, err := exec.Command("go", "env", "GOROOT").Output(); err != nil {
		fmt.Println("failed to retrieve GOROOT", err.Error())
		os.Exit(1)
	} else {
		goRoot = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("go", "env", "GOPATH").Output(); err != nil {
		fmt.Println("failed to retrieve GOPATH", err.Error())
		os.Exit(1)
	} else {
		goPath = string(bytes.TrimSpace(out))
	}

	conf, err := config.LoadConfigFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "config load error, could not load config from", args[0], ":", err)
		os.Exit(2)
	}

	gotsrpc.Build(conf, goPath, goRoot)
}
