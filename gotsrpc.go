package gotsrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"path"
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
	var errDecode error
	switch r.Header.Get("Content-Type") {
	case msgpackContentType:
		errDecode = codec.NewDecoder(r.Body, msgpackHandle).Decode(args)
	default:
		errDecode = codec.NewDecoder(r.Body, jsonHandle).Decode(args)
	}

	if errDecode != nil {
		return errors.Wrap(errDecode, "could not decode arguments")
	}
	if callStats != nil {
		callStats.Unmarshalling = time.Now().Sub(start)
		callStats.RequestSize = int(r.ContentLength)
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
	writer := newResponseWriterWithLength(w)
	serializationStart := time.Now()
	var errEncode error

	switch r.Header.Get("Accept") {
	case msgpackContentType:
		writer.Header().Set("Content-Type", msgpackContentType)
		errEncode = codec.NewEncoder(writer, msgpackHandle).Encode(response)
	case jsonContentType:
		writer.Header().Set("Content-Type", jsonContentType)
		errEncode = codec.NewEncoder(writer, jsonHandle).Encode(response)
	default:
		writer.Header().Set("Content-Type", jsonContentType)
		errEncode = codec.NewEncoder(writer, jsonHandle).Encode(response)
	}

	if errEncode != nil {
		fmt.Println(errEncode)
		http.Error(w, "could not encode data to accepted format", http.StatusInternalServerError)
		return
	}

	if stats != nil {
		stats.ResponseSize = writer.length
		stats.Marshalling = time.Now().Sub(serializationStart)
	}
	//writer.WriteHeader(http.StatusOK)
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
	for pkgName, pkg := range pkgs {
		if pkgName == strippedPackageName {
			return pkg, nil
		}
		foundPackages = append(foundPackages, pkgName)
	}
	return nil, errors.New("package \"" + packageName + "\" not found in " + strings.Join(foundPackages, ", "))
}
