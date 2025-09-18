package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/config"
)

var (
	version        = "dev"
	commitHash     = "n/a"
	buildTimestamp = "n/a"
)

func usage() {
	fmt.Print(`gotsrpc

Usage:
  gotsrpc [options] [command] <config-file>

Available Commands
  version    Display version information

Options:
  -debug     Print debug information

Examples:
  $ gotsrpc path/to/gotsrpc.yaml
`)
}

func main() {
	flagDebug := flag.Bool("debug", false, "debug")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	switch {
	case len(args) == 0:
		usage()
		os.Exit(1)
	case args[0] == "version":
		if *flagDebug {
			buildTime := buildTimestamp
			if value, err := strconv.ParseInt(buildTimestamp, 10, 64); err == nil {
				buildTime = time.Unix(value, 0).String()
			}
			fmt.Printf("Version: %s\nCommit: %s\nBuildTime: %s\n", version, commitHash, buildTime)
		} else {
			fmt.Println(version)
		}
		os.Exit(0)
	case len(args) != 1:
		usage()
		os.Exit(1)
	default:
		gotsrpc.ReaderTrace = *flagDebug
	}

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
		_, _ = fmt.Fprintln(os.Stderr, "config load error, could not load config from", args[0], ":", err)
		os.Exit(2)
	}

	gotsrpc.Build(conf, goPath, goRoot)
}
