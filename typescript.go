package gotsrpc

import (
	"errors"
	"fmt"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

var SkipGoTSRPC = false

func (f *Field) tsName() string {
	n := f.Name
	if f.JSONInfo != nil && len(f.JSONInfo.Name) > 0 {
		n = f.JSONInfo.Name
	}
	return n
}

func (v *Value) tsType(mappings config.TypeScriptMappings) string {
	switch true {
	case v.IsPtr:
		if v.StructType != nil {
			if len(v.StructType.Package) > 0 {
				mapping, ok := mappings[v.StructType.Package]
				var tsModule string
				if ok {
					tsModule = mapping.TypeScriptModule
				}
				return tsModule + "." + v.StructType.Name
			}
			return v.StructType.Name
		}
		return string(v.ScalarType)
	case v.Array != nil:
		return v.Array.Value.tsType(mappings) + "[]"
	case len(v.ScalarType) > 0:
		switch v.ScalarType {
		case ScalarTypeBool:
			return "boolean"
		default:
			return string(v.ScalarType)
		}

	default:
		return "any"
	}
}

func renderStruct(str *Struct, mappings config.TypeScriptMappings, ts *code) error {
	ts.l("// " + str.FullName())
	ts.l("export interface " + str.Name + " {").ind(1)
	for _, f := range str.Fields {
		if f.JSONInfo != nil && f.JSONInfo.Ignore {
			continue
		}
		ts.app(f.tsName())
		if f.Value.IsPtr {
			ts.app("?")
		}
		ts.app(":" + f.Value.tsType(mappings))
		ts.app(";")
		ts.nl()
	}
	ts.ind(-1).l("}")
	return nil
}

/*
   export class ServiceClient {
       static defaultInst = new ServiceClient()
       constructor(public endPoint:string = "/service") {  }
       hello(name:string, success:(reply:string, err:Err) => void, err:(request:XMLHttpRequest) => void) {
           GoTSRPC.call(this.endPoint, "Hello", [name], success, err);
       }
   }
*/

func renderService(service *Service, mappings config.TypeScriptMappings, ts *code) error {
	clientName := service.Name + "Client"
	ts.l("export class " + clientName + " {").ind(1).
		l("static defaultInst = new " + clientName + ";").
		l("constructor(public endPoint:string = \"/service\") {  }")
	for _, method := range service.Methods {

		ts.app(lcfirst(method.Name) + "(")
		// actual args
		args := []string{}
		callArgs := []string{}

		for index, arg := range method.Args {
			if index == 0 && arg.Value.isHTTPResponseWriter() {
				trace("skipping first arg is a http.ResponseWriter")
				continue
			}
			if index == 1 && arg.Value.isHTTPRequest() {
				trace("skipping second arg is a *http.Request")
				continue
			}

			args = append(args, arg.tsName()+":"+arg.Value.tsType(mappings))
			callArgs = append(callArgs, arg.Name)
		}
		ts.app(strings.Join(args, ", "))
		// success callback
		retArgs := []string{}
		for index, retField := range method.Return {
			retArgName := retField.tsName()
			if len(retArgName) == 0 {
				retArgName = "ret"
				if index > 0 {
					retArgName += "_" + fmt.Sprint(index)
				}
			}
			retArgs = append(retArgs, retArgName+":"+retField.Value.tsType(mappings))
		}
		if len(args) > 0 {
			ts.app(", ")
		}
		ts.app("success:(" + strings.Join(retArgs, ", ") + ") => void")
		ts.app(", err:(request:XMLHttpRequest) => void) {").nl()
		ts.ind(1)
		// generic framework call
		ts.l("GoTSRPC.call(this.endPoint, \"" + method.Name + "\", [" + strings.Join(callArgs, ", ") + "], success, err);")
		ts.ind(-1)
		ts.app("}")
		ts.nl()
	}
	ts.ind(-1)
	ts.l("}")
	return nil
}
func RenderStructsToPackages(structs map[string]*Struct, mappings config.TypeScriptMappings, mappedTypeScript map[string]map[string]*code) (err error) {

	codeMap := map[string]map[string]*code{}
	for _, mapping := range mappings {
		codeMap[mapping.GoPackage] = map[string]*code{} //newCode().l("module " + mapping.TypeScriptModule + " {").ind(1)
	}
	for name, str := range structs {
		if str == nil {
			err = errors.New("could not resolve: " + name)
			return
		}

		packageCodeMap, ok := codeMap[str.Package]
		if !ok {
			err = errors.New("missing code mapping for go package : " + str.Package + " => you have to add a mapping from this go package to a TypeScript module in your build-config.yml in the mappings section")
			return
		}
		packageCodeMap[str.Name] = newCode().ind(1)
		err = renderStruct(str, mappings, packageCodeMap[str.Name])
		if err != nil {
			return
		}
	}
	for _, mapping := range mappings {
		for structName, structCode := range codeMap[mapping.GoPackage] {
			_, ok := mappedTypeScript[mapping.GoPackage]
			if !ok {
				mappedTypeScript[mapping.GoPackage] = map[string]*code{}
			}
			mappedTypeScript[mapping.GoPackage][structName] = structCode
		}
		//.ind(-1).l("}").string()
	}
	return nil
}
func RenderTypeScriptServices(services []*Service, mappings config.TypeScriptMappings, tsModuleName string) (typeScript string, err error) {
	ts := newCode()
	if !SkipGoTSRPC {
		ts.l(`module GoTSRPC {
    export function call(endPoint:string, method:string, args:any[], success:any, err:any) {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));            
        request.onload = function() {
            if (request.status == 200) {
				try {
					var data = JSON.parse(request.responseText);
					success.apply(null, data);
				} catch(e) {
	                err(request);
				}
            } else {
                err(request);
            }
        };            
        request.onerror = function() {
            err(request);
        };
    }
}`)
	}

	ts.l("module " + tsModuleName + " {")
	ts.ind(1)

	for _, service := range services {
		err = renderService(service, mappings, ts)
		if err != nil {
			return
		}
	}
	ts.ind(-1)
	ts.l("}")
	typeScript = ts.string()
	return
}
