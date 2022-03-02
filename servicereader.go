package gotsrpc

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"sort"
	"strings"

	"github.com/foomo/gotsrpc/v2/config"
)

func (sl ServiceList) Len() int           { return len(sl) }
func (sl ServiceList) Swap(i, j int)      { sl[i], sl[j] = sl[j], sl[i] }
func (sl ServiceList) Less(i, j int) bool { return strings.Compare(sl[i].Name, sl[j].Name) > 0 }

func readServiceFile(file *ast.File, packageName string, services ServiceList) error {
	findService := func(serviceName string) (service *Service, ok bool) {
		for _, service := range services {
			if service.Name == serviceName {
				return service, true
			}
		}
		return nil, false
	}

	fileImports := getFileImports(file, packageName)

	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Recv != nil {
				trace("that is a method named", funcDecl.Name)
				if len(funcDecl.Recv.List) == 1 {

					firstReceiverField := funcDecl.Recv.List[0]
					if "*ast.StarExpr" == reflect.ValueOf(firstReceiverField.Type).Type().String() {
						starExpr := firstReceiverField.Type.(*ast.StarExpr)
						if "*ast.Ident" == reflect.ValueOf(starExpr.X).Type().String() {

							ident := starExpr.X.(*ast.Ident)
							service, ok := findService(ident.Name)
							firstCharOfMethodName := funcDecl.Name.Name[0:1]
							if !ok || strings.ToLower(firstCharOfMethodName) == firstCharOfMethodName {
								// skip this method
								continue
							}

							trace("	on sth:", ident.Name)

							service.Methods = append(service.Methods, &Method{
								Name:   funcDecl.Name.Name,
								Args:   readFields(funcDecl.Type.Params, fileImports),
								Return: readFields(funcDecl.Type.Results, fileImports),
							})
						}
					}
				}
			} else {
				trace("no receiver for", funcDecl.Name)
			}
		} else if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					ident := typeSpec.Name
					trace("that is an interface named", ident.Name)
					if service, ok := findService(ident.Name); ok {

						if iSpec, ok := typeSpec.Type.(*ast.InterfaceType); ok {
							service.IsInterface = true
							for _, fieldDecl := range iSpec.Methods.List {
								if funcDecl, ok := fieldDecl.Type.(*ast.FuncType); ok {
									if len(fieldDecl.Names) == 0 {
										continue
									}
									mname := fieldDecl.Names[0]
									trace(" on sth:", mname.Name)
									//fmt.Println("interface:", ident.Name, "method:", mname.Name)
									service.Methods = append(service.Methods, &Method{
										Name:   mname.Name,
										Args:   readFields(funcDecl.Params, fileImports),
										Return: readFields(funcDecl.Results, fileImports),
									})
								}
							}
						}
					}
				}
			}
		}
	}
	for _, s := range services {
		sort.Sort(s.Methods)
	}
	return nil
}

type importSpec struct {
	alias string
	name  string
	path  string
}

type fileImportSpecMap map[string]importSpec

func (fileImports fileImportSpecMap) getPackagePath(packageName string) string {
	importSpec, ok := fileImports[packageName]
	if ok {
		packageName = importSpec.path
	}
	return packageName
}

func standardImportName(importPath string) string {
	pathParts := strings.Split(importPath, "/")
	return pathParts[len(pathParts)-1]
}

func getFileImports(file *ast.File, packageName string) (imports fileImportSpecMap) {
	imports = fileImportSpecMap{"": importSpec{alias: "", name: "", path: packageName}}
	for _, decl := range file.Decls {
		if reflect.ValueOf(decl).Type().String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			if genDecl.Tok == token.IMPORT {
				trace("got an import", genDecl.Specs)
				for _, spec := range genDecl.Specs {
					if "*ast.ImportSpec" == reflect.ValueOf(spec).Type().String() {
						spec := spec.(*ast.ImportSpec)
						importPath := spec.Path.Value[1 : len(spec.Path.Value)-1]
						importName := spec.Name.String()
						if importName == "" || importName == "<nil>" {
							importName = standardImportName(importPath)
						}
						imports[importName] = importSpec{
							alias: importName,
							name:  standardImportName(importPath),
							path:  importPath,
						}
						//trace("  import   >>>>>>>>>>>>>>>>>>>>", importName, importPath)
					}
				}
			}
		}
	}
	return imports
}

func readFields(fieldList *ast.FieldList, fileImports fileImportSpecMap) (fields []*Field) {
	trace("reading fields")
	fields = []*Field{}
	if fieldList == nil {
		return
	}

	for _, param := range fieldList.List {
		names, value, _ := readField(param, fileImports)
		for _, name := range names {
			fields = append(fields, &Field{
				Name:  name,
				Value: value,
			})
		}
	}
	trace("done reading fields")
	return
}

func readServicesInPackage(pkg *ast.Package, packageName string, serviceMap map[string]string) (services ServiceList, err error) {
	if pkg == nil {
		return nil, errors.New("package cannot be nil")
	}
	services = ServiceList{}
	for endpoint, serviceName := range serviceMap {
		services = append(services, &Service{
			Name:     serviceName,
			Methods:  []*Method{},
			Endpoint: endpoint,
		})
	}
	pkgFiles := make([]string, 0, len(pkg.Files))
	for k := range pkg.Files {
		pkgFiles = append(pkgFiles, k)
	}
	sort.Strings(pkgFiles)

	for _, k := range pkgFiles {
		file := pkg.Files[k]
		err = readServiceFile(file, packageName, services)
		if err != nil {
			return
		}

	}
	sort.Sort(services)
	return
}

func loadConstantTypes(pkg *ast.Package) map[string]interface{} {
	constantTypes := map[string]interface{}{}
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if reflect.ValueOf(decl).Type().String() == "*ast.GenDecl" {
				genDecl := decl.(*ast.GenDecl)
				switch genDecl.Tok {
				case token.TYPE:
					trace("got a type", genDecl.Specs)
					for _, spec := range genDecl.Specs {
						if reflect.ValueOf(spec).Type().String() == "*ast.TypeSpec" {
							spec := spec.(*ast.TypeSpec)
							if _, ok := constantTypes[spec.Name.Name]; ok {
								continue
							}
							switch reflect.ValueOf(spec.Type).Type().String() {
							case "*ast.InterfaceType":
								constantTypes[spec.Name.Name] = "any"
							case "*ast.Ident":
								specIdent := spec.Type.(*ast.Ident)
								switch specIdent.Name {
								case "byte":
									constantTypes[spec.Name.Name] = "any"
								case "string":
									constantTypes[spec.Name.Name] = "string"
								case "bool":
									constantTypes[spec.Name.Name] = "boolean"
								case "float", "float32", "float64",
									"int", "int8", "int16", "int32", "int64",
									"uint", "uint8", "uint16", "uint32", "uint64":
									constantTypes[spec.Name.Name] = "number"
								default:
									trace("unhandled type", reflect.ValueOf(spec.Type).Type().String())
								}
							default:
								trace("ignoring type", reflect.ValueOf(spec.Type).Type().String())
							}
						}
					}
				case token.CONST:
					trace("got a const", genDecl.Specs)
					for _, spec := range genDecl.Specs {
						if reflect.ValueOf(spec).Type().String() == "*ast.ValueSpec" {
							spec := spec.(*ast.ValueSpec)
							if specType, ok := spec.Type.(*ast.Ident); ok {
								for _, val := range spec.Values {
									if reflect.ValueOf(val).Type().String() == "*ast.BasicLit" {
										if _, ok := constantTypes[specType.Name]; !ok {
											constantTypes[specType.Name] = map[string]*ast.BasicLit{}
										} else if _, ok := constantTypes[specType.Name].(map[string]*ast.BasicLit); !ok {
											constantTypes[specType.Name] = map[string]*ast.BasicLit{}
										}
										constantTypes[specType.Name].(map[string]*ast.BasicLit)[spec.Names[0].Name] = val.(*ast.BasicLit)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return constantTypes
}

func Read(
	goPaths []string,
	gomod config.Namespace,
	packageName string,
	serviceMap map[string]string,
	missingTypes map[string]bool,
	missingConstants map[string]bool,
) (
	pkgName string,
	services ServiceList,
	structs map[string]*Struct,
	scalars map[string]*Scalar,
	constantTypes map[string]map[string]interface{},
	err error,
) {
	if len(serviceMap) == 0 {
		err = errors.New("nothing to do service names are empty")
		return
	}
	pkg, parseErr := parsePackage(goPaths, gomod, packageName)
	if parseErr != nil {
		err = parseErr
		return
	}
	pkgName = pkg.Name
	services, err = readServicesInPackage(pkg, packageName, serviceMap)
	if err != nil {
		return
	}

	for _, s := range services {
		for _, m := range s.Methods {
			collectStructTypes(m.Return, missingTypes)
			collectStructTypes(m.Args, missingTypes)
			collectScalarTypes(m.Return, missingTypes)
			collectScalarTypes(m.Args, missingTypes)
		}
	}
	trace("missing")
	traceData(missingTypes)

	structs = map[string]*Struct{}
	scalars = map[string]*Scalar{}

	collectErr := collectTypes(goPaths, gomod, missingTypes, structs, scalars)
	if collectErr != nil {
		err = errors.New("error while collecting structs: " + collectErr.Error())
	}
	trace("---------------- found structs -------------------")
	traceData(structs)
	trace("---------------- /found structs -------------------")
	trace("---------------- found scalars -------------------")
	traceData(scalars)
	trace("---------------- /found scalars -------------------")
	allConstantTypes := map[string]map[string]interface{}{}
	for _, structDef := range structs {
		if structDef != nil {
			structPackage := structDef.Package
			if _, ok := allConstantTypes[structPackage]; !ok {
				if pkg, constPkgErr := parsePackage(goPaths, gomod, structPackage); constPkgErr != nil {
					err = constPkgErr
					return
				} else {
					allConstantTypes[structPackage] = loadConstantTypes(pkg)
				}
			}
		}
	}
	for _, scalarDef := range scalars {
		if scalarDef != nil {
			scalarPackage := scalarDef.Package
			if _, ok := allConstantTypes[scalarPackage]; !ok {
				if pkg, constPkgErr := parsePackage(goPaths, gomod, scalarPackage); constPkgErr != nil {
					err = constPkgErr
					return
				} else {
					allConstantTypes[scalarPackage] = loadConstantTypes(pkg)
				}
			}
		}
	}

	flatStructs := map[string]bool{}
	for _, s := range structs {
		loadFlatStructs(s, flatStructs)
	}

	constantTypes = map[string]map[string]interface{}{}
	for constantTypePackage, constantType := range allConstantTypes {
		for constantTypeName, constantTypeVales := range constantType {
			fullName := constantTypePackage + "." + constantTypeName
			_, scalarOK := scalars[fullName]
			_, structOK := flatStructs[fullName]
			_, constantsOK := missingConstants[fullName]

			if scalarOK || structOK || constantsOK {
				missingConstants[fullName] = false
				if _, ok := constantTypes[constantTypePackage]; !ok {
					constantTypes[constantTypePackage] = map[string]interface{}{}
				}
				constantTypes[constantTypePackage][constantTypeName] = constantTypeVales
			}
		}
	}

	for missingConstant, missing := range missingConstants {
		if missing {
			err = errors.New("could not resolve constant: " + missingConstant)
			return
		}
	}

	// fix arg and return field lists
	for _, service := range services {
		for _, method := range service.Methods {
			fixFieldStructs(method.Args, structs, scalars)
			fixFieldStructs(method.Return, structs, scalars)
		}
	}
	traceData("---------------------------", services)
	return
}

func loadFlatStructs(s *Struct, flatStructs map[string]bool) {
	if s.Map != nil {
		if s.Map.Key != nil {
			loadFlatStructsValue(s.Map.Key, flatStructs)
		}
		if s.Map.Value != nil && s.Map.Value.Scalar != nil {
			loadFlatStructsValue(s.Map.Value, flatStructs)
		}
	}
	if s.Fields != nil {
		for _, field := range s.Fields {
			loadFlatStructsValue(field.Value, flatStructs)
		}
	}
	flatStructs[s.FullName()] = true
}

func loadFlatStructsValue(s *Value, flatStructs map[string]bool) {
	if s.Map != nil {
		if s.Map.Key != nil {
			loadFlatStructsValue(s.Map.Key, flatStructs)
		}
		if s.Map.Value != nil && s.Map.Value.Scalar != nil {
			loadFlatStructsValue(s.Map.Value, flatStructs)
		}
	}
	if s.Struct != nil {
		loadFlatStructs(s.Struct, flatStructs)
	}
	if s.Scalar != nil {
		flatStructs[s.Scalar.FullName()] = true
	}
}
func fixFieldStructs(fields []*Field, structs map[string]*Struct, scalars map[string]*Scalar) {
	for _, f := range fields {
		if f.Value.StructType != nil {
			// do we have that struct or is it a hidden scalar
			name := f.Value.StructType.FullName()
			s, strctExists := structs[name]
			if strctExists {
				f.Value.IsError = s.IsError
				continue
			}
			scalar, scalarExists := scalars[name]
			if scalarExists {
				f.Value.StructType = nil
				f.Value.Scalar = scalar
			}
		}
	}
}

func collectTypes(goPaths []string, gomod config.Namespace, missingTypes map[string]bool, structs map[string]*Struct, scalars map[string]*Scalar) error {
	scannedPackageStructs := map[string]map[string]*Struct{}
	scannedPackageScalars := map[string]map[string]*Scalar{}
	missingTypeNames := func() []string {
		var missing []string
		for name, isMissing := range missingTypes {
			if isMissing {
				missing = append(missing, name)
			}
		}
		// fmt.Println("missing types", len(missingTypes), "missing", len(missing))
		return missing
	}
	lastNumMissing := len(missingTypeNames())

	for typesPending(structs, scalars, missingTypes) {
		trace("pending", missingTypeNames())
		for fullName, typeIsMissing := range missingTypes {
			if !typeIsMissing {
				continue
			}
			fullNameParts := strings.Split(fullName, ".")
			fullNameParts = fullNameParts[:len(fullNameParts)-1]

			//path := fullNameParts[:len(fullNameParts)-1][0]

			packageName := strings.Join(fullNameParts, ".")

			trace(fullName, "==========================>", fullNameParts, "=============>", packageName)

			packageStructs, structOK := scannedPackageStructs[packageName]
			packageScalars, scalarOK := scannedPackageScalars[packageName]
			if !structOK || !scalarOK {
				parsedPackageStructs, parsedPackageScalars, err := getTypesInPackage(goPaths, gomod, packageName)
				if err != nil {
					return err
				}

				trace("found structs in", goPaths, packageName)
				for structName, strct := range packageStructs {
					trace("	struct", structName, strct)
					if strct == nil {
						panic("how could that be")
					}
				}
				trace("found scalars in", goPaths, packageName)
				for scalarName, scalar := range packageScalars {
					trace("	scalar", scalarName, scalar)
				}
				traceData(parsedPackageScalars)
				packageStructs = parsedPackageStructs
				packageScalars = parsedPackageScalars
				scannedPackageStructs[packageName] = packageStructs
				scannedPackageScalars[packageName] = packageScalars
			}
			traceData("packageStructs", packageName, packageStructs)
			for packageStructName, packageStruct := range packageStructs {
				missing, needed := missingTypes[packageStructName]
				if needed && missing {
					trace("picked up package struct", packageStructName, packageStruct)
					missingTypes[packageStructName] = false
					if packageStruct == nil {
						panic("waaaaaaaaa")
					}
					structs[packageStructName] = packageStruct
				}
			}

			traceData("packageScalars", packageScalars)
			for packageScalarName, packageScalar := range packageScalars {
				missing, needed := missingTypes[packageScalarName]
				if needed && missing {
					trace("picked up package scalar", packageScalarName, packageScalar)
					missingTypes[packageScalarName] = false
					scalars[packageScalarName] = packageScalar
				}
			}

		}
		newNumMissingTypes := len(missingTypeNames())
		if newNumMissingTypes > 0 && newNumMissingTypes == lastNumMissing {
			//packageStructs, structOK := scannedPackageStructs[packageName]
			for scalarName, scalars := range scannedPackageScalars {
				fmt.Println("scanned scalars ", scalarName)
				for _, scalar := range scalars {
					fmt.Println("	", scalar.Name)
				}
			}
			for structName, strcts := range scannedPackageStructs {
				fmt.Println("scanned struct ", structName)
				for _, strct := range strcts {
					fmt.Println("	", strct.Name)
				}
			}
			return errors.New(fmt.Sprintln("could not resolve at least one of the following types", missingTypeNames()))
		}
		lastNumMissing = newNumMissingTypes
	}
	return nil
}

func typesPending(structs map[string]*Struct, scalars map[string]*Scalar, missingTypes map[string]bool) bool {
	for _, missing := range missingTypes {
		if missing {
			return true
		}
	}
	for _, structType := range structs {
		if !structType.DepsSatisfied(missingTypes, structs, scalars) {
			return true
		}
	}
	return false
}

func (s *Struct) DepsSatisfied(missingTypes map[string]bool, structs map[string]*Struct, scalars map[string]*Scalar) bool {
	needsWork := func(fullName string) bool {
		strct, strctOK := structs[fullName]
		scalar, scalarOK := scalars[fullName]
		if !strctOK && !scalarOK {
			// hey there is more todo
			missingTypes[fullName] = true
			trace("need work ----------------------" + fullName)
			return true
		}
		if strct == nil && scalar == nil {
			trace("need work ----------------------" + fullName)
			return true
		}
		return false
	}
	needWorksFields := func(fields []*Field) bool {
		for _, field := range fields {
			var fieldStructType *StructType = nil
			if field.Value.StructType != nil {
				fieldStructType = field.Value.StructType
			} else if field.Value.Array != nil && field.Value.Array.Value.StructType != nil {
				fieldStructType = field.Value.Array.Value.StructType
			} else if field.Value.Map != nil && field.Value.Map.Value.StructType != nil {
				fieldStructType = field.Value.Map.Value.StructType
			} else if field.Value.Scalar != nil && needsWork(field.Value.Scalar.FullName()) {
				return false
			} else if field.Value.Array != nil && field.Value.Array.Value.Scalar != nil && needsWork(field.Value.Array.Value.Scalar.FullName()) {
				return false
			} else if field.Value.Map != nil && field.Value.Map.Value.Scalar != nil && needsWork(field.Value.Map.Value.Scalar.FullName()) {
				return false
			}
			if fieldStructType != nil {
				if needsWork(fieldStructType.FullName()) {
					return false
				}
			}
		}
		return true
	}
	if ok := needWorksFields(s.Fields); !ok {
		return false
	} else if ok := needWorksFields(s.InlineFields); !ok {
		return false
	} else if ok := needWorksFields(s.UnionFields); !ok {
		return false
	}
	//// special handling of union only structs
	//if len(s.Fields) == 0 {
	//	for _, field := range s.UnionFields {
	//		var fieldStructType *StructType = nil
	//		if field.Value.StructType != nil {
	//			fieldStructType = field.Value.StructType
	//		} else if field.Value.Scalar != nil && needsWork(field.Value.Scalar.FullName()) {
	//			return false
	//		}
	//		if fieldStructType != nil {
	//			if needsWork(fieldStructType.FullName()) {
	//				return false
	//			}
	//		}
	//	}
	//}
	if s.Array != nil && s.Array.Value != nil {
		if s.Array.Value.StructType != nil {
			if needsWork(s.Array.Value.StructType.FullName()) {
				return false
			}
		} else if s.Array.Value.Scalar != nil {
			if needsWork(s.Array.Value.Scalar.FullName()) {
				return false
			}
		}
	}
	if s.Map != nil && s.Map.Value != nil {
		if s.Map.Value.StructType != nil {
			if needsWork(s.Map.Value.StructType.FullName()) {
				return false
			}
		} else if s.Map.Value.Scalar != nil {
			if needsWork(s.Map.Value.Scalar.FullName()) {
				return false
			}
		}
	}
	return !needsWork(s.FullName())
}

func (s *Struct) FullName() string {
	fullName := s.Package + "." + s.Name
	if len(fullName) == 0 {
		fullName = s.Name
	}
	return fullName
}

func (st *StructType) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}
	return fullName
}

func getTypesInPackage(
	goPaths []string,
	gomod config.Namespace,
	packageName string,
) (
	structs map[string]*Struct,
	scalars map[string]*Scalar,
	err error,
) {
	pkg, err := parsePackage(goPaths, gomod, packageName)
	if err != nil {
		return nil, nil, err
	}
	structs, scalars, err = readStructs(pkg, packageName)
	if err != nil {
		return nil, nil, err
	}
	return structs, scalars, nil
}

func getStructTypeForField(value *Value) *StructType {
	//field.Value.StructType
	var strType *StructType
	switch true {
	case value.StructType != nil:
		strType = value.StructType
		//case field.Value.ArrayType
	case value.Map != nil:
		strType = getStructTypeForField(value.Map.Value)
	case value.Array != nil:
		strType = getStructTypeForField(value.Array.Value)
	}
	return strType
}

func getScalarForField(value *Value) []*Scalar {
	//field.Value.StructType
	var scalarTypes []*Scalar
	switch true {
	case value.Scalar != nil:
		scalarTypes = append(scalarTypes, value.Scalar)
		//case field.Value.ArrayType
	case value.Map != nil:
		if value.Map.Key != nil {
			if v := getScalarForField(value.Map.Key); v != nil {
				scalarTypes = append(scalarTypes, v...)
			}
		}
		scalarTypes = append(scalarTypes, getScalarForField(value.Map.Value)...)
	case value.Array != nil:
		scalarTypes = append(scalarTypes, getScalarForField(value.Array.Value)...)
	}
	return scalarTypes
}

func collectScalarTypes(fields []*Field, scalarTypes map[string]bool) {
	for _, field := range fields {
		for _, scalarType := range getScalarForField(field.Value) {
			if scalarType != nil {
				fullName := scalarType.Package + "." + scalarType.Name
				if len(scalarType.Package) == 0 {
					fullName = scalarType.Name
				}
				switch fullName {
				case "error", "net/http.Request", "net/http.ResponseWriter":
					continue
				default:
					scalarTypes[fullName] = true
				}
			}
		}
	}
}

func collectStructTypes(fields []*Field, structTypes map[string]bool) {
	for _, field := range fields {
		strType := getStructTypeForField(field.Value)
		if strType != nil {
			fullName := strType.Package + "." + strType.Name
			if len(strType.Package) == 0 {
				fullName = strType.Name
			}
			switch fullName {
			case "error", "net/http.Request", "net/http.ResponseWriter":
				continue
			default:
				structTypes[fullName] = true
			}
		}
	}
}

//func collectStructs(goPath, structs)
