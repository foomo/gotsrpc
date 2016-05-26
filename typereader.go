package gotsrpc

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

var ReaderTrace = false

func trace(args ...interface{}) {
	if ReaderTrace {
		fmt.Println(args...)
	}
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

func ReadStructs(pkg *ast.Package, services []string) error {
	structs := map[string]*Struct{}
	for _, file := range pkg.Files {
		//readFile(filename, file)
		extractStructs(file, structs)
	}
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
	if len(structType) > 0 {
		v.StructType = &StructType{
			Name: structType,
		}
	}
}

func readAstStarExpr(v *Value, starExpr *ast.StarExpr) {
	v.IsPtr = true
	switch reflect.ValueOf(starExpr.X).Type().String() {
	case "*ast.Ident":
		ident := starExpr.X.(*ast.Ident)
		v.StructType = &StructType{
			Name: ident.Name,
		}
	case "*ast.StructType":
		// nested anonymous
		readAstStructType(v, starExpr.X.(*ast.StructType))
	case "*ast.SelectorExpr":
		readAstSelectorExpr(v, starExpr.X.(*ast.SelectorExpr))
	default:
		trace("a pointer on what", reflect.ValueOf(starExpr.X).Type().String())
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
	trace("		map key", mapType.Key, reflect.ValueOf(mapType.Key).Type().String())
	trace("		map value", mapType.Value, reflect.ValueOf(mapType.Value).Type().String())
	// key
	switch reflect.ValueOf(mapType.Key).Type().String() {
	case "*ast.Ident":
		_, scalarType := getTypesFromAstType(mapType.Key.(*ast.Ident))
		m.KeyType = string(scalarType)
	}
	// value
	m.Value.loadExpr(mapType.Value)
}

func readAstSelectorExpr(v *Value, selectorExpr *ast.SelectorExpr) {
	switch reflect.ValueOf(selectorExpr.X).Type().String() {
	case "*ast.Ident":
		// that could be the package name
		selectorIdent := selectorExpr.X.(*ast.Ident)
		v.StructType = &StructType{
			Package: selectorIdent.Name,
			Name:    selectorExpr.Sel.Name,
		}
	default:
		fmt.Println("selectorExpr.Sel !?", selectorExpr.X, reflect.ValueOf(selectorExpr.X).Type().String())
	}
}

func readAstStructType(v *Value, structType *ast.StructType) {
	v.Struct = &Struct{}
	v.Struct.Fields = readFieldList(structType.Fields.List)
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
	case "*ast.SelectorExpr":
		readAstSelectorExpr(v, expr.(*ast.SelectorExpr))
	case "*ast.StructType":
		readAstStructType(v, expr.(*ast.StructType))
	default:
		trace("what kind of field ident would that be ?!", reflect.ValueOf(expr).Type().String())
	}
}

func readField(astField *ast.Field) (name string, v *Value, jsonInfo *JSONInfo) {
	name = ""
	if len(astField.Names) > 0 {
		name = astField.Names[0].Name
	}
	trace("	", name)
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

		if strings.Compare(strings.ToLower(name[:1]), name[:1]) == 0 {
			continue
		}

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

func extractStructs(file *ast.File, structs map[string]*Struct) {
	for _, imp := range file.Imports {
		fmt.Println("import", imp.Name, imp.Path)
	}
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
					trace("StructType", obj.Name)
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
