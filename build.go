package gotsrpc

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/foomo/gotsrpc/v2/config"
)

func deriveCommonJSMapping(conf *config.Config) {
	replacer := strings.NewReplacer(".", "_", "/", "_", "-", "_")
	for _, mapping := range conf.Mappings {
		mapping.TypeScriptModule = replacer.Replace(mapping.GoPackage)
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

func commonJSImports(conf *config.Config, c *code, tsFilename string, code string) {
	packageNames := make([]string, 0, len(conf.Mappings))
	for packageName := range conf.Mappings {
		packageNames = append(packageNames, packageName)
	}
	sort.Strings(packageNames)
	for _, packageName := range packageNames {
		importMapping := conf.Mappings[packageName]

		if len(code) > 0 && !strings.Contains(code, importMapping.TypeScriptModule+".") {
			continue
		}

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

func Build(conf *config.Config, goPath, goRoot string) { //nolint:maintidx
	deriveCommonJSMapping(conf)

	mappedTypeScript := map[string]map[string]*code{}

	// preserve alphabetic order
	var names []string
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

	for _, name := range names {
		target := conf.Targets[name]

		packageName := target.Package
		outputPath := getPathForTarget(conf.Module, goPath, target)
		_, _ = fmt.Fprintf(os.Stderr, "building target %s (%s -> %s)\n", name, packageName, outputPath)

		goRPCProxiesFilename := path.Join(outputPath, "gorpc_gen.go")
		goRPCClientsFilename := path.Join(outputPath, "gorpcclient_gen.go")
		goTSRPCProxiesFilename := path.Join(outputPath, "gotsrpc_gen.go")
		goTSRPCClientsFilename := path.Join(outputPath, "gotsrpcclient_gen.go")

		remove := func(filename string) {
			_, err := os.Stat(filename)
			if err == nil {
				_, _ = fmt.Fprintln(os.Stderr, "	removing existing", filename)
				os.Remove(filename)
			}
		}
		remove(goRPCProxiesFilename)
		remove(goRPCClientsFilename)
		remove(goTSRPCProxiesFilename)
		remove(goTSRPCClientsFilename)

		workDirectory, err := os.Getwd()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		vendorDirectory := path.Join(workDirectory, "vendor")

		goPaths := []string{goPath, goRoot}

		if _, err := os.Stat(vendorDirectory); !os.IsNotExist(err) {
			goPaths = append(goPaths, vendorDirectory)
		}

		pkgName, services, structs, scalars, constantTypes, err := Read(goPaths, conf.Module, packageName, target.Services, missingTypes, missingConstants)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "\t an error occurred while trying to understand your code: ", err)
			os.Exit(2)
		}

		// collect all union structs
		unions := map[string][]string{}
		for _, s := range structs {
			if len(s.Fields) == 0 && len(s.UnionFields) > 0 {
				unions[s.Package] = append(unions[s.Package], s.Name)
			}
		}

		if target.Out != "" {
			ts, err := RenderTypeScriptServices(services, conf.Mappings, scalars, structs, target)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not generate ts code", err)
				os.Exit(3)
			}

			// workaround to remove unneeded imports
			importsCode := newCode("	")
			commonJSImports(conf, importsCode, target.Out, ts)
			importsCode.l("").l("")
			ts = importsCode.string() + ts

			// _, _ = fmt.Fprintln(os.Stdout, ts)
			updateErr := updateCode(target.Out, getTSHeaderComment()+ts)
			if updateErr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not write service file", target.Out, updateErr)
				os.Exit(3)
			}

			err = renderTypescriptStructsToPackages(structs, conf.Mappings, constantTypes, scalars, mappedTypeScript)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "struct gen err for target", name, err)
				os.Exit(4)
			}
		}

		formatAndWrite := func(code string, filename string) {
			formattedGoBytes, formattingError := format.Source([]byte(code))
			if formattingError == nil {
				code = string(formattedGoBytes)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "	could not format go ts rpc proxies code", formattingError)
			}

			codeBytes, errProcessImports := imports.Process(filename, []byte(code), nil)
			if errProcessImports != nil {
				_, _ = fmt.Fprintln(
					os.Stderr,
					"	goimports does not like the generated code: ",
					errProcessImports,
				)

				// write code into file for debugging
				writeErr := os.WriteFile(filename, []byte(code), 0644) //nolint:gosec
				if writeErr != nil {
					_, _ = fmt.Fprintln(os.Stderr, "	could not write go source to file", writeErr)
					os.Exit(5)
				}
				_, _ = fmt.Fprintln(os.Stderr, "wrote code for debugging into file", filename)

				os.Exit(5)
			}

			writeErr := os.WriteFile(filename, codeBytes, 0644) //nolint:gosec
			if writeErr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not write go source to file", writeErr)
				os.Exit(5)
			}
		}
		if len(target.TSRPC) > 0 {
			goTSRPCProxiesCode, goerr := RenderGoTSRPCProxies(services, packageName, pkgName, target, unions)
			if goerr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not generate go ts rpc proxies code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goTSRPCProxiesCode, goTSRPCProxiesFilename)
		}
		if len(target.TSRPC) > 0 && !target.SkipTSRPCClient {
			goTSRPCClientsCode, goerr := RenderGoTSRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not generate go ts rpc clients code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goTSRPCClientsCode, goTSRPCClientsFilename)
		}

		if len(target.GoRPC) > 0 {
			goRPCProxiesCode, goerr := RenderGoRPCProxies(services, packageName, pkgName, target)
			if goerr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not generate go rpc proxies code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goRPCProxiesCode, goRPCProxiesFilename)

			goRPCClientsCode, goerr := RenderGoRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				_, _ = fmt.Fprintln(os.Stderr, "	could not generate go rpc clients code in target", name, goerr)
				os.Exit(4)
			}
			formatAndWrite(goRPCClientsCode, goRPCClientsFilename)
		}
	}

	for goPackage, mappedStructsMap := range mappedTypeScript {
		mapping, ok := conf.Mappings[goPackage]
		if !ok {
			_, _ = fmt.Fprintln(os.Stderr, "reverse mapping error in struct generation for package", goPackage)
			os.Exit(6)
		}

		_, _ = fmt.Fprintln(os.Stderr, "building structs for go package", goPackage, "to ts module", mapping.TypeScriptModule, "in file", mapping.Out)
		moduleCode := newCode("	")
		structIndent := -3

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

		moduleCode.l("// end of common js")

		// workaround to remove unneeded imports
		importsCode := newCode("	")
		commonJSImports(conf, importsCode, mapping.Out, moduleCode.string())
		importsCode.l("").l("")
		ts := importsCode.string() + moduleCode.string()

		updateErr := updateCode(mapping.Out, getTSHeaderComment()+ts)
		if updateErr != nil {
			_, _ = fmt.Fprintln(os.Stderr, "	failed to update code in", mapping.Out, updateErr)
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
	oldCode, _ := os.ReadFile(file)
	if string(oldCode) != code {
		fmt.Println("	writing file", file)
		return os.WriteFile(file, []byte(code), 0644) //nolint:gosec
	}
	fmt.Println("	update file not necessary - unchanged", file)
	return nil
}
