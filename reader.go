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

func jsonDump(v interface{}) {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	fmt.Println(err, string(jsonBytes))
}

func extractJSONInfo(tag string) *JSONInfo {
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
func filterFiles(info os.FileInfo) bool {
	name := info.Name()
	parseIt := !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
	return parseIt
}

func ReadFile(file string) {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fset, file, nil, parser.Trace)
	if err != nil {
		fmt.Println(err)
		return
	}
	readFile(file, f)
}

func Read(dir string, services []string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pkgs, err)

	structs := map[string]*Struct{}

	for pkgName, pkg := range pkgs {

		fmt.Println("pkg", pkgName) //, "scope", pkg.Scope, "files", pkg.Files)
		fmt.Println("-------------------------------------------------------------------------")

		for filename, file := range pkg.Files {
			//readFile(filename, file)
			extractStructs(filename, file, structs)
		}
	}
	jsonDump(structs)
	return nil
}

func getScalarFromAstIdent(ident *ast.Ident) ScalarType {
	switch ident.Name {
	case "string":
		return ScalarTypeString
	case "bool":
		return ScalarTypeBool
	case "int", "int32", "int64", "float", "float32", "float64":
		return ScalarTypeNumber
	default:
		return ScalarTypeNone
	}

}

func getTypesFromAstType(ident *ast.Ident) (structType string, scalarType ScalarType) {
	scalarType = getScalarFromAstIdent(ident)
	switch scalarType {
	case ScalarTypeNone:
		structType = ident.Name
	}
	return
}

func readAstType(v *Value, fieldIdent *ast.Ident) {
	structType, scalarType := getTypesFromAstType(fieldIdent)
	v.ScalarType = scalarType
	v.StructType = structType
}

func readAstStarExpr(v *Value, starExpr *ast.StarExpr) {
	switch reflect.ValueOf(starExpr.X).Type().String() {
	case "*ast.Ident":
		ident := starExpr.X.(*ast.Ident)
		v.StructType = ident.Name
		v.IsPtr = true
	case "*ast.StructType":
		// nested anonymous
		structType := starExpr.X.(*ast.StructType)
		v.Struct = &Struct{}
		v.Struct.Fields = readFieldList(structType.Fields.List)
	default:
		fmt.Println("a pointer on what", reflect.ValueOf(starExpr.X).Type().String())
	}
}

func readAstArrayType(v *Value, arrayType *ast.ArrayType) {

	switch reflect.ValueOf(arrayType.Elt).Type().String() {
	case "*ast.StarExpr":
		readAstStarExpr(v, arrayType.Elt.(*ast.StarExpr))
	default:
		fmt.Println("array type elt", reflect.ValueOf(arrayType.Elt).Type().String())
	}

}

func readAstMapType(m *Map, mapType *ast.MapType) {

	fmt.Println("		map key", mapType.Key, reflect.ValueOf(mapType.Key).Type().String())
	fmt.Println("		map value", mapType.Value, reflect.ValueOf(mapType.Value).Type().String())

	// key
	switch reflect.ValueOf(mapType.Key).Type().String() {
	case "*ast.Ident":
		_, scalarType := getTypesFromAstType(mapType.Key.(*ast.Ident))
		m.KeyType = string(scalarType)
	}

	// value
	m.Value.loadExpr(mapType.Value)
}

func (v *Value) loadExpr(expr ast.Expr) {
	//fmt.Println(field.Names[0].Name, field.Type, reflect.ValueOf(field.Type).Type().String())
	switch reflect.ValueOf(expr).Type().String() {
	case "*ast.ArrayType":
		fieldArray := expr.(*ast.ArrayType)
		v.Array = &Array{Value: &Value{}}
		switch reflect.ValueOf(fieldArray.Elt).Type().String() {
		case "*ast.Ident":
			readAstType(v.Array.Value, fieldArray.Elt.(*ast.Ident))
		case "*ast.StarExpr":
			readAstStarExpr(v.Array.Value, fieldArray.Elt.(*ast.StarExpr))
		case "*ast.ArrayType":
			readAstArrayType(v.Array.Value, fieldArray.Elt.(*ast.ArrayType))
		case "*ast.MapType":
			v.Array.Value.Map = &Map{
				Value: &Value{},
			}
			readAstMapType(v.Array.Value.Map, fieldArray.Elt.(*ast.MapType))
		default:
			fmt.Println("---------------------> array of", reflect.ValueOf(fieldArray.Elt).Type().String())
		}
	case "*ast.Ident":
		fieldIdent := expr.(*ast.Ident)
		readAstType(v, fieldIdent)

	case "*ast.StarExpr":
		// a pointer on sth
		readAstStarExpr(v, expr.(*ast.StarExpr))
	case "*ast.MapType":
		v.Map = &Map{
			Value: &Value{},
		}
		readAstMapType(v.Map, expr.(*ast.MapType))
	default:
		fmt.Println("what kind of field ident would that be ?!", reflect.ValueOf(expr).Type().String())

	}
}

func readField(astField *ast.Field) (name string, v *Value, jsonInfo *JSONInfo) {

	name = astField.Names[0].Name
	if strings.Compare(strings.ToLower(name[:1]), name[:1]) == 0 {
		// not exported
		return "", nil, nil
	}
	fmt.Println("	", astField.Names[0].Name)
	v = &Value{}
	v.loadExpr(astField.Type)

	if astField.Tag != nil {
		jsonInfo = extractJSONInfo(astField.Tag.Value[1 : len(astField.Tag.Value)-1])
	}
	return
}
func readFieldList(fieldList []*ast.Field) (fields map[string]*Field) {
	fields = map[string]*Field{}
	for _, field := range fieldList {

		name, value, jsonInfo := readField(field)
		if value != nil {
			fields[name] = &Field{
				Name:     name,
				Value:    value,
				JSONInfo: jsonInfo,
			}
		}

	}
	return
}

func extractStructs(filename string, file *ast.File, structs map[string]*Struct) {
	//fmt.Println("-------------------------------------------------------------------------")
	//jsonDump(file)
	//fmt.Println(filename, file.Scope, len(file.Scope.Objects), file.Scope.Objects)
	//fmt.Println("-------------------------------------------------------------------------")
	for name, obj := range file.Scope.Objects {
		//fmt.Println(name, obj.Kind, obj.Data)
		if obj.Kind == ast.Typ && obj.Decl != nil {
			//ast.StructType
			structs[name] = &Struct{
				Name:   name,
				Fields: map[string]*Field{},
			}
			if reflect.ValueOf(obj.Decl).Type().String() == "*ast.TypeSpec" {
				typeSpec := obj.Decl.(*ast.TypeSpec)
				typeSpecRefl := reflect.ValueOf(typeSpec.Type)
				if typeSpecRefl.Type().String() == "*ast.StructType" {
					structType := typeSpec.Type.(*ast.StructType)
					fmt.Println("structType.Fields", structType.Fields)
					structs[name].Fields = readFieldList(structType.Fields.List)
				} else {
					//	fmt.Println("	what would that be", typeSpecRefl.Type().String())
				}
			} else {
				//fmt.Println("	!!!!!!!!!!!!!!!!!!!!!!!!!!!! not a type spec", r.Type().String())
			}
		} else if obj.Kind == ast.Fun {
			//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> a fucking func / method", obj.Kind, obj)
		} else {
			//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", obj.Kind, obj)
		}
	}
}

func readFile(filename string, file *ast.File) {
	for _, decl := range file.Decls {
		if reflect.ValueOf(decl).Type().String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Recv != nil {
				fmt.Println("that is a method named", funcDecl.Name)
				if len(funcDecl.Recv.List) == 1 {
					firstReceiverField := funcDecl.Recv.List[0]
					if "*ast.StarExpr" == reflect.ValueOf(firstReceiverField.Type).Type().String() {
						starExpr := firstReceiverField.Type.(*ast.StarExpr)
						if "*ast.Ident" == reflect.ValueOf(starExpr.X).Type().String() {
							ident := starExpr.X.(*ast.Ident)
							fmt.Println("	on sth:", ident.Name)
						}

					}
				}
			} else {
				fmt.Println("no receiver for", funcDecl.Name)
			}
		}
	}
}
