package gotsrpc

import (
	"fmt"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

func renderTypescriptClient(skipGoTSRPC bool, moduleKind config.ModuleKind, service *Service, mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, structs map[string]*Struct, ts *code) error {
	clientName := service.Name + "Client"

	ts.l("export class " + clientName + " {").ind(1)

	if moduleKind == config.ModuleKindCommonJS {
		if skipGoTSRPC {
			ts.l("constructor(public endPoint:string = \"" + service.Endpoint + "\", public transport:(endPoint:string, method:string, args:any[], success:any, err:any) => void) {  }")
		} else {
			ts.l("static defaultInst = new " + clientName + ";")
			ts.l("constructor(public endPoint:string = \"" + service.Endpoint + "\", public transport = call) {  }")
		}

	} else {
		ts.l("static defaultInst = new " + clientName + ";")
		ts.l("constructor(public endPoint:string = \"" + service.Endpoint + "\", public transport = GoTSRPC.call) {  }")
	}

	for _, method := range service.Methods {

		ts.app(lcfirst(method.Name) + "(")
		// actual args
		//args := []string{}
		callArgs := []string{}

		argOffset := 0
		for index, arg := range method.Args {
			if index == 0 && arg.Value.isHTTPResponseWriter() {
				trace("skipping first arg is a http.ResponseWriter")
				argOffset = 1
				continue
			}
			if index == 1 && arg.Value.isHTTPRequest() {
				trace("skipping second arg is a *http.Request")
				argOffset = 2
				continue
			}
		}
		argCount := 0
		for index, arg := range method.Args {
			if index < argOffset {
				continue
			}
			if index > argOffset {
				ts.app(", ")
			}
			ts.app(arg.tsName() + ":")
			arg.Value.tsType(mappings, scalarTypes, structs, ts)
			callArgs = append(callArgs, arg.Name)
			argCount++
		}
		if argCount > 0 {
			ts.app(", ")
		}
		ts.app("success:(")
		// + strings.Join(retArgs, ", ") +

		for index, retField := range method.Return {
			retArgName := retField.tsName()
			if len(retArgName) == 0 {
				retArgName = "ret"
				if index > 0 {
					retArgName += "_" + fmt.Sprint(index)
				}
			}
			if index > 0 {
				ts.app(", ")
			}
			ts.app(retArgName + ":")
			retField.Value.tsType(mappings, scalarTypes, structs, ts)
		}

		ts.app(") => void")
		ts.app(", err:(request:XMLHttpRequest, e?:Error) => void) {").nl()
		ts.ind(1)
		// generic framework call
		ts.l("this.transport(this.endPoint, \"" + method.Name + "\", [" + strings.Join(callArgs, ", ") + "], success, err);")
		ts.ind(-1)
		ts.app("}")
		ts.nl()
	}
	ts.ind(-1)
	ts.l("}")
	return nil
}
