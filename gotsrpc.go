package gotsrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"path"
	"sort"
	"strings"
	"time"
)

const contextStatsKey = "gotsrpcStats"

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

func LoadArgs(args interface{}, callStats *CallStats, r *http.Request) error {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	errLoad := loadArgs(&args, body)
	if errLoad != nil {
		return errLoad
	}
	if callStats != nil {
		callStats.Unmarshalling = time.Now().Sub(start)
		callStats.RequestSize = len(body)
	}
	return nil
}

func loadArgs(args interface{}, jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, &args); err != nil {
		return err
	}
	return nil
}

func RequestWithStatsContext(r *http.Request) *http.Request {
	stats := &CallStats{}
	return r.WithContext(context.WithValue(r.Context(), contextStatsKey, stats))
}

func GetStatsForRequest(r *http.Request) *CallStats {
	value := r.Context().Value(contextStatsKey)
	if value == nil {
		return nil
	}
	return value.(*CallStats)
}

func ClearStats(r *http.Request) {
	*r = *r.WithContext(context.WithValue(r.Context(), contextStatsKey, nil))
}

// Reply despite the fact, that this is a public method - do not call it, it will be called by generated code
func Reply(response []interface{}, stats *CallStats, r *http.Request, w http.ResponseWriter) {
	serializationStart := time.Now()
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize response"))
		return
	}
	if stats != nil {
		stats.ResponseSize = len(jsonBytes)
		stats.Marshalling = time.Now().Sub(serializationStart)
	}
	//r = r.WithContext(ctx)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonBytes)
	//fmt.Println("replied with stats", stats, "on", ctx)
}

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		fmt.Println("an error occured", err)
	}
	fmt.Println(string(jsonBytes))
}

func parseDir(goPaths []string, packageName string) (map[string]*ast.Package, error) {
	errorStrings := map[string]string{}
	for _, goPath := range goPaths {
		fset := token.NewFileSet()
		var dir string
		if strings.HasSuffix(goPath, "vendor") {
			dir = path.Join(goPath, packageName)
		} else {
			dir = path.Join(goPath, "src", packageName)
		}
		pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
		if err == nil {
			return pkgs, nil
		}
		errorStrings[dir] = err.Error()
	}
	return nil, errors.New("could not parse dir for package name: " + packageName + " in goPaths " + strings.Join(goPaths, ", ") + " : " + fmt.Sprint(errorStrings))
}

type byLen []string

func (a byLen) Len() int {
	return len(a)
}

func (a byLen) Less(i, j int) bool {
	return len(a[i]) > len(a[j])
}

func (a byLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func parsePackage(goPaths []string, packageName string) (pkg *ast.Package, err error) {
	pkgs, err := parseDir(goPaths, packageName)
	if err != nil {
		return nil, errors.New("could not parse package " + packageName + ": " + err.Error())
	}
	packageNameParts := strings.Split(packageName, "/")
	if len(packageNameParts) == 0 {
		return nil, errors.New("invalid package name given")
	}
	strippedPackageName := packageNameParts[len(packageNameParts)-1]
	foundPackages := []string{}
	sortedGoPaths := make([]string, len(goPaths))
	for iGoPath := range goPaths {
		sortedGoPaths[iGoPath] = goPaths[iGoPath]
	}
	sort.Sort(byLen(sortedGoPaths))

	for pkgName, pkg := range pkgs {
		// fmt.Println("---------------------> got", pkgName, "looking for", packageName, strippedPackageName)
		// fmt.Println(goPaths)
		// if pkgName == "stripe" {
		// 	//spew.Dump(pkg)
		// 	for pkgFile, pkg := range pkg.Files {
		// 		fmt.Println("file = ", pkgFile)
		// 		spew.Dump(pkg)
		// 	}
		// }
		if pkgName == strippedPackageName {
			return pkg, nil
		}

		for pkgFile := range pkg.Files {
			for _, goPath := range sortedGoPaths {
				// fmt.Println("::::::::::::::::::::::::::::::::", iGoPath, goPath)
				prefix := goPath + "/" // + "/src/"
				if strings.HasPrefix(pkgFile, prefix) && !strings.HasSuffix(pkgFile, "_test.go") && !strings.HasSuffix(pkgFile, "_generator.go") {
					trimmedFilename := strings.TrimPrefix(pkgFile, prefix)
					parts := strings.Split(trimmedFilename, "/")
					if len(parts) > 1 {
						parts = parts[0 : len(parts)-1]
						// fmt.Println(">>>>>>", strings.Join(parts, "/"))
						// fmt.Println("==========>", pkgFile, prefix)
						if strings.Join(parts, "/") == packageName {
							return pkg, nil
						}
					}
				}
			}
		}

		foundPackages = append(foundPackages, pkgName)
	}
	return nil, errors.New("package \"" + packageName + "\" not found in " + strings.Join(foundPackages, ", ") + " looking in go paths" + strings.Join(goPaths, ", "))
}
