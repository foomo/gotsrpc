package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/foomo/gotsrpc"
	"github.com/foomo/gotsrpc/config"
)

const Version = "0.11.0"

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Fprintln(os.Stderr, err, string(jsonBytes))
}
func usage() {
	fmt.Println("Usage")
	fmt.Println(os.Args[0], " path/to/build-config.yml")
	flag.PrintDefaults()
}
func main() {

	flagSkipGotsprc := flag.Bool("skipgotsrpc", false, "if true, module GoTSRPC will not be generated")
	flagDebug := flag.Bool("debug", false, "debug")

	flag.Parse()
	gotsrpc.SkipGoTSRPC = *flagSkipGotsprc
	args := flag.Args()
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}
	if flag.Arg(0) == "version" {
		fmt.Println(Version)
		os.Exit(0)
	}
	gotsrpc.ReaderTrace = *flagDebug
	goPath := os.Getenv("GOPATH")

	if len(goPath) == 0 {
		fmt.Fprintln(os.Stderr, "GOPATH not set")
		os.Exit(1)
	}

	conf, err := config.LoadConfigFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "config load error, could not load config from", args[0], ":", err)
		os.Exit(2)
	}
	gotsrpc.Build(conf, goPath)
}
