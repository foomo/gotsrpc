package gotsrpc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/foomo/gotsrpc/v2/config"
)

func renderTypescriptClient(service *Service, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) error {
	clientName := service.Name + "Client"

	ts.l("export class " + clientName + " {")

	ts.ind(1)
	// ts.l(`static defaultInst = new ` + clientName + `()`)
	// ts.l(`constructor(public endpoint = "` + service.Endpoint + `") {}`)
	ts.l(`public static defaultEndpoint = "` + service.Endpoint + `";`)
	ts.l("constructor(")
	ts.ind(1)
	ts.l("public transport:<T>(method: string, data?: any[]) => Promise<T>")
	ts.ind(-1)
	ts.l(") {}")

	for _, method := range service.Methods {
		ts.app("async " + lcfirst(method.Name) + "(")
		var callArgs []string
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
			ts.app(arg.tsName())
			ts.app(":")
			arg.Value.tsType(mappings, scalars, structs, ts, arg.JSONInfo)
			callArgs = append(callArgs, arg.Name)
			argCount++
		}
		ts.app("):")

		returnTypeTS := newCode("	")
		returnTypeTS.app("{")
		innerReturnTypeTS := newCode("	")
		innerReturnTypeTS.app("{")
		firstReturnType := ""
		countReturns := 0
		countInnerReturns := 0
		responseObjectPrefix := ""
		responseObject := "return {"

		for index, retField := range method.Return {
			countInnerReturns++
			retArgName := retField.tsName()

			if len(retArgName) == 0 {
				retArgName = "ret"
				if index > 0 {
					retArgName += "_" + fmt.Sprint(index)
				}
			}
			if index > 0 {
				returnTypeTS.app("; ")
				innerReturnTypeTS.app("; ")
			}

			innerReturnTypeTS.app(strconv.Itoa(index))
			innerReturnTypeTS.app(":")
			retField.Value.tsType(mappings, scalars, structs, innerReturnTypeTS, retField.JSONInfo)

			if index == 0 {
				firstReturnTypeTS := newCode("	")
				retField.Value.tsType(mappings, scalars, structs, firstReturnTypeTS, retField.JSONInfo)
				firstReturnType = firstReturnTypeTS.string()
			}
			countReturns++
			returnTypeTS.app(retArgName)
			returnTypeTS.app(":")
			responseObject += responseObjectPrefix + retArgName + " : response[" + strconv.Itoa(index) + "]"
			retField.Value.tsType(mappings, scalars, structs, returnTypeTS, retField.JSONInfo)
			responseObjectPrefix = ", "
		}
		responseObject += "};"
		returnTypeTS.app("}")
		innerReturnTypeTS.app("}")
		if countReturns == 0 {
			ts.app("Promise<void> {")
		} else if countReturns == 1 {
			ts.app("Promise<" + firstReturnType + "> {")
		} else if countReturns > 1 {
			ts.app("Promise<" + returnTypeTS.string() + "> {")
		}
		ts.nl()

		ts.ind(1)

		innerCallTypeString := "void"
		if countInnerReturns > 0 {
			innerCallTypeString = innerReturnTypeTS.string()
		}

		call := "this.transport<" + innerCallTypeString + ">(\"" + method.Name + "\", [" + strings.Join(callArgs, ", ") + "])"
		switch countReturns {
		case 0:
			ts.l("await " + call)
		case 1:
			ts.l("return (await " + call + ")[0]")
		default:
			ts.l("const response = await " + call)
			ts.l(responseObject)
		}

		ts.ind(-1)
		ts.app("}")
		ts.nl()
	}
	ts.ind(-1)
	ts.l("}")
	return nil
}
