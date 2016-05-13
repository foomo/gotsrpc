package gotsrpc

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

type ScalarType string

const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeNumber            = "number"
	ScalarTypeBool              = "bool"
)

type JSONInfo struct {
	Name            string
	OmitEmpty       bool
	ForceStringType bool
	Ignore          bool
}

type Field struct {
	Name       string     `json:",omitempty"`
	JSONInfo   *JSONInfo  `json:",omitempty"`
	ScalarType ScalarType `json:",omitempty"`
	Struct     *Struct    `json:",omitempty"`
	IsArray    bool       `json:",omitempty"`
	IsMap      bool       `json:",omitempty"`
}

type Func struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	Name   string
	Fields map[string]*Field
	Funcs  map[string]*Func
}

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}

func extractJSON(tag string) *JSONInfo {
	t := reflect.StructTag(tag)
	jsonTagString := t.Get("json")
	if len(jsonTagString) == 0 {
		return nil
	}
	jsonTagParts := strings.Split(jsonTagString, ",")

	name := ""
	ignore := false
	omit := false
	forceStringType := false
	cleanParts := []string{}
	for _, jsonTagPart := range jsonTagParts {
		cleanParts = append(cleanParts, strings.TrimSpace(jsonTagPart))
	}
	switch len(cleanParts) {
	case 1:
		switch cleanParts[0] {
		case "-":
			ignore = true
		default:
			name = cleanParts[0]
		}
	case 2:
		if len(cleanParts[0]) > 0 {
			name = cleanParts[0]
		}
		switch cleanParts[1] {
		case "omitempty":
			omit = true
		case "string":
			forceStringType = true
		}
	}
	return &JSONInfo{
		Name:            name,
		OmitEmpty:       omit,
		ForceStringType: forceStringType,
		Ignore:          ignore,
	}
}

func Read(dir string, services []string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		name := info.Name()
		parseIt := !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
		return parseIt
	}, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pkgs, err)

	structs := map[string]*Struct{}

	for pkgName, pkg := range pkgs {
		fmt.Println("pkg", pkgName) //, "scope", pkg.Scope, "files", pkg.Files)
		for filename, file := range pkg.Files {
			fmt.Println(filename, file.Scope.Objects)
			for name, obj := range file.Scope.Objects {

				//fmt.Println(name, obj.Kind, obj.Data)
				if obj.Kind == ast.Typ && obj.Decl != nil {
					//ast.StructType
					r := reflect.ValueOf(obj.Decl)
					fmt.Println(name, "typ decl", obj.Decl, r.Type().String())
					structs[name] = &Struct{
						Name:   name,
						Fields: map[string]*Field{},
						Funcs:  map[string]*Func{},
					}
					if r.Type().String() == "*ast.TypeSpec" {
						typeSpec := obj.Decl.(*ast.TypeSpec)
						typeSpecRefl := reflect.ValueOf(typeSpec.Type)
						if typeSpecRefl.Type().String() == "*ast.StructType" {
							structType := typeSpec.Type.(*ast.StructType)
							fmt.Println("structType.Fields", structType.Fields)
							for _, field := range structType.Fields.List {
								fmt.Println(name + "." + field.Names[0].Name)
								currentField := &Field{
									Name: field.Names[0].Name,
								}
								structs[name].Fields[field.Names[0].Name] = currentField
								if field.Tag != nil {
									jsonTagString := field.Tag.Value[1 : len(field.Tag.Value)-1]
									fmt.Println("there is a tag too", jsonTagString)
									jsonDump(extractJSON(jsonTagString))
									currentField.JSONInfo = extractJSON(jsonTagString)
								}
							}
						}
					}
				}
			}
		}
	}
	jsonDump(structs)

	sepp := reflect.StructTag(`json:",hello"`)
	fmt.Println(sepp.Get("json"))

}
