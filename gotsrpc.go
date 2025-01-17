package gotsrpc

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"

	"github.com/foomo/gotsrpc/v2/config"
)

type contextKey string

var (
	// Read-only global compressor pools
	globalCompressorWriterPools = map[Compressor]*sync.Pool{
		CompressorGZIP:   {New: func() interface{} { return gzip.NewWriter(io.Discard) }},
		CompressorSnappy: {New: func() interface{} { return snappy.NewBufferedWriter(io.Discard) }},
	}
	globalCompressorReaderPools = map[Compressor]*sync.Pool{
		CompressorGZIP:   {New: func() interface{} { return &gzip.Reader{} }},
		CompressorSnappy: {New: func() interface{} { return snappy.NewReader(nil) }},
	}
)

const contextStatsKey contextKey = "gotsrpcStats"

func GetCalledFunc(r *http.Request, endPoint string) string {
	return strings.TrimPrefix(r.URL.Path, endPoint+"/")
}

func ErrorFuncNotFound(w http.ResponseWriter) {
	http.Error(w, "method not found", http.StatusNotFound)
}

func ErrorCouldNotReply(w http.ResponseWriter) {
	http.Error(w, "could not reply", http.StatusInternalServerError)
}

func ErrorCouldNotLoadArgs(w http.ResponseWriter) {
	http.Error(w, "could not load args", http.StatusBadRequest)
}

func ErrorMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "you gotta POST", http.StatusMethodNotAllowed)
}

func LoadArgs(args interface{}, callStats *CallStats, r *http.Request) error {
	start := time.Now()
	var bodyReader io.Reader = r.Body
	switch r.Header.Get("Content-Encoding") {
	case "snappy":
		if snappyReader, ok := globalCompressorReaderPools[CompressorSnappy].Get().(*snappy.Reader); ok {
			defer globalCompressorReaderPools[CompressorSnappy].Put(snappyReader)

			snappyReader.Reset(r.Body)
			bodyReader = snappyReader
		}
	case "gzip":
		if gzipReader, ok := globalCompressorReaderPools[CompressorGZIP].Get().(*gzip.Reader); ok {
			defer globalCompressorReaderPools[CompressorGZIP].Put(gzipReader)

			err := gzipReader.Reset(r.Body)
			if err != nil {
				return NewClientError(errors.Wrap(err, "could not create gzip reader"))
			}
			bodyReader = gzipReader
		}
	}
	handle := getHandlerForContentType(r.Header.Get("Content-Type")).handle

	if err := codec.NewDecoder(bodyReader, handle).Decode(args); err != nil {
		return errors.Wrap(err, "could not decode arguments")
	}

	// We need to close the reader at the end of the function
	if writer, ok := bodyReader.(io.Closer); ok {
		if err := writer.Close(); err != nil {
			return errors.Wrap(err, "failed to write to response body")
		}
	}

	if callStats != nil {
		callStats.Unmarshalling = time.Since(start)
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

func GetStatsForRequest(r *http.Request) (*CallStats, bool) {
	if value, ok := r.Context().Value(contextStatsKey).(*CallStats); ok && value != nil {
		return value, true
	} else {
		return &CallStats{}, false
	}
}

func ClearStats(r *http.Request) {
	*r = *r.WithContext(context.WithValue(r.Context(), contextStatsKey, nil))
}

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
			parsedPkg = pkg
			break
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
