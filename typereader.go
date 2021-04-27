package gotsrpc

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

var ReaderTrace = false

func readStructs(pkg *ast.Package, packageName string) (structs map[string]*Struct, scalars map[string]*Scalar, err error) {
	structs = map[string]*Struct{}
	trace("reading files in package", packageName)
	scalars = map[string]*Scalar{}
	errorTypes := map[string]bool{}
	for _, file := range pkg.Files {
		err = extractTypes(file, packageName, structs, scalars)
		if err != nil {
			return
		}

		err = extractErrorTypes(file, packageName, errorTypes)
		if err != nil {
			return
		}
	}
	for name, structType := range structs {
		_, isErrorType := errorTypes[name]
		if isErrorType {
			structType.IsError = true
		}
	}
	//jsonDump(errorTypes)
	//jsonDump(scalarTypes)
	//jsonDump(structs)
	return
}

func trace(args ...interface{}) {
	if ReaderTrace {
		fmt.Fprintln(os.Stderr, args...)
	}
}
func traceData(args ...interface{}) {
	if ReaderTrace {
		for _, arg := range args {
			yamlBytes, errMarshal := yaml.Marshal(arg)
			if errMarshal != nil {
				trace(arg)
				continue
			}
			trace(string(yamlBytes))
		}
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
	omit := false
	inline := false
	ignore := false
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
		case "inline":
			inline = true
		case "omitempty":
			omit = true
		case "string":
			forceStringType = true
		}
	}
	return &JSONInfo{
		Name:            name,
		Inline:          inline,
		OmitEmpty:       omit,
		ForceStringType: forceStringType,
		Ignore:          ignore,
	}
}

func getScalarFromAstIdent(ident *ast.Ident) ScalarType {
	switch ident.Name {
	case "interface":
		return ScalarTypeInterface
	case "string":
		return ScalarTypeString
	case "bool":
		return ScalarTypeBool
	case "byte":
		return ScalarTypeByte
	case "float", "float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return ScalarTypeNumber
	default:
		if ident.Obj != nil && ident.Obj.Decl != nil && reflect.ValueOf(ident.Obj.Decl).Type().String() == "*ast.TypeSpec" {
			typeSpec := ident.Obj.Decl.(*ast.TypeSpec)
			if reflect.ValueOf(typeSpec.Type).Type().String() == "*ast.Ident" {
				return ScalarType(ident.Name) //getScalarFromAstIdent(typeSpec.Type.(*ast.Ident))
			}
		} else if ident.Obj == nil {
			return ScalarType(ident.Name)
		}
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

func readAstType(v *Value, fieldIdent *ast.Ident, fileImports fileImportSpecMap, packageName string) {
	structType, scalarType := getTypesFromAstType(fieldIdent)
	v.ScalarType = scalarType
	if len(structType) > 0 {
		v.StructType = &StructType{
			Name:    structType,
			Package: fileImports.getPackagePath(packageName),
		}
	} else if fieldIdent.Name[:1] == strings.ToUpper(fieldIdent.Name[:1]) {
		v.Scalar = &Scalar{
			Package: fileImports.getPackagePath(packageName),
			Name:    fieldIdent.Name,
			Type:    scalarType,
		}
	} else {
		v.GoScalarType = fieldIdent.Name
	}
}

func readAstStarExpr(v *Value, starExpr *ast.StarExpr, fileImports fileImportSpecMap) {
	v.IsPtr = true
	switch reflect.ValueOf(starExpr.X).Type().String() {
	case "*ast.Ident":
		ident := starExpr.X.(*ast.Ident)
		readAstType(v, ident, fileImports, "")
	case "*ast.StructType":
		// nested anonymous
		readAstStructType(v, starExpr.X.(*ast.StructType), fileImports)
	case "*ast.SelectorExpr":
		readAstSelectorExpr(v, starExpr.X.(*ast.SelectorExpr), fileImports)
	default:
		trace("a pointer on what", reflect.ValueOf(starExpr.X).Type().String())
	}
}

func readAstMapType(m *Map, mapType *ast.MapType, fileImports fileImportSpecMap) {
	trace("		map key", mapType.Key, reflect.ValueOf(mapType.Key).Type().String())
	trace("		map value", mapType.Value, reflect.ValueOf(mapType.Value).Type().String())
	// key
	switch reflect.ValueOf(mapType.Key).Type().String() {
	case "*ast.Ident":
		_, scalarType := getTypesFromAstType(mapType.Key.(*ast.Ident))
		m.KeyType = string(scalarType)
		m.KeyGoType = mapType.Key.(*ast.Ident).Name
	default:
		// todo: implement support for "*ast.Scalar" type (sca)
		// this is important for scalar types in map keys
		// Example:
		//(*ast.MapType)(0xc420e2cc90)({
		//Map: (token.Pos) 276258,
		//		Key: (*ast.SelectorExpr)(0xc420301900)({
		//	X: (*ast.Ident)(0xc4203018c0)(constants),
		//		Sel: (*ast.Ident)(0xc4203018e0)(Site)
		//	}),
		//Value: (*ast.ArrayType)(0xc420e2cc60)({
		//	Lbrack: (token.Pos) 276277,
		//			Len: (ast.Expr) <nil>,
		//			Elt: (*ast.SelectorExpr)(0xc420301960)({
		//		X: (*ast.Ident)(0xc420301920)(elastic),
		//			Sel: (*ast.Ident)(0xc420301940)(CategoryDocument)
		//		})
		//	})
		//})

		//fmt.Println("--------------------------->", reflect.ValueOf(mapType.Key).Type().String())
	}
	// value
	m.Value.loadExpr(mapType.Value, fileImports)
}

func readAstSelectorExpr(v *Value, selectorExpr *ast.SelectorExpr, fileImports fileImportSpecMap) {
	switch reflect.ValueOf(selectorExpr.X).Type().String() {
	case "*ast.Ident":
		// that could be the package name
		//selectorIdent := selectorExpr.X.(*ast.Ident)
		// fmt.Println(selectorExpr, selectorExpr.X.(*ast.Ident))
		//readAstType(v, selectorExpr.X.(*ast.Ident), fileImports)
		readAstType(v, selectorExpr.Sel, fileImports, selectorExpr.X.(*ast.Ident).Name)
		if v.StructType != nil {
			v.StructType.Package = fileImports.getPackagePath(v.StructType.Name)
			v.StructType.Name = selectorExpr.Sel.Name
		}
		//fmt.Println(selectorExpr.X.(*ast.Ident).Name, ".", selectorExpr.Sel)
		//readAstType(v, selectorExpr.Sel, fileImports)
	default:
		trace("selectorExpr.Sel !?", selectorExpr.X, reflect.ValueOf(selectorExpr.X).Type().String())
	}
}

func readAstStructType(v *Value, structType *ast.StructType, fileImports fileImportSpecMap) {
	v.Struct = &Struct{}
	v.Struct.Fields = readFieldList(structType.Fields.List, fileImports)
}

func readAstInterfaceType(v *Value, interfaceType *ast.InterfaceType, fileImports fileImportSpecMap) {
	v.IsInterface = true

}

func (v *Value) loadExpr(expr ast.Expr, fileImports fileImportSpecMap) {

	switch reflect.ValueOf(expr).Type().String() {
	case "*ast.ArrayType":
		fieldArray := expr.(*ast.ArrayType)
		v.Array = &Array{Value: &Value{}}

		switch reflect.ValueOf(fieldArray.Elt).Type().String() {
		case "*ast.ArrayType":
			//readAstArrayType(v.Array.Value, fieldArray.Elt.(*ast.ArrayType), fileImports)
			v.Array.Value.loadExpr(fieldArray.Elt.(*ast.ArrayType), fileImports)
		case "*ast.Ident":
			readAstType(v.Array.Value, fieldArray.Elt.(*ast.Ident), fileImports, "")
		case "*ast.StarExpr":
			readAstStarExpr(v.Array.Value, fieldArray.Elt.(*ast.StarExpr), fileImports)
		case "*ast.MapType":
			v.Array.Value.Map = &Map{
				Value: &Value{},
			}
			readAstMapType(v.Array.Value.Map, fieldArray.Elt.(*ast.MapType), fileImports)
		case "*ast.SelectorExpr":
			readAstSelectorExpr(v.Array.Value, fieldArray.Elt.(*ast.SelectorExpr), fileImports)
		case "*ast.StructType":
			readAstStructType(v.Array.Value, fieldArray.Elt.(*ast.StructType), fileImports)
		case "*ast.InterfaceType":
			readAstInterfaceType(v.Array.Value, fieldArray.Elt.(*ast.InterfaceType), fileImports)
		default:
			trace("---------------------> array of", reflect.ValueOf(fieldArray.Elt).Type().String())
		}
	case "*ast.Ident":
		fieldIdent := expr.(*ast.Ident)
		readAstType(v, fieldIdent, fileImports, "")
	case "*ast.StarExpr":
		// a pointer on sth
		readAstStarExpr(v, expr.(*ast.StarExpr), fileImports)
	case "*ast.MapType":
		v.Map = &Map{
			Value: &Value{},
		}
		readAstMapType(v.Map, expr.(*ast.MapType), fileImports)
	case "*ast.SelectorExpr":
		readAstSelectorExpr(v, expr.(*ast.SelectorExpr), fileImports)
	case "*ast.StructType":
		readAstStructType(v, expr.(*ast.StructType), fileImports)
	case "*ast.InterfaceType":
		readAstInterfaceType(v, expr.(*ast.InterfaceType), fileImports)
	default:
		trace("what kind of field ident would that be ?!", reflect.ValueOf(expr).Type().String())
	}
}

func readField(astField *ast.Field, fileImports fileImportSpecMap) (names []string, v *Value, jsonInfo *JSONInfo) {
	if len(astField.Names) == 0 {
		names = append(names, "")
	} else {
		for _, name := range astField.Names {
			names = append(names, name.Name)
		}
	}
	v = &Value{}
	v.loadExpr(astField.Type, fileImports)
	if astField.Tag != nil {
		jsonInfo = extractJSONInfo(astField.Tag.Value[1 : len(astField.Tag.Value)-1])
	}
	return
}

func readFieldList(fieldList []*ast.Field, fileImports fileImportSpecMap) (fields []*Field) {
	fields = []*Field{}
	for _, field := range fieldList {
		names, value, jsonInfo := readField(field, fileImports)
		if value != nil {
			for _, name := range names {
				if len(name) == 0 {
					if jsonInfo != nil && jsonInfo.Inline {
						if identType, ok := field.Type.(*ast.Ident); ok {
							if typeSpec, ok := identType.Obj.Decl.(*ast.TypeSpec); ok {
								if structType, ok := typeSpec.Type.(*ast.StructType); ok {
									trace("Inline IdentType", identType.Name)
									fields = append(fields, readFieldList(structType.Fields.List, fileImports)...)
									continue
								}
							}
						}
					}
					trace("i do not understand this one", field, names, value, jsonInfo)
					continue
				}
				// this is not unicode proof
				if strings.Compare(strings.ToLower(name[:1]), name[:1]) == 0 {
					continue
				}
				fields = append(fields, &Field{
					Name:     name,
					Value:    value,
					JSONInfo: jsonInfo,
				})
			}
		}
	}
	return
}

func extractErrorTypes(file *ast.File, packageName string, errorTypes map[string]bool) (err error) {
	for _, d := range file.Decls {
		if reflect.ValueOf(d).Type().String() == "*ast.FuncDecl" {
			funcDecl := d.(*ast.FuncDecl)
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) == 1 {
				firstReceiverField := funcDecl.Recv.List[0]
				if "*ast.StarExpr" == reflect.ValueOf(firstReceiverField.Type).Type().String() {
					starExpr := firstReceiverField.Type.(*ast.StarExpr)
					if "*ast.Ident" == reflect.ValueOf(starExpr.X).Type().String() {
						ident := starExpr.X.(*ast.Ident)
						if funcDecl.Name.Name == "Error" && funcDecl.Type.Params.NumFields() == 0 && funcDecl.Type.Results.NumFields() == 1 {
							returnValueField := funcDecl.Type.Results.List[0]
							refl := reflect.ValueOf(returnValueField.Type)
							if refl.Type().String() == "*ast.Ident" {
								returnValueIdent := returnValueField.Type.(*ast.Ident)
								if returnValueIdent.Name == "string" {
									errorTypes[packageName+"."+ident.Name] = true
								}
								//fmt.Println("error for:", ident.Name, returnValueIdent.Name)
							}
						}
					}
				}
			}
		}
	}
	return
}

func extractTypes(file *ast.File, packageName string, structs map[string]*Struct, scalars map[string]*Scalar) error {
	fileImports := getFileImports(file, packageName)
	for name, obj := range file.Scope.Objects {
		if obj.Kind == ast.Typ && obj.Decl != nil {
			structName := packageName + "." + name

			if reflect.ValueOf(obj.Decl).Type().String() == "*ast.TypeSpec" {
				typeSpec := obj.Decl.(*ast.TypeSpec)
				typeSpecRefl := reflect.ValueOf(typeSpec.Type)
				typeName := typeSpecRefl.Type().String()
				switch typeName {
				case "*ast.StructType":
					structs[structName] = &Struct{
						Name:    name,
						Fields:  []*Field{},
						Package: packageName,
					}
					structType := typeSpec.Type.(*ast.StructType)
					trace("StructType", obj.Name)
					structs[structName].Fields = readFieldList(structType.Fields.List, fileImports)
				case "*ast.InterfaceType":
					trace("Interface", obj.Name)
					scalars[structName] = &Scalar{
						Name:    structName,
						Package: packageName,
						Type:    ScalarTypeInterface,
					}
				case "*ast.Ident":
					trace("Scalar", obj.Name)
					scalarIdent := typeSpec.Type.(*ast.Ident)
					scalars[structName] = &Scalar{
						Name:    structName,
						Package: packageName,
						Type:    getScalarFromAstIdent(scalarIdent),
					}
				case "*ast.ArrayType":
					arrayValue := &Value{}
					arrayValue.loadExpr(typeSpec.Type, fileImports)
					structs[structName] = &Struct{
						Name:    name,
						Package: packageName,
						Array:   arrayValue.Array,
					}
				case "*ast.MapType":
					mapValue := &Value{}
					mapValue.loadExpr(typeSpec.Type, fileImports)
					structs[structName] = &Struct{
						Name:    name,
						Package: packageName,
						Map:     mapValue.Map,
					}
				default:
					fmt.Println("	ignoring", obj.Name, typeSpecRefl.Type().String())
				}
			}
		}
	}
	return nil
}
