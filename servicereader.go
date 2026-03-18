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

// interfaceInfo holds a parsed interface type and its file imports.
type interfaceInfo struct {
	iface      *ast.InterfaceType
	imports    fileImportSpecMap
	typeParams []string
}

// collectPackageInterfaces scans all files in the package and builds a map
// of interface names to their AST and file imports.
func collectPackageInterfaces(pkg *ast.Package, packageName string) map[string]interfaceInfo {
	result := map[string]interfaceInfo{}
	for _, file := range pkg.Files {
		fileImports := getFileImports(file, packageName)
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				iface, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}
				var typeParams []string
				if typeSpec.TypeParams != nil {
					for _, tp := range typeSpec.TypeParams.List {
						for _, n := range tp.Names {
							typeParams = append(typeParams, n.Name)
						}
					}
				}
				result[typeSpec.Name.Name] = interfaceInfo{
					iface:      iface,
					imports:    fileImports,
					typeParams: typeParams,
				}
			}
		}
	}
	return result
}

// resolvedMethod pairs an AST function type with the file imports it was declared in.
type resolvedMethod struct {
	name      string
	funcTyp   *ast.FuncType
	imports   fileImportSpecMap
	typeSubst map[string]ast.Expr
}

// resolveExpr resolves an AST expression through a type substitution map.
func resolveExpr(expr ast.Expr, typeSubst map[string]ast.Expr) ast.Expr {
	if ident, ok := expr.(*ast.Ident); ok {
		if sub, ok := typeSubst[ident.Name]; ok {
			return sub
		}
	}
	return expr
}

// resolveInterfaceMethods recursively collects all methods from an interface,
// following embedded interfaces via the pkgInterfaces map. Uses visited for cycle protection.
// typeSubst maps type parameter names to concrete type expressions.
func resolveInterfaceMethods(iface *ast.InterfaceType, imports fileImportSpecMap, pkgInterfaces map[string]interfaceInfo, visited map[string]bool, typeSubst map[string]ast.Expr) []resolvedMethod {
	var methods []resolvedMethod
	for _, field := range iface.Methods.List {
		switch ft := field.Type.(type) {
		case *ast.FuncType:
			if len(field.Names) == 0 {
				continue
			}
			methods = append(methods, resolvedMethod{
				name:      field.Names[0].Name,
				funcTyp:   ft,
				imports:   imports,
				typeSubst: typeSubst,
			})
		case *ast.Ident:
			// Embedded interface reference (non-generic)
			if visited[ft.Name] {
				continue
			}
			visited[ft.Name] = true
			if info, ok := pkgInterfaces[ft.Name]; ok {
				methods = append(methods, resolveInterfaceMethods(info.iface, info.imports, pkgInterfaces, visited, nil)...)
			}
		case *ast.IndexExpr:
			// Generic embedded interface with single type arg: Base[string] or Base[T]
			ident, ok := ft.X.(*ast.Ident)
			if !ok {
				continue
			}
			if visited[ident.Name] {
				continue
			}
			visited[ident.Name] = true
			info, ok := pkgInterfaces[ident.Name]
			if !ok {
				continue
			}
			// Build substitution map for the embedded interface's type params
			newSubst := map[string]ast.Expr{}
			resolvedArg := resolveExpr(ft.Index, typeSubst)
			if len(info.typeParams) > 0 {
				newSubst[info.typeParams[0]] = resolvedArg
			}
			methods = append(methods, resolveInterfaceMethods(info.iface, info.imports, pkgInterfaces, visited, newSubst)...)
		case *ast.IndexListExpr:
			// Generic embedded interface with multiple type args: Keyed[string, int]
			ident, ok := ft.X.(*ast.Ident)
			if !ok {
				continue
			}
			if visited[ident.Name] {
				continue
			}
			visited[ident.Name] = true
			info, ok := pkgInterfaces[ident.Name]
			if !ok {
				continue
			}
			newSubst := map[string]ast.Expr{}
			for i, idx := range ft.Indices {
				resolvedArg := resolveExpr(idx, typeSubst)
				if i < len(info.typeParams) {
					newSubst[info.typeParams[i]] = resolvedArg
				}
			}
			methods = append(methods, resolveInterfaceMethods(info.iface, info.imports, pkgInterfaces, visited, newSubst)...)
		}
	}
	return methods
}

func readServiceFile(file *ast.File, packageName string, services ServiceList, pkgInterfaces map[string]interfaceInfo) error {
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
					if starExpr, ok := firstReceiverField.Type.(*ast.StarExpr); ok {
						if ident, ok := starExpr.X.(*ast.Ident); ok {
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
							resolved := resolveInterfaceMethods(iSpec, fileImports, pkgInterfaces, map[string]bool{ident.Name: true}, nil)
							for _, m := range resolved {
								trace(" on sth:", m.name)
								var tpNames []string
								for k := range m.typeSubst {
									tpNames = append(tpNames, k)
								}
								args := readFields(m.funcTyp.Params, m.imports, tpNames...)
								ret := readFields(m.funcTyp.Results, m.imports, tpNames...)
								if len(m.typeSubst) > 0 {
									substituteTypeParams(args, m.typeSubst, m.imports)
									substituteTypeParams(ret, m.typeSubst, m.imports)
								}
								service.Methods = append(service.Methods, &Method{
									Name:   m.name,
									Args:   args,
									Return: ret,
								})
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
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Tok == token.IMPORT {
				trace("got an import", genDecl.Specs)
				for _, spec := range genDecl.Specs {
					if spec, ok := spec.(*ast.ImportSpec); ok {
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
						// trace("  import   >>>>>>>>>>>>>>>>>>>>", importName, importPath)
					}
				}
			}
		}
	}
	return imports
}

func readFields(fieldList *ast.FieldList, fileImports fileImportSpecMap, typeParams ...string) (fields []*Field) {
	trace("reading fields")
	fields = []*Field{}
	if fieldList == nil {
		return
	}

	for _, param := range fieldList.List {
		names, value, _ := readField(param, fileImports, typeParams)
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

// substituteTypeParams replaces TypeParam entries in fields with concrete types from the substitution map.
func substituteTypeParams(fields []*Field, subst map[string]ast.Expr, imports fileImportSpecMap) {
	for _, f := range fields {
		substituteValue(f.Value, subst, imports)
	}
}

// substituteValue recursively replaces TypeParam references with concrete types.
func substituteValue(v *Value, subst map[string]ast.Expr, imports fileImportSpecMap) {
	if v == nil {
		return
	}
	if v.TypeParam != "" {
		if expr, ok := subst[v.TypeParam]; ok {
			wasPtr := v.IsPtr
			*v = Value{}
			v.IsPtr = wasPtr
			v.loadExpr(expr, imports, nil)
		}
		return
	}
	if v.Array != nil {
		substituteValue(v.Array.Value, subst, imports)
	}
	if v.Map != nil {
		substituteValue(v.Map.Key, subst, imports)
		substituteValue(v.Map.Value, subst, imports)
	}
	if v.StructType != nil {
		for _, arg := range v.StructType.TypeArgs {
			substituteValue(arg, subst, imports)
		}
	}
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
	pkgInterfaces := collectPackageInterfaces(pkg, packageName)

	pkgFiles := make([]string, 0, len(pkg.Files))
	for k := range pkg.Files {
		pkgFiles = append(pkgFiles, k)
	}
	sort.Strings(pkgFiles)

	for _, k := range pkgFiles {
		file := pkg.Files[k]
		err = readServiceFile(file, packageName, services, pkgInterfaces)
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
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				switch genDecl.Tok { //nolint:exhaustive
				case token.TYPE:
					trace("got a type", genDecl.Specs)
					for _, spec := range genDecl.Specs {
						if spec, ok := spec.(*ast.TypeSpec); ok {
							if _, ok := constantTypes[spec.Name.Name]; ok {
								continue
							}
							switch specType := spec.Type.(type) {
							case *ast.InterfaceType:
								constantTypes[spec.Name.Name] = "any"
							case *ast.Ident:
								switch specType.Name {
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
						if spec, ok := spec.(*ast.ValueSpec); ok {
							if specType, ok := spec.Type.(*ast.Ident); ok {
								for _, val := range spec.Values {
									if valType, ok := val.(*ast.BasicLit); ok {
										if _, ok := constantTypes[specType.Name]; !ok {
											constantTypes[specType.Name] = map[string]*ast.BasicLit{}
										} else if _, ok := constantTypes[specType.Name].(map[string]*ast.BasicLit); !ok {
											constantTypes[specType.Name] = map[string]*ast.BasicLit{}
										}
										constantTypes[specType.Name].(map[string]*ast.BasicLit)[spec.Names[0].Name] = valType //nolint:forcetypeassert
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
	if s.StructType != nil {
		for _, arg := range s.StructType.TypeArgs {
			loadFlatStructsValue(arg, flatStructs)
		}
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

			// path := fullNameParts[:len(fullNameParts)-1][0]

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
			// packageStructs, structOK := scannedPackageStructs[packageName]
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

func needsWorkValue(value *Value, needsWork func(fullName string) bool) bool {
	switch {
	case value.Scalar != nil:
		if needsWork(value.Scalar.FullName()) {
			return true
		}
	case value.StructType != nil:
		if needsWork(value.StructType.FullName()) {
			return true
		}
		for _, arg := range value.StructType.TypeArgs {
			if needsWorkValue(arg, needsWork) {
				return true
			}
		}
	case value.Array != nil:
		if needsWorkValue(value.Array.Value, needsWork) {
			return true
		}
	case value.Map != nil:
		if needsWorkValue(value.Map.Key, needsWork) || needsWorkValue(value.Map.Value, needsWork) {
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
			if needsWorkValue(field.Value, needsWork) {
				return false
			}

			// var fieldStructType *StructType = nil
			// if field.Value.StructType != nil {
			// 	fieldStructType = field.Value.StructType
			// } else if field.Value.Array != nil && field.Value.Array.Value.StructType != nil {
			// 	fieldStructType = field.Value.Array.Value.StructType
			// } else if field.Value.Map != nil && field.Value.Map.Value.StructType != nil {
			// 	fieldStructType = field.Value.Map.Value.StructType
			// } else if field.Value.Scalar != nil && needsWork(field.Value.Scalar.FullName()) {
			// 	return false
			// } else if field.Value.Array != nil && field.Value.Array.Value.Scalar != nil && needsWork(field.Value.Array.Value.Scalar.FullName()) {
			// 	return false
			// } else if field.Value.Map != nil && field.Value.Map.Value.Scalar != nil && needsWork(field.Value.Map.Value.Scalar.FullName()) {
			// 	return false
			// }
			// if fieldStructType != nil {
			// 	if needsWork(fieldStructType.FullName()) {
			// 		return false
			// 	}
			// }
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
	// // special handling of union only structs
	// if len(s.Fields) == 0 {
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
	// }
	if s.Array != nil {
		if s.Array.Value != nil && needsWorkValue(s.Array.Value, needsWork) {
			return false
		}
	}
	if s.Map != nil {
		if s.Map.Key != nil && needsWorkValue(s.Map.Key, needsWork) {
			return false
		}
		if s.Map.Value != nil && needsWorkValue(s.Map.Value, needsWork) {
			return false
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

func getStructTypesForField(value *Value) []*StructType {
	var types []*StructType
	switch {
	case value.StructType != nil:
		types = append(types, value.StructType)
		for _, arg := range value.StructType.TypeArgs {
			types = append(types, getStructTypesForField(arg)...)
		}
	case value.Map != nil:
		types = append(types, getStructTypesForField(value.Map.Value)...)
	case value.Array != nil:
		types = append(types, getStructTypesForField(value.Array.Value)...)
	}
	return types
}

func getScalarForField(value *Value) []*Scalar {
	// field.Value.StructType
	var scalarTypes []*Scalar
	switch {
	case value.Scalar != nil:
		scalarTypes = append(scalarTypes, value.Scalar)
		// case field.Value.ArrayType
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
				case "error", "net/http.Request", "net/http.ResponseWriter", "context.Context":
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
		for _, strType := range getStructTypesForField(field.Value) {
			if strType != nil {
				fullName := strType.Package + "." + strType.Name
				if len(strType.Package) == 0 {
					fullName = strType.Name
				}
				switch fullName {
				case "error", "net/http.Request", "net/http.ResponseWriter", "context.Context":
					continue
				default:
					structTypes[fullName] = true
				}
			}
		}
	}
}

// func collectStructs(goPath, structs)
