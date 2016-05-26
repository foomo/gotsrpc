package gotsrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"path"
	"strings"
)

func GetCalledFunc(r *http.Request, endPoint string) string {
	return strings.TrimPrefix(r.URL.Path, endPoint+"/")
}

func ErrorFuncNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("method not found"))
}

func ErrorCouldNotLoadArgs(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("could not load args"))
}

func ErrorMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("you gotta POST"))
}

func LoadArgs(args []interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func Reply(response []interface{}, w http.ResponseWriter) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize response"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}

func parsePackage(goPath string, packageName string) (pkg *ast.Package, err error) {
	fset := token.NewFileSet()
	dir := path.Join(goPath, "src", packageName)
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	packageNameParts := strings.Split(packageName, "/")
	if len(packageNameParts) == 0 {
		return nil, errors.New("invalid package name given")
	}
	strippedPackageName := packageNameParts[len(packageNameParts)-1]
	foundPackages := []string{}
	for pkgName, pkg := range pkgs {
		fmt.Println("pkgName", pkgName)
		if pkgName == strippedPackageName {
			return pkg, nil
		}
		foundPackages = append(foundPackages, pkgName)
	}
	return nil, errors.New("package \"" + packageName + "\" not found in " + strings.Join(foundPackages, ", "))
}
