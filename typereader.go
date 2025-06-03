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
	structTag := reflect.StructTag(tag)

	jsonTags := strings.Split(structTag.Get("json"), ",")
	gotsrpcTags := strings.Split(structTag.Get("gotsrpc"), ",")
	if len(jsonTags) == 0 && len(gotsrpcTags) == 0 {
		return nil
	}

	for k, value := range jsonTags {
		jsonTags[k] = strings.TrimSpace(value)
	}
	for k, value := range gotsrpcTags {
		gotsrpcTags[k] = strings.TrimSpace(value)
	}

	name := ""
	tsType := ""
	omit := false
	union := false
	inline := false
	ignore := false

	if len(jsonTags) > 0 {
		switch jsonTags[0] {
		case "":
			// do nothing
		case "-":
			ignore = true
		default:
			name = jsonTags[0]
		}
	}
	if len(jsonTags) > 1 {
		for _, value := range jsonTags[1:] {
			switch value {
			case "inline":
				inline = true
			case "omitempty":
				omit = true
			}
		}
	}

	for _, value := range gotsrpcTags {
		switch {
		case value == "union":
			union = true
		case strings.HasPrefix(value, "type:"):
			tsType = strings.TrimPrefix(value, "type:")
		}
	}

	// TODO split up gotsrpc info
	return &JSONInfo{
		Name:      name,
		Type:      tsType,
		Union:     union,
		Inline:    inline,
		OmitEmpty: omit,
		Ignore:    ignore,
	}
}

func getScalarFromAstIdent(ident *ast.Ident) ScalarType {
	switch ident.Name {
	case "any", "interface":
		return ScalarTypeAny
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
			if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
				if reflect.ValueOf(typeSpec.Type).Type().String() == "*ast.Ident" {
					return ScalarType(ident.Name) // getScalarFromAstIdent(typeSpec.Type.(*ast.Ident))
				}
			}
		} else if ident.Obj == nil {
			return ScalarType(ident.Name)
		}
		return ScalarTypeNone
	}
}

func getTypesFromAstType(ident *ast.Ident) (structType string, scalarType ScalarType) {
	scalarType = getScalarFromAstIdent(ident)
	switch scalarType { //nolint:gocritic,exhaustive
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
		if fieldIdent.Name == "error" {
			v.IsError = true
		}
	}
}

func readAstStarExpr(v *Value, starExpr *ast.StarExpr, fileImports fileImportSpecMap) {
	v.IsPtr = true
	switch starExprType := starExpr.X.(type) {
	case *ast.Ident:
		readAstType(v, starExprType, fileImports, "")
	case *ast.StructType:
		// nested anonymous
		readAstStructType(v, starExprType, fileImports)
	case *ast.SelectorExpr:
		readAstSelectorExpr(v, starExprType, fileImports)
	default:
		trace("a pointer on what", reflect.ValueOf(starExpr.X).Type().String())
	}
}

func readAstMapType(m *Map, mapType *ast.MapType, fileImports fileImportSpecMap) {
	trace("		map key", mapType.Key, reflect.ValueOf(mapType.Key).Type().String())
	trace("		map value", mapType.Value, reflect.ValueOf(mapType.Value).Type().String())
	// key
	switch keyType := mapType.Key.(type) {
	case *ast.Ident:
		_, scalarType := getTypesFromAstType(keyType)
		m.KeyType = string(scalarType)
		m.KeyGoType = keyType.Name
		m.Key = &Value{}
		readAstType(m.Key, keyType, fileImports, "")
	case *ast.SelectorExpr:
		m.Key = &Value{}
		readAstSelectorExpr(m.Key, keyType, fileImports)
	default:
		// todo: implement support for "*ast.Scalar" type (sca)
		// this is important for scalar types in map keys
		// Example:
		// (*ast.MapType)(0xc420e2cc90)({
		// Map: (token.Pos) 276258,
		// 		Key: (*ast.SelectorExpr)(0xc420301900)({
		// 	X: (*ast.Ident)(0xc4203018c0)(constants),
		// 		Sel: (*ast.Ident)(0xc4203018e0)(Site)
		// 	}),
		// Value: (*ast.ArrayType)(0xc420e2cc60)({
		// 	Lbrack: (token.Pos) 276277,
		// 			Len: (ast.Expr) <nil>,
		// 			Elt: (*ast.SelectorExpr)(0xc420301960)({
		// 		X: (*ast.Ident)(0xc420301920)(elastic),
		// 			Sel: (*ast.Ident)(0xc420301940)(CategoryDocument)
		// 		})
		// 	})
		// })
	}
	// value
	m.Value.loadExpr(mapType.Value, fileImports)
}

func readAstSelectorExpr(v *Value, selectorExpr *ast.SelectorExpr, fileImports fileImportSpecMap) {
	switch selExpType := selectorExpr.X.(type) {
	case *ast.Ident:
		// that could be the package name
		// selectorIdent := selectorExpr.X.(*ast.Ident)
		// fmt.Println(selectorExpr, selectorExpr.X.(*ast.Ident))
		// readAstType(v, selectorExpr.X.(*ast.Ident), fileImports)
		readAstType(v, selectorExpr.Sel, fileImports, selExpType.Name)
		if v.StructType != nil {
			v.StructType.Package = fileImports.getPackagePath(v.StructType.Name)
			v.StructType.Name = selectorExpr.Sel.Name
		}
		// fmt.Println(selectorExpr.X.(*ast.Ident).Name, ".", selectorExpr.Sel)
		// readAstType(v, selectorExpr.Sel, fileImports)
	default:
		trace("selectorExpr.Sel !?", selectorExpr.X, reflect.ValueOf(selectorExpr.X).Type().String())
	}
}

func readAstStructType(v *Value, structType *ast.StructType, fileImports fileImportSpecMap) {
	v.Struct = &Struct{}
	v.Struct.Fields, v.Struct.InlineFields, v.Struct.UnionFields = readFieldList(structType.Fields.List, fileImports)
}

func readAstInterfaceType(v *Value, interfaceType *ast.InterfaceType, fileImports fileImportSpecMap) {
	v.IsInterface = true
}

func (v *Value) loadExpr(expr ast.Expr, fileImports fileImportSpecMap) {
	switch exprType := expr.(type) {
	case *ast.ArrayType:
		v.Array = &Array{Value: &Value{}}

		switch exprEltType := exprType.Elt.(type) {
		case *ast.ArrayType:
			// readAstArrayType(v.Array.Value, fieldArray.Elt.(*ast.ArrayType), fileImports)
			v.Array.Value.loadExpr(exprEltType, fileImports)
		case *ast.Ident:
			readAstType(v.Array.Value, exprEltType, fileImports, "")
		case *ast.StarExpr:
			readAstStarExpr(v.Array.Value, exprEltType, fileImports)
		case *ast.MapType:
			v.Array.Value.Map = &Map{
				Value: &Value{},
			}
			readAstMapType(v.Array.Value.Map, exprEltType, fileImports)
		case *ast.SelectorExpr:
			readAstSelectorExpr(v.Array.Value, exprEltType, fileImports)
		case *ast.StructType:
			readAstStructType(v.Array.Value, exprEltType, fileImports)
		case *ast.InterfaceType:
			readAstInterfaceType(v.Array.Value, exprEltType, fileImports)
		default:
			trace("---------------------> array of", reflect.ValueOf(exprType.Elt).Type().String())
		}
	case *ast.Ident:
		readAstType(v, exprType, fileImports, "")
	case *ast.StarExpr:
		// a pointer on sth
		readAstStarExpr(v, exprType, fileImports)
	case *ast.MapType:
		v.Map = &Map{
			Value: &Value{},
		}
		readAstMapType(v.Map, exprType, fileImports)
	case *ast.SelectorExpr:
		readAstSelectorExpr(v, exprType, fileImports)
	case *ast.StructType:
		readAstStructType(v, exprType, fileImports)
	case *ast.InterfaceType:
		readAstInterfaceType(v, exprType, fileImports)
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

func readFieldList(fieldList []*ast.Field, fileImports fileImportSpecMap) (fields []*Field, inlineFields []*Field, unionFields []*Field) {
	fields = []*Field{}
	for _, field := range fieldList {
		if names, value, jsonInfo := readField(field, fileImports); value != nil {
			for _, name := range names {
				if len(name) == 0 {
					if jsonInfo == nil {
						trace("i do not understand this one", field, names, value, jsonInfo)
						continue
					} else if jsonInfo.Ignore {
						trace("ignoring this one", field, names, value, jsonInfo)
						continue
					} else if jsonInfo.Inline {
						inlineFields = append(inlineFields, &Field{
							Name:     name,
							Value:    value,
							JSONInfo: jsonInfo,
						})
						continue
					}
				} else if strings.Compare(strings.ToLower(name[:1]), name[:1]) == 0 {
					// this is not unicode proof
					continue
				} else if jsonInfo != nil && jsonInfo.Union {
					unionFields = append(unionFields, &Field{
						Name:     name,
						Value:    value,
						JSONInfo: jsonInfo,
					})
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
		if funcDecl, ok := d.(*ast.FuncDecl); ok {
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) == 1 {
				firstReceiverField := funcDecl.Recv.List[0]
				if starExpr, ok := firstReceiverField.Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok {
						if funcDecl.Name.Name == "Error" && funcDecl.Type.Params.NumFields() == 0 && funcDecl.Type.Results.NumFields() == 1 {
							returnValueField := funcDecl.Type.Results.List[0]
							if returnValueIdent, ok := returnValueField.Type.(*ast.Ident); ok {
								if returnValueIdent.Name == "string" {
									errorTypes[packageName+"."+ident.Name] = true
								}
								// fmt.Println("error for:", ident.Name, returnValueIdent.Name)
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

			if typeSpec, ok := obj.Decl.(*ast.TypeSpec); ok {
				switch typeSpecType := typeSpec.Type.(type) {
				case *ast.StructType:
					structs[structName] = &Struct{
						Name:    name,
						Fields:  []*Field{},
						Package: packageName,
					}
					trace("StructType", obj.Name)
					fields, inlineFields, unionFields := readFieldList(typeSpecType.Fields.List, fileImports)
					structs[structName].Fields = fields
					structs[structName].InlineFields = inlineFields
					structs[structName].UnionFields = unionFields
				case *ast.InterfaceType:
					trace("Interface", obj.Name)
					scalars[structName] = &Scalar{
						Name:    structName,
						Package: packageName,
						Type:    ScalarTypeAny,
					}
				case *ast.Ident:
					trace("Scalar", obj.Name)
					scalars[structName] = &Scalar{
						Name:    structName,
						Package: packageName,
						Type:    getScalarFromAstIdent(typeSpecType),
					}
				case *ast.SelectorExpr:
					trace("SelectorExpr", obj.Name)
					structs[structName] = &Struct{
						Name:    name,
						Package: packageName,
					}
				case *ast.ArrayType:
					arrayValue := &Value{}
					arrayValue.loadExpr(typeSpec.Type, fileImports)
					structs[structName] = &Struct{
						Name:    name,
						Package: packageName,
						Array:   arrayValue.Array,
					}
				case *ast.MapType:
					mapValue := &Value{}
					mapValue.loadExpr(typeSpec.Type, fileImports)
					structs[structName] = &Struct{
						Name:    name,
						Package: packageName,
						Map:     mapValue.Map,
					}
				default:
					fmt.Println("	ignoring", obj.Name, reflect.ValueOf(typeSpec.Type).Type().String())
				}
			}
		}
	}
	return nil
}
