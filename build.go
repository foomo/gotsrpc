package gotsrpc

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/foomo/gotsrpc/v2/config"
)

func deriveCommonJSMapping(conf *config.Config) {
	replacer := strings.NewReplacer(".", "_", "/", "_", "-", "_")
	for _, mapping := range conf.Mappings {
		mapping.TypeScriptModule = replacer.Replace(mapping.GoPackage) //strings.Replace(strings.Replace(mapping.GoPackage, ".", "_", -1), "/", "_", -1)
	}
}

func relativeFilePath(a, b string) (r string, e error) {
	r, e = filepath.Rel(path.Dir(a), b)
	if e != nil {
		return
	}
	r = strings.TrimSuffix(r, ".ts")
	return
}

func commonJSImports(conf *config.Config, c *code, tsFilename string) {
	c.l("// hello commonjs - we need some imports - sorted in alphabetical order, by go package")
	packageNames := []string{}
	for packageName := range conf.Mappings {
		packageNames = append(packageNames, packageName)
	}
	sort.Strings(packageNames)
	for _, packageName := range packageNames {
		importMapping := conf.Mappings[packageName]
		relativePath, relativeErr := relativeFilePath(tsFilename, importMapping.Out)
		if relativeErr != nil {
			fmt.Println("can not derive a relative path between", tsFilename, "and", importMapping.Out, relativeErr)
			os.Exit(1)
		}

		c.l("import * as " + importMapping.TypeScriptModule + " from './" + relativePath + "'; // " + tsFilename + " to " + importMapping.Out)
	}

}

func getPathForTarget(gomod config.Namespace, goPath string, target *config.Target) (outputPath string) {
	if gomod.Name != "" && strings.HasPrefix(target.Package, gomod.Name) {
		relative := strings.TrimPrefix(target.Package, gomod.Name)
		return path.Join(gomod.Path, relative)
	} else {
		return path.Join(goPath, "src", target.Package)
	}
}

func Build(conf *config.Config, goPath string) {

	if conf.ModuleKind == config.ModuleKindCommonJS {
		deriveCommonJSMapping(conf)
	}

	mappedTypeScript := map[string]map[string]*code{}

	// preserve alphabetic order
	names := []string{}
	for name := range conf.Targets {
		names = append(names, name)
	}
	sort.Strings(names)

	missingTypes := map[string]bool{}
	for _, mapping := range conf.Mappings {
		for _, include := range mapping.Structs {
			missingTypes[include] = true
		}
	}

	missingConstants := map[string]bool{}
	for _, mapping := range conf.Mappings {
		for _, include := range mapping.Scalars {
			missingConstants[include] = true
		}
	}

	for name, target := range conf.Targets {

		packageName := target.Package
		outputPath := getPathForTarget(conf.Module, goPath, target)
		fmt.Fprintf(os.Stderr, "building target %s (%s -> %s)\n", name, packageName, outputPath)

		goRPCProxiesFilename := path.Join(outputPath, "gorpc_gen.go")
		goRPCClientsFilename := path.Join(outputPath, "gorpcclient_gen.go")
		goTSRPCProxiesFilename := path.Join(outputPath, "gotsrpc_gen.go")
		goTSRPCClientsFilename := path.Join(outputPath, "gotsrpcclient_gen.go")

		remove := func(filename string) {
			_, err := os.Stat(filename)
			if err == nil {
				fmt.Fprintln(os.Stderr, "	removing existing", filename)
				os.Remove(filename)
			}
		}
		remove(goRPCProxiesFilename)
		remove(goRPCClientsFilename)
		remove(goTSRPCProxiesFilename)
		remove(goTSRPCClientsFilename)

		workDirectory, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		vendorDirectory := path.Join(workDirectory, "vendor")

		goPaths := []string{goPath, runtime.GOROOT()}

		if _, err := os.Stat(vendorDirectory); !os.IsNotExist(err) {
			goPaths = append(goPaths, vendorDirectory)
		}

		pkgName, services, structs, scalars, constantTypes, err := Read(goPaths, conf.Module, packageName, target.Services, missingTypes, missingConstants)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\t an error occured while trying to understand your code: ", err)
			os.Exit(2)
		}

		if target.Out != "" {

			ts, err := RenderTypeScriptServices(conf.ModuleKind, conf.TSClientFlavor, services, conf.Mappings, scalars, structs, target)
			if err != nil {
				fmt.Fprintln(os.Stderr, "	could not generate ts code", err)
				os.Exit(3)
			}
			if conf.ModuleKind == config.ModuleKindCommonJS {
				tsClientCode := newCode("	")
				commonJSImports(conf, tsClientCode, target.Out)
				tsClientCode.l("").l("")
				ts = tsClientCode.string() + ts
			}

			// fmt.Fprintln(os.Stdout, ts)
			updateErr := updateCode(target.Out, getTSHeaderComment()+ts)
			if updateErr != nil {
				fmt.Fprintln(os.Stderr, "	could not write service file", target.Out, updateErr)
				os.Exit(3)
			}

			err = renderTypescriptStructsToPackages(conf.ModuleKind, structs, conf.Mappings, constantTypes, scalars, mappedTypeScript)
			if err != nil {
				fmt.Fprintln(os.Stderr, "struct gen err for target", name, err)
				os.Exit(4)
			}
		}

		formatAndWrite := func(code string, filename string) {
			formattedGoBytes, formattingError := format.Source([]byte(code))
			if formattingError == nil {
				code = string(formattedGoBytes)
			} else {
				fmt.Fprintln(os.Stderr, "	could not format go ts rpc proxies code", formattingError)
			}

			codeBytes, errProcessImports := imports.Process(filename, []byte(code), nil)
			if errProcessImports != nil {
				fmt.Fprintln(
					os.Stderr,
					"	goimports does not like the generated code: ",
					errProcessImports,
					", this is the code: ",
					code,
				)
				os.Exit(5)
			}

			writeErr := ioutil.WriteFile(filename, codeBytes, 0644)
			if writeErr != nil {
				fmt.Fprintln(os.Stderr, "	could not write go source to file", writeErr)
				os.Exit(5)
			}
		}
		if len(target.TSRPC) > 0 {
			goTSRPCProxiesCode, goerr := RenderGoTSRPCProxies(services, packageName, pkgName, target)
			if goerr != nil {
				fmt.Fprintln(os.Stderr, "	could not generate go ts rpc proxies code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goTSRPCProxiesCode, goTSRPCProxiesFilename)

			goTSRPCClientsCode, goerr := RenderGoTSRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				fmt.Fprintln(os.Stderr, "	could not generate go ts rpc clients code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goTSRPCClientsCode, goTSRPCClientsFilename)
		}

		if len(target.GoRPC) > 0 {
			goRPCProxiesCode, goerr := RenderGoRPCProxies(services, packageName, pkgName, target)
			if goerr != nil {
				fmt.Fprintln(os.Stderr, "	could not generate go rpc proxies code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goRPCProxiesCode, goRPCProxiesFilename)

			goRPCClientsCode, goerr := RenderGoRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				fmt.Fprintln(os.Stderr, "	could not generate go rpc clients code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goRPCClientsCode, goRPCClientsFilename)
		}
	}

	for goPackage, mappedStructsMap := range mappedTypeScript {
		mapping, ok := conf.Mappings[goPackage]
		if !ok {
			fmt.Fprintln(os.Stderr, "reverse mapping error in struct generation for package", goPackage)
			os.Exit(6)
		}

		fmt.Fprintln(os.Stderr, "building structs for go package", goPackage, "to ts module", mapping.TypeScriptModule, "in file", mapping.Out)
		moduleCode := newCode("	")
		structIndent := -1
		if conf.ModuleKind == config.ModuleKindCommonJS {
			structIndent = -3
			commonJSImports(conf, moduleCode, mapping.Out)
		} else {
			moduleCode.l("module " + mapping.TypeScriptModule + " {").ind(1)
		}

		var structNames []string

		for structName := range mappedStructsMap {
			structNames = append(structNames, structName)
		}
		sort.Strings(structNames)
		for _, structName := range structNames {

			structCode, ok := mappedStructsMap[structName]
			if ok {
				moduleCode.app(structCode.ind(structIndent).l("").string())
			}
		}

		if conf.ModuleKind == config.ModuleKindCommonJS {
			moduleCode.l("// end of common js")
		} else {
			moduleCode.ind(-1).l("}")
		}
		updateErr := updateCode(mapping.Out, getTSHeaderComment()+moduleCode.string())
		if updateErr != nil {
			fmt.Fprintln(os.Stderr, "	failed to update code in", mapping.Out, updateErr)
		}
	}

}

func updateCode(file string, code string) error {
	if len(file) > 0 {
		if file[0] == '~' {
			home := os.Getenv("HOME")
			if len(home) == 0 {
				return errors.New("could not resolve home dir")
			}
			file = path.Join(home, file[1:])
		}
	}
	errMkdirAll := os.MkdirAll(path.Dir(file), 0755)
	if errMkdirAll != nil {
		return errMkdirAll
	}
	oldCode, _ := ioutil.ReadFile(file)
	if string(oldCode) != code {
		fmt.Println("	writing file", file)
		return ioutil.WriteFile(file, []byte(code), 0644)
	}
	fmt.Println("	update file not necessary - unchanged", file)
	return nil
}
