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
	"github.com/foomo/gotsrpc/config"
)

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Fprintln(os.Stderr, err, string(jsonBytes))
}
func usage() {
	fmt.Println("Usage")
	fmt.Println(os.Args[0], " path/to/build-config.yml [target, [target], ...]")
	flag.PrintDefaults()
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	gotsrpc.ReaderTrace = true
	goPath := os.Getenv("GOPATH")

	if len(goPath) == 0 {
		fmt.Fprintln(os.Stderr, "GOPATH not set")
		os.Exit(1)
	}

	conf, err := config.LoadConfigFile(args[0])
	if err != nil {
		fmt.Println(os.Stderr, "config load error")
		os.Exit(2)
	}
	fmt.Println(conf, err)
	buildTargets := map[string]*config.Target{}
	if len(args) > 1 {
		for _, target := range args[1:] {
			fmt.Println(os.Stderr, "will build target", target)
			_, ok := conf.Targets[target]
			if !ok {
				fmt.Println(os.Stderr, "invalid target has to be one of:")
				for existingTarget := range conf.Targets {
					fmt.Println(os.Stderr, "	", existingTarget)
				}
				os.Exit(1)
			}
			buildTargets[target] = conf.Targets[target]
		}
	} else {
		fmt.Println(os.Stderr, "will build all targets in config")
		buildTargets = conf.Targets
	}
	fmt.Println(os.Stderr, buildTargets)

	for name, target := range buildTargets {
		fmt.Println(os.Stderr, "building target", name)
		longPackageName := target.Package
		longPackageNameParts := strings.Split(longPackageName, "/")
		goFilename := path.Join(goPath, "src", longPackageName, "gotsrpc.go")

		_, err := os.Stat(goFilename)
		if err == nil {
			fmt.Fprintln(os.Stderr, "removing existing", goFilename)
			os.Remove(goFilename)
		}

		packageName := longPackageNameParts[len(longPackageNameParts)-1]
		services, structs, err := gotsrpc.Read(goPath, longPackageName, target.Services)

		if err != nil {
			fmt.Fprintln(os.Stderr, "an error occured while trying to understand your code", err)
			os.Exit(2)
		}
		ts, err := gotsrpc.RenderTypeScriptServices(services, conf.Mappings, target.TypeScriptModule)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not generate ts code", err)
			os.Exit(3)
		}

		fmt.Println(os.Stdout, ts)

		mappedCode, err := gotsrpc.RenderStructsToPackages(structs, conf.Mappings)
		if err != nil {
			fmt.Println("struct gen err", err)
			os.Exit(4)
		}

		for tsModule, code := range mappedCode {
			fmt.Println("-----------------", tsModule, "--------------------")
			fmt.Println(code)
		}

		gocode, goerr := gotsrpc.RenderGo(services, packageName)
		if goerr != nil {
			fmt.Fprintln(os.Stderr, "could not generate go code", goerr)
			os.Exit(4)
		}

		formattedGoBytes, formattingError := format.Source([]byte(gocode))
		if formattingError == nil {
			gocode = string(formattedGoBytes)
		} else {
			fmt.Fprintln(os.Stderr, "could not format go code", formattingError)
		}

		writeErr := ioutil.WriteFile(goFilename, []byte(gocode), 0644)
		if writeErr != nil {
			fmt.Fprintln(os.Stderr, "could not write go source to file", writeErr)
			os.Exit(5)
		}
	}

}
