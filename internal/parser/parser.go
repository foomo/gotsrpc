package parser

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"

	"github.com/foomo/gotsrpc/v2/config"
)

// parsedPackage is a minimal stand-in for the deprecated go/ast.Package.
// It exposes only the fields the rest of the parser actually consumes:
// the package's short name and its parsed source files keyed by absolute path.
type parsedPackage struct {
	Name    string
	Files   map[string]*ast.File
	FileSet *token.FileSet
}

// loadPackage resolves and parses a single Go package by import path using
// golang.org/x/tools/go/packages. Delegating to `go list` gives correct
// handling of:
//   - module-versioned import paths (e.g. go.mongodb.org/mongo-driver/v2),
//     whose on-disk directory does not contain the /v2 suffix;
//   - go.mod replace and require directives;
//   - vendored and module-cache lookups.
//
// This replaces the previous hand-rolled goPaths walk together with the
// deprecated go/parser.ParseDir and go/ast.NewPackage calls.
func loadPackage(gomod config.Namespace, packageName string) (*parsedPackage, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedSyntax,
		Dir:   gomod.Path,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, packageName)
	if err != nil {
		return nil, errors.New("could not load package " + packageName + ": " + err.Error())
	}

	if len(pkgs) == 0 {
		return nil, errors.New("package not found: " + packageName)
	}

	var firstErr error

	for _, pkg := range pkgs {
		if pkg.Name == "" || len(pkg.Syntax) == 0 {
			if firstErr == nil && len(pkg.Errors) > 0 {
				firstErr = fmt.Errorf("%v", pkg.Errors)
			}

			continue
		}

		fileNames := pkg.CompiledGoFiles
		if len(fileNames) != len(pkg.Syntax) {
			fileNames = pkg.GoFiles
		}

		files := make(map[string]*ast.File, len(pkg.Syntax))

		for i, syntax := range pkg.Syntax {
			var name string
			if i < len(fileNames) {
				name = fileNames[i]
			} else {
				name = fmt.Sprintf("%s.%d.go", pkg.PkgPath, i)
			}

			if strings.HasSuffix(name, "_test.go") {
				continue
			}

			files[name] = syntax
		}

		if len(files) == 0 {
			continue
		}

		// go/packages parses each file independently, so cross-file identifier
		// references (the ones go/ast.NewPackage used to resolve) come back
		// with Ident.Obj == nil. Several readers downstream rely on Obj being
		// linked to follow named-map / array aliases through the package, so
		// we re-run that linking step ourselves. This avoids depending on the
		// deprecated ast.NewPackage while keeping behavior identical.
		resolvePackageScope(files)

		return &parsedPackage{
			Name:    pkg.Name,
			Files:   files,
			FileSet: pkg.Fset,
		}, nil
	}

	if firstErr != nil {
		return nil, errors.New("could not load package " + packageName + ": " + firstErr.Error())
	}

	return nil, errors.New("package not found: " + packageName)
}

// parsePackage retains its prior signature so existing callers compile
// unchanged. The goPaths argument is no longer used — package resolution
// is fully delegated to go/packages via the module config in gomod.
func parsePackage(_ []string, gomod config.Namespace, packageName string) (*parsedPackage, error) {
	return loadPackage(gomod, packageName)
}

// resolvePackageScope re-creates the cross-file identifier linking that the
// deprecated go/ast.NewPackage used to perform. It walks every file in the
// package, merges all top-level declarations into a single package-level
// scope, then re-resolves each file's Unresolved identifiers against that
// scope, populating Ident.Obj when a match is found.
//
// The legacy ast.Object machinery is itself deprecated in favor of go/types,
// but the existing readers in this package still navigate the AST via
// Ident.Obj, so we replicate that contract locally instead of leaking the
// deprecated ast.NewPackage symbol into the call graph.
func resolvePackageScope(files map[string]*ast.File) {
	pkgScope := ast.NewScope(nil)

	for _, f := range files {
		if f.Scope == nil {
			continue
		}

		for _, obj := range f.Scope.Objects {
			// first declaration wins, mirroring ast.NewPackage's behavior
			_ = pkgScope.Insert(obj)
		}
	}

	for _, f := range files {
		stillUnresolved := f.Unresolved[:0]

		for _, ident := range f.Unresolved {
			if obj := pkgScope.Lookup(ident.Name); obj != nil {
				ident.Obj = obj
			} else {
				stillUnresolved = append(stillUnresolved, ident)
			}
		}

		f.Unresolved = stillUnresolved
	}
}
