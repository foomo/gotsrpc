package gotsrpc

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"sort"
	"strings"
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
		if reflect.ValueOf(decl).Type().String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
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
		name, value, _ := readField(param, fileImports)
		fields = append(fields, &Field{
			Name:  name,
			Value: value,
		})
	}
	trace("done reading fields")
	return

}

func readServicesInPackage(pkg *ast.Package, packageName string, serviceMap map[string]string) (services ServiceList, err error) {
	services = ServiceList{}
	for endpoint, serviceName := range serviceMap {
		services = append(services, &Service{
			Name:     serviceName,
			Methods:  []*Method{},
			Endpoint: endpoint,
		})
	}
	for _, file := range pkg.Files {
		err = readServiceFile(file, packageName, services)
		if err != nil {
			return
		}

	}
	sort.Sort(services)
	return
}

func loadConstants(pkg *ast.Package) map[string]*ast.BasicLit {
	constants := map[string]*ast.BasicLit{}
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if reflect.ValueOf(decl).Type().String() == "*ast.GenDecl" {
				genDecl := decl.(*ast.GenDecl)
				if genDecl.Tok == token.CONST {
					trace("got a const", genDecl.Specs)
					for _, spec := range genDecl.Specs {
						if "*ast.ValueSpec" == reflect.ValueOf(spec).Type().String() {
							spec := spec.(*ast.ValueSpec)
							for _, val := range spec.Values {
								if reflect.ValueOf(val).Type().String() == "*ast.BasicLit" {
									firstValueLit := val.(*ast.BasicLit)
									constName := spec.Names[0].String()
									for indexRune, r := range constName {
										if indexRune == 0 {
											if string(r) == strings.ToUpper(string(r)) {
												constants[constName] = firstValueLit
											}
											break
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return constants

}

func Read(
	goPaths []string,
	packageName string,
	serviceMap map[string]string,
) (
	services ServiceList,
	structs map[string]*Struct,
	scalars map[string]*Scalar,
	constants map[string]map[string]*ast.BasicLit,
	err error,
) {
	if len(serviceMap) == 0 {
		err = errors.New("nothing to do service names are empty")
		return
	}
	pkg, parseErr := parsePackage(goPaths, packageName)
	if parseErr != nil {
		err = parseErr
		return
	}
	services, err = readServicesInPackage(pkg, packageName, serviceMap)
	if err != nil {
		return
	}

	missingTypes := map[string]bool{}
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

	collectErr := collectTypes(goPaths, missingTypes, structs, scalars)
	if collectErr != nil {
		err = errors.New("error while collecting structs: " + collectErr.Error())
	}
	trace("---------------- found structs -------------------")
	traceData(structs)
	trace("---------------- /found structs -------------------")
	trace("---------------- found scalars -------------------")
	traceData(scalars)
	trace("---------------- /found scalars -------------------")
	constants = map[string]map[string]*ast.BasicLit{}
	for _, structDef := range structs {
		if structDef != nil {
			structPackage := structDef.Package
			_, ok := constants[structPackage]
			if !ok {
				// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", structPackage)
				pkg, constPkgErr := parsePackage(goPaths, structPackage)
				if constPkgErr != nil {
					err = constPkgErr
					return
				}
				constants[structPackage] = loadConstants(pkg)

			}
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

func collectTypes(goPaths []string, missingTypes map[string]bool, structs map[string]*Struct, scalars map[string]*Scalar) error {
	scannedPackageStructs := map[string]map[string]*Struct{}
	scannedPackageScalars := map[string]map[string]*Scalar{}
	missingTypeNames := func() []string {
		missing := []string{}
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
				parsedPackageStructs, parsedPackageScalars, err := getTypesInPackage(goPaths, packageName)
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
				for scalarName, scalar := range parsedPackageScalars {
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

func (s *Struct) DepsSatisfied(missingTypes map[string]bool, structs map[string]*Struct, scalarTypes map[string]*Scalar) bool {
	needsWork := func(fullName string) bool {
		strct, strctOK := structs[fullName]
		scalar, scalarOK := scalarTypes[fullName]
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
	for _, field := range s.Fields {
		var fieldStructType *StructType
		fieldStructType = nil
		if field.Value.StructType != nil {
			fieldStructType = field.Value.StructType
		} else if field.Value.Array != nil && field.Value.Array.Value.StructType != nil {
			fieldStructType = field.Value.Array.Value.StructType
		} else if field.Value.Map != nil && field.Value.Map.Value.StructType != nil {
			fieldStructType = field.Value.Map.Value.StructType
		}
		if fieldStructType != nil {
			if needsWork(fieldStructType.FullName()) {
				return false
			}
		}
	}
	if s.Array != nil {
		if s.Array.Value != nil {
			if s.Array.Value.StructType != nil {
				if needsWork(s.Array.Value.StructType.FullName()) {
					return false
				}
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
	packageName string,
) (
	structs map[string]*Struct,
	scalars map[string]*Scalar,
	err error,
) {
	pkg, err := parsePackage(goPaths, packageName)
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

func getScalarForField(value *Value) *Scalar {
	//field.Value.StructType
	var scalarType *Scalar
	switch true {
	case value.Scalar != nil:
		scalarType = value.Scalar
	//case field.Value.ArrayType
	case value.Map != nil:
		scalarType = getScalarForField(value.Map.Value)
	case value.Array != nil:
		scalarType = getScalarForField(value.Array.Value)
	}
	return scalarType
}

func collectScalarTypes(fields []*Field, scalarTypes map[string]bool) {
	for _, field := range fields {

		scalarType := getScalarForField(field.Value)
		if scalarType != nil {
			fullName := scalarType.Package + "." + scalarType.Name
			if len(scalarType.Package) == 0 {
				fullName = scalarType.Name
			}
			scalarTypes[fullName] = true
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
