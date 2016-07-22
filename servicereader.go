package gotsrpc

import (
	"errors"
	"go/ast"
	"go/token"
	"reflect"
	"runtime"
	"strings"
)

func readServiceFile(file *ast.File, packageName string, services []*Service) error {
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
							trace("	on sth:", ident.Name)

							service, ok := findService(ident.Name)
							firstCharOfMethodName := funcDecl.Name.Name[0:1]
							if ok && strings.ToLower(firstCharOfMethodName) != firstCharOfMethodName {
								service.Methods = append(service.Methods, &Method{
									Name:   funcDecl.Name.Name,
									Args:   readFields(funcDecl.Type.Params, fileImports),
									Return: readFields(funcDecl.Type.Results, fileImports),
								})
							}
						}
					}
				}
			} else {
				trace("no receiver for", funcDecl.Name)
			}
		}
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
	return

}

func readServicesInPackage(pkg *ast.Package, packageName string, serviceNames []string) (services []*Service, err error) {
	services = []*Service{}
	for _, serviceName := range serviceNames {
		services = append(services, &Service{
			Name:    serviceName,
			Methods: []*Method{},
		})
	}
	for _, file := range pkg.Files {
		err = readServiceFile(file, packageName, services)
		if err != nil {
			return
		}

	}
	return
}

func Read(goPath string, packageName string, serviceNames []string) (services []*Service, structs map[string]*Struct, err error) {
	if len(serviceNames) == 0 {
		err = errors.New("nothing to do service names are empty")
		return
	}
	pkg, err := parsePackage(goPath, packageName)
	if err != nil {
		return
	}
	services, err = readServicesInPackage(pkg, packageName, serviceNames)
	if err != nil {
		return
	}

	structTypes := map[string]*StructType{}
	for _, s := range services {
		for _, m := range s.Methods {
			collecStructTypes(m.Return, structTypes)
			collecStructTypes(m.Args, structTypes)
		}
	}
	structs = map[string]*Struct{}
	for wantedName := range structTypes {
		structs[wantedName] = nil
	}
	collectErr := collectStructs(goPath, structs)
	if collectErr != nil {

		err = errors.New("error while collecting structs: " + collectErr.Error())
	}
	return
}

func collectStructs(goPath string, structs map[string]*Struct) error {
	scannedPackages := map[string]map[string]*Struct{}
	for structsPending(structs) {
		trace("pending", len(structs))
		for fullName, strct := range structs {
			if strct != nil {
				continue
			}
			fullNameParts := strings.Split(fullName, ".")
			fullNameParts = fullNameParts[:len(fullNameParts)-1]

			//path := fullNameParts[:len(fullNameParts)-1][0]

			packageName := strings.Join(fullNameParts, ".")

			// trace(fullName, "==========================>", fullNameParts, "=============>", packageName)

			packageStructs, ok := scannedPackages[packageName]
			if !ok {
				parsedPackageStructs, err := getStructsInPackage(goPath, packageName)
				if err != nil {
					return err
				}
				packageStructs = parsedPackageStructs
				scannedPackages[packageName] = packageStructs
			}
			for packageStructName, packageStruct := range packageStructs {
				// trace("------------------------------------>", packageStructName, packageStruct)
				existingStruct, needed := structs[packageStructName]
				if needed && existingStruct == nil {
					structs[packageStructName] = packageStruct
				}
			}
		}
	}
	return nil
}

func structsPending(structs map[string]*Struct) bool {
	for name, structType := range structs {
		if structType == nil {
			trace("missing", name)
			return true
		}
		if !structType.DepsSatisfied(structs) {
			return true
		}
	}
	return false
}

func (s *Struct) DepsSatisfied(structs map[string]*Struct) bool {
	needsWork := func(fullName string) bool {
		strct, ok := structs[fullName]
		if !ok {
			// hey there is more todo
			structs[fullName] = nil
			trace("need work ----------------------" + fullName)
			return true
		}
		if strct == nil {
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

func getStructsInPackage(goPath string, packageName string) (structs map[string]*Struct, err error) {
	pkg, err := parsePackage(goPath, packageName)
	if err != nil {
		pkg, err = parsePackage(runtime.GOROOT(), packageName)
		if err != nil {
			return nil, err
		}
	}
	structs, err = readStructs(pkg, packageName)
	if err != nil {
		return nil, err
	}
	return structs, nil
}

func collecStructTypes(fields []*Field, structTypes map[string]*StructType) {
	for _, field := range fields {
		if field.Value.StructType != nil {
			fullName := field.Value.StructType.Package + "." + field.Value.StructType.Name
			if len(field.Value.StructType.Package) == 0 {
				fullName = field.Value.StructType.Name
			}
			switch fullName {
			case "error", "net/http.Request", "net/http.ResponseWriter":
				continue
			default:
				structTypes[fullName] = field.Value.StructType
			}

		}
	}
}

//func collectStructs(goPath, structs)
