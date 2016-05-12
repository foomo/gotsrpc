package main

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

type Field struct {
	Name       string
	ScalarType ScalarType
	Struct     Struct
	IsArray    bool
	IsMap      bool
}

type Func struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	Name   string
	Fields map[string]*Field
	Funcs  map[string]*Field
}

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}

func main() {
	fmt.Println("hello", os.Args[1])
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, os.Args[1], func(info os.FileInfo) bool {
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
					}
					if r.Type().String() == "*ast.TypeSpec" {
						typeSpec := obj.Decl.(*ast.TypeSpec)
						typeSpecRefl := reflect.ValueOf(typeSpec.Type)
						if typeSpecRefl.Type().String() == "*ast.StructType" {
							structType := typeSpec.Type.(*ast.StructType)
							fmt.Println("structType.Fields", structType.Fields)
							for _, field := range structType.Fields.List {
								fmt.Println(name + "." + field.Names[0].Name)
								structs[name].Fields[field.Names[0].Name] = &Field{
									Name: field.Names[0].Name,
								}
								if field.Tag != nil {
									fmt.Println("there is a tag too", field.Tag.Value)
								}
							}
						}
					}
				}
			}
		}
	}
	jsonDump(structs)
}
