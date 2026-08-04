package build

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
	"golang.org/x/tools/imports"

	"github.com/foomo/gotsrpc/v3/config"
	"github.com/foomo/gotsrpc/v3/internal/codegen"
	"github.com/foomo/gotsrpc/v3/internal/parser"
)

func Build(conf *config.Config, goPath, goRoot string) error { //nolint:maintidx
	deriveCommonJSMapping(conf)

	mappedTypeScript := map[string]map[string]*codegen.Code{}

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
		log.Print("Building target", "name", name, "package", packageName, "output", outputPath)

		goRPCProxiesFilename := path.Join(outputPath, "gorpc_gen.go")
		goRPCClientsFilename := path.Join(outputPath, "gorpcclient_gen.go")
		goTSRPCProxiesFilename := path.Join(outputPath, "gotsrpc_gen.go")
		goTSRPCClientsFilename := path.Join(outputPath, "gotsrpcclient_gen.go")

		remove := func(filename string) {
			if _, err := os.Stat(filename); err == nil {
				log.Debug("removing existing file", "file", filename)
				_ = os.Remove(filename)
			}
		}
		remove(goRPCProxiesFilename)
		remove(goRPCClientsFilename)
		remove(goTSRPCProxiesFilename)
		remove(goTSRPCClientsFilename)

		workDirectory, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not determine working directory: %w", err)
		}

		vendorDirectory := path.Join(workDirectory, "vendor")

		goPaths := []string{goPath, goRoot}

		if _, err := os.Stat(vendorDirectory); !os.IsNotExist(err) {
			goPaths = append(goPaths, vendorDirectory)
		}

		pkgName, services, structs, scalars, constantTypes, err := parser.Read(goPaths, conf.Module, packageName, target.Services, missingTypes, missingConstants)
		if err != nil {
			return fmt.Errorf("an error occurred while trying to understand your code: %w", err)
		}

		// collect all union structs
		unions := map[string][]string{}

		for _, s := range structs {
			if len(s.Fields) == 0 && len(s.UnionFields) > 0 {
				unions[s.Package] = append(unions[s.Package], s.Name)
			}
		}

		if target.Out != "" {
			ts, err := codegen.RenderTypeScriptServices(services, conf.Mappings, scalars, structs, target)
			if err != nil {
				return fmt.Errorf("could not generate ts code: %w", err)
			}

			// workaround to remove unneeded imports
			importsCode := codegen.NewCode("	")
			if err := commonJSImports(conf, importsCode, target.Out, ts); err != nil {
				return err
			}

			importsCode.L("").L("")

			ts = importsCode.String() + ts

			if err := updateCode(target.Out, codegen.GetTSHeaderComment()+ts); err != nil {
				return fmt.Errorf("could not write service file %s: %w", target.Out, err)
			}

			if err := codegen.RenderTypescriptStructsToPackages(structs, conf.Mappings, constantTypes, scalars, mappedTypeScript); err != nil {
				return fmt.Errorf("struct gen err for target %s: %w", name, err)
			}
		}

		formatAndWrite := func(code string, filename string) error {
			formattedGoBytes, formattingError := format.Source([]byte(code))
			if formattingError == nil {
				code = string(formattedGoBytes)
			} else {
				log.Warn("could not format generated go code", "file", filename, "err", formattingError)
			}

			codeBytes, errProcessImports := imports.Process(filename, []byte(code), nil)
			if errProcessImports != nil {
				if writeErr := os.WriteFile(filename, []byte(code), 0644); writeErr != nil { //nolint:gosec
					return fmt.Errorf("could not write go source to file %s: %w", filename, writeErr)
				}

				return fmt.Errorf("goimports does not like the generated code (wrote raw code to %s for debugging): %w", filename, errProcessImports)
			}

			if writeErr := os.WriteFile(filename, codeBytes, 0644); writeErr != nil { //nolint:gosec
				return fmt.Errorf("could not write go source to file %s: %w", filename, writeErr)
			}

			return nil
		}

		if len(target.TSRPC) > 0 {
			goTSRPCProxiesCode, goerr := codegen.RenderGoTSRPCProxies(services, packageName, pkgName, target, unions)
			if goerr != nil {
				return fmt.Errorf("could not generate go ts rpc proxies code in target %s: %w", name, goerr)
			}

			if err := formatAndWrite(goTSRPCProxiesCode, goTSRPCProxiesFilename); err != nil {
				return err
			}
		}

		if len(target.TSRPC) > 0 && !target.SkipTSRPCClient {
			goTSRPCClientsCode, goerr := codegen.RenderGoTSRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				return fmt.Errorf("could not generate go ts rpc clients code in target %s: %w", name, goerr)
			}

			if err := formatAndWrite(goTSRPCClientsCode, goTSRPCClientsFilename); err != nil {
				return err
			}
		}

		if len(target.GoRPC) > 0 {
			goRPCProxiesCode, goerr := codegen.RenderGoRPCProxies(services, packageName, pkgName, target)
			if goerr != nil {
				return fmt.Errorf("could not generate go rpc proxies code in target %s: %w", name, goerr)
			}

			if err := formatAndWrite(goRPCProxiesCode, goRPCProxiesFilename); err != nil {
				return err
			}

			goRPCClientsCode, goerr := codegen.RenderGoRPCClients(services, packageName, pkgName, target)
			if goerr != nil {
				return fmt.Errorf("could not generate go rpc clients code in target %s: %w", name, goerr)
			}

			if err := formatAndWrite(goRPCClientsCode, goRPCClientsFilename); err != nil {
				return err
			}
		}
	}

	for goPackage, mappedStructsMap := range mappedTypeScript {
		mapping, ok := conf.Mappings[goPackage]
		if !ok {
			return fmt.Errorf("reverse mapping error in struct generation for package %s", goPackage)
		}

		log.Debug("building structs for go package", "package", goPackage, "module", mapping.TypeScriptModule, "file", mapping.Out)

		moduleCode := codegen.NewCode("	")
		structIndent := -3

		var structNames []string

		for structName := range mappedStructsMap {
			structNames = append(structNames, structName)
		}
		// sort and keep enums on top
		slices.SortFunc(structNames, func(e1 string, e2 string) int {
			es1, ok1 := mappedStructsMap[e1]

			es2, ok2 := mappedStructsMap[e2]
			if !ok1 || !ok2 {
				return strings.Compare(e1, e2)
			}

			es1E := strings.Contains(es1.String(), "export enum ")
			es2E := strings.Contains(es2.String(), "export enum ")

			switch {
			case es1E && !es2E:
				return -1
			case !es1E && es2E:
				return 1
			default:
				return strings.Compare(e1, e2)
			}
		})

		for _, structName := range structNames {
			structCode, ok := mappedStructsMap[structName]
			if ok {
				moduleCode.App(structCode.Ind(structIndent).L("").String())
			}
		}

		moduleCode.L("// end of common js")

		// workaround to remove unneeded imports
		importsCode := codegen.NewCode("	")
		if err := commonJSImports(conf, importsCode, mapping.Out, moduleCode.String()); err != nil {
			return err
		}

		importsCode.L("").L("")

		ts := importsCode.String() + moduleCode.String()

		if err := updateCode(mapping.Out, codegen.GetTSHeaderComment()+ts); err != nil {
			log.Warn("failed to update code", "file", mapping.Out, "err", err)
		}
	}

	return nil
}

func deriveCommonJSMapping(conf *config.Config) {
	replacer := strings.NewReplacer(".", "_", "/", "_", "-", "_")
	for _, mapping := range conf.Mappings {
		mapping.TypeScriptModule = replacer.Replace(mapping.GoPackage)
	}
}

func relativeFilePath(a, b string, appendJsExtension bool) (r string, e error) {
	r, e = filepath.Rel(path.Dir(a), b)
	if e != nil {
		return
	}

	r = strings.TrimSuffix(r, ".ts")
	if appendJsExtension {
		r += ".js"
	}

	return
}

func commonJSImports(conf *config.Config, c *codegen.Code, tsFilename string, code string) error {
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

		relativePath, relativeErr := relativeFilePath(tsFilename, importMapping.Out, conf.TSImportJsExtension)
		if relativeErr != nil {
			return fmt.Errorf("can not derive a relative path between %s and %s: %w", tsFilename, importMapping.Out, relativeErr)
		}

		c.L("import * as " + importMapping.TypeScriptModule + " from './" + relativePath + "'; // " + tsFilename + " to " + importMapping.Out)
	}

	return nil
}

func getPathForTarget(gomod config.Namespace, goPath string, target *config.Target) string {
	if gomod.Name != "" && strings.HasPrefix(target.Package, gomod.Name) {
		relative := strings.TrimPrefix(target.Package, gomod.Name)
		return path.Join(gomod.Path, relative)
	} else {
		return path.Join(goPath, "src", target.Package)
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

	if err := os.MkdirAll(path.Dir(file), 0700); err != nil {
		return err
	}

	oldCode, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if string(oldCode) != code {
		log.Info("writing file", "file", file)

		return os.WriteFile(file, []byte(code), 0600)
	}

	log.Debug("update not necessary - unchanged", "file", file)

	return nil
}
