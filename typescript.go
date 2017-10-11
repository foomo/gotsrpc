package gotsrpc

import (
	"errors"
	"fmt"
	"go/ast"
	"sort"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

const goConstPseudoPackage = "__goConstants"

// @todo refactor this is wrong
var SkipGoTSRPC = false

func (f *Field) tsName() string {
	n := f.Name
	if f.JSONInfo != nil && len(f.JSONInfo.Name) > 0 {
		n = f.JSONInfo.Name
	}
	return n
}

func (v *Value) tsType(mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, ts *code) {

	switch true {
	case v.Map != nil:
		ts.app("{[index:" + v.Map.KeyType + "]:")
		v.Map.Value.tsType(mappings, scalarTypes, ts)
		ts.app("}")
	case v.Array != nil:
		v.Array.Value.tsType(mappings, scalarTypes, ts)
		ts.app("[]")
	case v.Scalar != nil:
		ts.app(string(v.Scalar.Type))
	case v.StructType != nil:
		if len(v.StructType.Package) > 0 {
			mapping, ok := mappings[v.StructType.Package]
			var tsModule string
			if ok {
				tsModule = mapping.TypeScriptModule
			}
			scalarName := v.StructType.FullName()
			// is it a hdden scalar ?!
			hiddenScalar, ok := scalarTypes[scalarName]
			if ok {
				ts.app(tsTypeFromScalarType(hiddenScalar.Type))
				return
			}
			ts.app(tsModule + "." + v.StructType.Name)
			return
		}
		ts.app(v.StructType.Name)
	case v.Struct != nil:
		//v.Struct.Value.tsType(mappings, ts)
		ts.l("{").ind(1)
		renderStructFields(v.Struct.Fields, mappings, scalarTypes, ts)
		ts.ind(-1).app("}")
	case len(v.ScalarType) > 0:
		ts.app(tsTypeFromScalarType(v.ScalarType))
	default:
		ts.app("any")
	}
	return
}

func tsTypeFromScalarType(scalarType ScalarType) string {
	switch scalarType {
	case ScalarTypeBool:
		return "boolean"
	}
	return string(scalarType)
}

func renderStructFields(fields []*Field, mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, ts *code) {
	for _, f := range fields {
		if f.JSONInfo != nil && f.JSONInfo.Ignore {
			continue
		}
		ts.app(f.tsName())
		if f.Value.IsPtr || (f.JSONInfo != nil && f.JSONInfo.OmitEmpty) {
			ts.app("?")
		}
		ts.app(":")
		f.Value.tsType(mappings, scalarTypes, ts)
		ts.app(";")
		ts.nl()
	}

}

func renderStruct(str *Struct, mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, ts *code) error {
	ts.l("// " + str.FullName())
	ts.l("export interface " + str.Name + " {").ind(1)
	renderStructFields(str.Fields, mappings, scalarTypes, ts)
	ts.ind(-1).l("}")
	return nil
}

func renderService(skipGoTSRPC bool, moduleKind config.ModuleKind, service *Service, mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, ts *code) error {
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
			arg.Value.tsType(mappings, scalarTypes, ts)
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
			retField.Value.tsType(mappings, scalarTypes, ts)
		}

		ts.app(") => void")
		ts.app(", err:(request:XMLHttpRequest) => void) {").nl()
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
func RenderStructsToPackages(structs map[string]*Struct, mappings config.TypeScriptMappings, constants map[string]map[string]*ast.BasicLit, scalarTypes map[string]*Scalar, mappedTypeScript map[string]map[string]*code) (err error) {

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
		packageCodeMap[str.Name] = newCode("	").ind(1)
		err = renderStruct(str, mappings, scalarTypes, packageCodeMap[str.Name])
		if err != nil {
			return
		}
	}
	ensureCodeInPackage := func(goPackage string) {
		_, ok := mappedTypeScript[goPackage]
		if !ok {
			mappedTypeScript[goPackage] = map[string]*code{}
		}
		return
	}
	for _, mapping := range mappings {
		for structName, structCode := range codeMap[mapping.GoPackage] {
			ensureCodeInPackage(mapping.GoPackage)
			mappedTypeScript[mapping.GoPackage][structName] = structCode
		}
	}
	for packageName, packageConstants := range constants {
		if len(packageConstants) > 0 {
			ensureCodeInPackage(packageName)
			_, done := mappedTypeScript[packageName][goConstPseudoPackage]
			if done {
				continue
			}
			constCode := newCode("	").ind(1).l("// constants from " + packageName).l("export const GoConst = {").ind(1)
			//constCode.l()
			mappedTypeScript[packageName][goConstPseudoPackage] = constCode
			constPrefixParts := split(packageName, []string{"/", ".", "-"})
			constPrefix := ""
			for _, constPrefixPart := range constPrefixParts {
				constPrefix += ucFirst(constPrefixPart)
			}
			constNames := []string{}
			for constName, _ := range packageConstants {
				constNames = append(constNames, constName)
			}
			sort.Strings(constNames)
			for _, constName := range constNames {
				basicLit := packageConstants[constName]
				constCode.l(fmt.Sprint(constName, " : ", basicLit.Value, ","))
			}
			constCode.ind(-1).l("}")

		}

	}
	return nil
}

func split(str string, seps []string) []string {
	res := []string{}
	strs := []string{str}
	for _, sep := range seps {
		nextStrs := []string{}
		for _, str := range strs {
			for _, part := range strings.Split(str, sep) {
				nextStrs = append(nextStrs, part)
			}
		}
		strs = nextStrs
		res = nextStrs
	}
	return res
}

func ucFirst(str string) string {
	strUpper := strings.ToUpper(str)
	constPrefix := ""
	var firstRune rune
	for _, strUpperRune := range strUpper {
		firstRune = strUpperRune
		break
	}
	constPrefix += string(firstRune)
	for i, strRune := range str {
		if i == 0 {
			continue
		}
		constPrefix += string(strRune)
	}
	return constPrefix
}

func RenderTypeScriptServices(moduleKind config.ModuleKind, services ServiceList, mappings config.TypeScriptMappings, scalarTypes map[string]*Scalar, target *config.Target) (typeScript string, err error) {
	ts := newCode("	")
	if !SkipGoTSRPC {

		if moduleKind != config.ModuleKindCommonJS {
			ts.l(`module GoTSRPC {`)
		}

		ts.l(`export const call = (endPoint:string, method:string, args:any[], success:any, err:any) => {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
		// this causes problems, when the browser decides to do a cors OPTIONS request
        // request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        request.send(JSON.stringify(args));
        request.onload = function() {
            if (request.status == 200) {
				try {
					var data = JSON.parse(request.responseText);
					success.apply(null, data);
				} catch(e) {
	                err(e);
				}
            } else {
                err(request);
            }
        };
        request.onerror = function() {
            err(request);
        };
    }
`)

	}
	if config.ModuleKindCommonJS != moduleKind {
		if !SkipGoTSRPC {
			ts.l("} // close")
		}
		ts.l("module " + target.TypeScriptModule + " {")
		ts.ind(1)
	}

	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !target.IsTSRPC(service.Name) {
			continue
		}
		err = renderService(SkipGoTSRPC, moduleKind, service, mappings, scalarTypes, ts)
		if err != nil {
			return
		}
	}
	if config.ModuleKindCommonJS != moduleKind {
		ts.ind(-1)
		ts.l("}")
	}
	typeScript = ts.string()
	return
}
