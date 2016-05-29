package gotsrpc

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func readServiceFile(file *ast.File, services []*Service) error {
	findService := func(serviceName string) (service *Service, ok bool) {
		for _, service := range services {
			if service.Name == serviceName {
				return service, true
			}
		}
		return nil, false
	}
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
							fmt.Println("	on sth:", ident.Name)

							service, ok := findService(ident.Name)

							if ok {
								service.Methods = append(service.Methods, &Method{
									Name:   funcDecl.Name.Name,
									Args:   readFields(funcDecl.Type.Params),
									Return: readFields(funcDecl.Type.Results),
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

func readFields(fieldList *ast.FieldList) (fields []*Field) {
	fields = []*Field{}
	if fieldList == nil {
		return
	}

	for _, param := range fieldList.List {
		name, value, _ := readField(param)
		fields = append(fields, &Field{
			Name:  name,
			Value: value,
		})
	}
	return

}

func readServicesInPackage(pkg *ast.Package, serviceNames []string) (services []*Service, err error) {
	services = []*Service{}
	for _, serviceName := range serviceNames {
		services = append(services, &Service{
			Name:    serviceName,
			Methods: []*Method{},
		})
	}
	for _, file := range pkg.Files {
		err = readServiceFile(file, services)
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
	services, err = readServicesInPackage(pkg, serviceNames)
	if err != nil {
		return
	}
	structs, err = readStructs(pkg)
	return
}
