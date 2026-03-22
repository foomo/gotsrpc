package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/foomo/gotsrpc/v2/config"
	"github.com/foomo/gotsrpc/v2/internal/model"
)

func renderTypescriptClient(service *model.Service, mappings config.TypeScriptMappings, scalars map[string]*model.Scalar, structs map[string]*model.Struct, ts *Code) error {
	clientName := service.Name + "Client"

	ts.L("export class " + clientName + " {")

	ts.Ind(1)
	ts.L(`public static defaultEndpoint = "` + service.Endpoint + `";`)
	ts.L("constructor(")
	ts.Ind(1)
	ts.L("public transport:<T>(method: string, data?: any[]) => Promise<T>")
	ts.Ind(-1)
	ts.L(") {}")

	for _, method := range service.Methods {
		ts.App("async " + lcfirst(method.Name) + "(")

		var callArgs []string

		argOffset := 0

		for index, arg := range method.Args {
			if index == 0 && valueIsHTTPResponseWriter(arg.Value) {
				trace("skipping first arg is a http.ResponseWriter")

				argOffset = 1

				continue
			} else if index == 0 && valueIsContext(arg.Value) {
				trace("skipping first arg is a context.Context")

				argOffset = 1

				continue
			}

			if index == 1 && valueIsHTTPRequest(arg.Value) {
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
				ts.App(", ")
			}

			ts.App(fieldTSName(arg))
			ts.App(":")
			valueTSType(arg.Value, mappings, scalars, structs, ts, arg.JSONInfo)
			callArgs = append(callArgs, arg.Name)
			argCount++
		}

		ts.App("):")

		returnTypeTS := NewCode("	")
		returnTypeTS.App("{")

		innerReturnTypeTS := NewCode("	")
		innerReturnTypeTS.App("{")

		firstReturnType := ""
		countReturns := 0
		countInnerReturns := 0
		responseObjectPrefix := ""

		var responseObject strings.Builder
		responseObject.WriteString("return {")

		for index, retField := range method.Return {
			countInnerReturns++
			retArgName := fieldTSName(retField)

			if len(retArgName) == 0 {
				retArgName = "ret"
				if index > 0 {
					retArgName += "_" + fmt.Sprint(index)
				}
			}

			if index > 0 {
				returnTypeTS.App("; ")
				innerReturnTypeTS.App("; ")
			}

			innerReturnTypeTS.App(strconv.Itoa(index))
			innerReturnTypeTS.App(":")
			valueTSType(retField.Value, mappings, scalars, structs, innerReturnTypeTS, retField.JSONInfo)

			if index == 0 {
				firstReturnTypeTS := NewCode("	")
				valueTSType(retField.Value, mappings, scalars, structs, firstReturnTypeTS, retField.JSONInfo)
				firstReturnType = firstReturnTypeTS.String()
			}

			countReturns++

			returnTypeTS.App(retArgName)
			returnTypeTS.App(":")

			responseObject.WriteString(responseObjectPrefix + retArgName + " : response[" + strconv.Itoa(index) + "]")

			valueTSType(retField.Value, mappings, scalars, structs, returnTypeTS, retField.JSONInfo)

			responseObjectPrefix = ", "
		}

		responseObject.WriteString("};")

		returnTypeTS.App("}")
		innerReturnTypeTS.App("}")

		if countReturns == 0 {
			ts.App("Promise<void> {")
		} else if countReturns == 1 {
			ts.App("Promise<" + firstReturnType + "> {")
		} else if countReturns > 1 {
			ts.App("Promise<" + returnTypeTS.String() + "> {")
		}

		ts.NL()

		ts.Ind(1)

		innerCallTypeString := "void"
		if countInnerReturns > 0 {
			innerCallTypeString = innerReturnTypeTS.String()
		}

		call := "this.transport<" + innerCallTypeString + ">(\"" + method.Name + "\", [" + strings.Join(callArgs, ", ") + "])"

		switch countReturns {
		case 0:
			ts.L("await " + call)
		case 1:
			ts.L("return (await " + call + ")[0]")
		default:
			ts.L("const response = await " + call)
			ts.L(responseObject.String())
		}

		ts.Ind(-1)
		ts.App("}")
		ts.NL()
	}

	ts.Ind(-1)
	ts.L("}")

	return nil
}
