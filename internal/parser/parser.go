package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/foomo/gotsrpc/v2/config"
)

func parserExcludeFiles(info os.FileInfo) bool {
	return !strings.HasSuffix(info.Name(), "_test.go")
}

func parseDir(goPaths []string, gomod config.Namespace, packageName string) (map[string]*ast.Package, *token.FileSet, error) {
	if gomod.Name != "" && strings.HasPrefix(packageName, gomod.Name) {
		fset := token.NewFileSet()
		dir := strings.Replace(packageName, gomod.Name, gomod.Path, 1)
		pkgs, err := parser.ParseDir(fset, dir, parserExcludeFiles, parser.DeclarationErrors|parser.AllErrors)
		return pkgs, fset, err
	}

	errorStrings := map[string]string{}
	for _, goPath := range goPaths {
		var dir string
		fset := token.NewFileSet()
		if gomod.ModFile != nil {
			for _, rep := range gomod.ModFile.Replace {
				if packageName == rep.Old.Path || strings.HasPrefix(packageName, rep.Old.Path+"/") {
					if strings.HasPrefix(rep.New.String(), ".") || strings.HasPrefix(rep.New.Path, "/") {
						trace("replacing package with local dir", packageName, rep.Old.String(), rep.New.String())
						dir = strings.Replace(packageName, rep.Old.Path, filepath.Join(gomod.Path, rep.New.Path), 1)
					} else {
						trace("replacing package", packageName, rep.Old.String(), rep.New.String())
						dir = strings.TrimSuffix(path.Join(goPath, "pkg", "mod", rep.New.String(), strings.TrimPrefix(packageName, rep.Old.Path)), "/")
					}
					break
				}
			}
			if dir == "" {
				for _, req := range gomod.ModFile.Require {
					if packageName == req.Mod.Path || strings.HasPrefix(packageName, req.Mod.Path+"/") {
						trace("resolving mod package", packageName, req.Mod.String())
						dir = strings.TrimSuffix(path.Join(goPath, "pkg", "mod", req.Mod.String(), strings.TrimPrefix(packageName, req.Mod.Path)), "/")
						break
					}
				}
			}
		}
		if dir == "" {
			if strings.HasSuffix(goPath, "vendor") {
				dir = path.Join(goPath, packageName)
			} else {
				dir = path.Join(goPath, "src", packageName)
			}
		}
		pkgs, err := parser.ParseDir(fset, dir, parserExcludeFiles, parser.AllErrors)
		if err == nil {
			return pkgs, fset, nil
		}
		errorStrings[dir] = err.Error()
	}
	return nil, nil, errors.New("could not parse dir for package name: " + packageName + " in goPaths " + strings.Join(goPaths, ", ") + " : " + fmt.Sprint(errorStrings))
}

func parsePackage(goPaths []string, gomod config.Namespace, packageName string) (pkg *ast.Package, err error) {
	pkgs, fset, err := parseDir(goPaths, gomod, packageName)
	if err != nil {
		return nil, errors.New("could not parse package " + packageName + ": " + err.Error())
	}
	packageNameParts := strings.Split(packageName, "/")
	if len(packageNameParts) == 0 {
		return nil, errors.New("invalid package name given")
	}
	strippedPackageName := packageNameParts[len(packageNameParts)-1]
	if len(pkgs) == 1 {
		for _, v := range pkgs {
			strippedPackageName = v.Name
			break
		}
	}
	var foundPackages []string
	sortedGoPaths := make([]string, len(goPaths))
	copy(sortedGoPaths, goPaths)
	sort.Sort(byLen(sortedGoPaths))

	var parsedPkg *ast.Package

Loop:
	for pkgName, pkg := range pkgs {
		if pkgName == strippedPackageName {
			parsedPkg = pkg
			break
		}

		for pkgFile := range pkg.Files {
			for _, goPath := range sortedGoPaths {
				prefix := goPath + "/"
				if strings.HasPrefix(pkgFile, prefix) && !strings.HasSuffix(pkgFile, "_test.go") && !strings.HasSuffix(pkgFile, "_generator.go") {
					trimmedFilename := strings.TrimPrefix(pkgFile, prefix)
					parts := strings.Split(trimmedFilename, "/")
					if len(parts) > 1 {
						parts = parts[0 : len(parts)-1]
						if strings.Join(parts, "/") == packageName {
							parsedPkg = pkg
							break Loop
						}
					}
				}
			}
		}

		foundPackages = append(foundPackages, pkgName)
	}

	if parsedPkg == nil {
		return nil, errors.New("package \"" + packageName + "\" not found in " + strings.Join(foundPackages, ", ") + " looking in go paths" + strings.Join(goPaths, ", "))
	}

	// create new package with resolved objects
	resolvedPkg, _ := ast.NewPackage(fset, parsedPkg.Files, nil, nil) // ignore error
	return resolvedPkg, nil
}
