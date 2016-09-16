package gotsrpc

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

func Build(conf *config.Config, goPath string) {
	mappedTypeScript := map[string]map[string]*code{}
	for name, target := range conf.Targets {
		fmt.Fprintln(os.Stderr, "building target", name)
		longPackageName := target.Package
		longPackageNameParts := strings.Split(longPackageName, "/")
		goFilename := path.Join(goPath, "src", longPackageName, "gotsrpc.go")

		_, err := os.Stat(goFilename)
		if err == nil {
			fmt.Fprintln(os.Stderr, "	removing existing", goFilename)
			os.Remove(goFilename)
		}

		packageName := longPackageNameParts[len(longPackageNameParts)-1]

		services, structs, scalarTypes, constants, err := Read(goPath, longPackageName, target.Services)

		if err != nil {
			fmt.Fprintln(os.Stderr, "	an error occured while trying to understand your code", err)
			os.Exit(2)
		}
		ts, err := RenderTypeScriptServices(services, conf.Mappings, scalarTypes, target.TypeScriptModule)
		if err != nil {
			fmt.Fprintln(os.Stderr, "	could not generate ts code", err)
			os.Exit(3)
		}

		// fmt.Fprintln(os.Stdout, ts)
		updateErr := updateCode(target.Out, ts)
		if updateErr != nil {
			fmt.Fprintln(os.Stderr, "	could not write service file", target.Out, updateErr)
			os.Exit(3)
		}
		err = RenderStructsToPackages(structs, conf.Mappings, constants, scalarTypes, mappedTypeScript)
		if err != nil {
			fmt.Fprintln(os.Stderr, "struct gen err for target", name, err)
			os.Exit(4)
		}

		gocode, goerr := RenderGo(services, longPackageName, packageName)
		if goerr != nil {
			fmt.Fprintln(os.Stderr, "	could not generate go code in target", name, goerr)
			os.Exit(4)
		}

		formattedGoBytes, formattingError := format.Source([]byte(gocode))
		if formattingError == nil {
			gocode = string(formattedGoBytes)
		} else {
			fmt.Fprintln(os.Stderr, "	could not format go code", formattingError)
		}

		writeErr := ioutil.WriteFile(goFilename, []byte(gocode), 0644)
		if writeErr != nil {
			fmt.Fprintln(os.Stderr, "	could not write go source to file", writeErr)
			os.Exit(5)
		}
	}

	for goPackage, mappedStructsMap := range mappedTypeScript {
		mapping, ok := conf.Mappings[goPackage]
		if !ok {
			fmt.Fprintln(os.Stderr, "reverse mapping error in struct generation for package", goPackage)
			os.Exit(6)
		}

		fmt.Fprintln(os.Stderr, "building structs for go package", goPackage, "to ts module", mapping.TypeScriptModule, "in file", mapping.Out)
		moduleCode := newCode("	").l("module " + mapping.TypeScriptModule + " {").ind(1)
		structNames := []string{"___goConstants"}

		for structName := range mappedStructsMap {
			structNames = append(structNames, structName)
		}
		sort.Strings(structNames)
		for _, structName := range structNames {

			structCode, ok := mappedStructsMap[structName]
			if ok {
				moduleCode.app(structCode.ind(-1).l("").string())
			}
		}

		moduleCode.ind(-1).l("}")
		updateErr := updateCode(mapping.Out, moduleCode.string())
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
	oldCode, _ := ioutil.ReadFile(file)
	if string(oldCode) != code {
		fmt.Println("	update file", file)
		return ioutil.WriteFile(file, []byte(code), 0644)
	}
	fmt.Println("	update file not necessary - unchanged", file)
	return nil
}
