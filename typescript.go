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

func (v *Value) tsType(mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) {
	switch true {
	case v.Map != nil:
		ts.app("{[index:" + v.Map.KeyType + "]:")
		v.Map.Value.tsType(mappings, scalars, structs, ts)
		ts.app("}")
	case v.Array != nil:
		v.Array.Value.tsType(mappings, scalars, structs, ts)
		if v.Array.Value.ScalarType != ScalarTypeByte {
			ts.app("[]")
		}
	case v.Scalar != nil:
		ts.app(tsTypeFromScalarType(v.Scalar.Type))
	case v.StructType != nil:
		if v.StructType.Package != "" {
			mapping, ok := mappings[v.StructType.Package]
			var tsModule string
			if ok {
				tsModule = mapping.TypeScriptModule
			}
			scalarName := v.StructType.FullName()
			// is it a hidden scalar ?!
			hiddenScalar, ok := scalars[scalarName]
			if ok {
				ts.app(tsTypeFromScalarType(hiddenScalar.Type))
				return
			}

			hiddenStruct, okHiddenStruct := structs[scalarName]
			if okHiddenStruct && hiddenStruct.Array != nil {
				if hiddenStruct.Array.Value.StructType != nil {
					hiddenMapping, hiddenMappingOK := mappings[hiddenStruct.Array.Value.StructType.Package]
					var tsModule string
					if hiddenMappingOK {
						tsModule = hiddenMapping.TypeScriptModule
					}
					ts.app(tsModule + "." + hiddenStruct.Array.Value.StructType.Name + "[]")
					return
				}
				if hiddenStruct.Array.Value.StructType == nil {
					if hiddenStruct.Array.Value.GoScalarType == "byte" { // this fixes types like primitive.ID [12]byte
						ts.app(tsTypeFromScalarType(hiddenStruct.Array.Value.ScalarType))
					}
					return
				}
			}

			ts.app(tsModule + "." + v.StructType.Name)
			return
		}
		ts.app(v.StructType.Name)
	case v.Struct != nil:
		// v.Struct.Value.tsType(mappings, ts)
		ts.l("{").ind(1)
		renderStructFields(v.Struct.Fields, mappings, scalars, structs, ts)
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
	case ScalarTypeByte:
		return "string"
	case ScalarTypeBool:
		return "boolean"
	}
	return string(scalarType)
}

func renderStructFields(fields []*Field, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) {
	for _, f := range fields {
		if f.JSONInfo != nil && f.JSONInfo.Ignore {
			continue
		}
		ts.app(f.tsName())
		if f.Value.IsPtr || (f.JSONInfo != nil && f.JSONInfo.OmitEmpty) {
			ts.app("?")
		}
		ts.app(":")
		f.Value.tsType(mappings, scalars, structs, ts)
		ts.app(";")
		ts.nl()
	}
}

func renderTypescriptStruct(str *Struct, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) error {
	if str.Array != nil {
		// skipping array type
		return nil
	}
	ts.l("// " + str.FullName())
	ts.l("export interface " + str.Name + " {").ind(1)
	switch {
	case str.Map != nil:
		ts.app("[index:" + str.Map.KeyType + "]:")
		str.Map.Value.tsType(mappings, scalars, structs, ts)
		ts.app(";")
		ts.nl()
	default:
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
	}
	ts.ind(-1).l("}")
	return nil
}

func renderTypescriptStructsToPackages(
	moduleKind config.ModuleKind,
	structs map[string]*Struct,
	mappings config.TypeScriptMappings,
	constants map[string]map[string]*ast.BasicLit,
	scalarTypes map[string]*Scalar,
	mappedTypeScript map[string]map[string]*code,
) (err error) {
	codeMap := map[string]map[string]*code{}
	for _, mapping := range mappings {
		codeMap[mapping.GoPackage] = map[string]*code{} // newCode().l("module " + mapping.TypeScriptModule + " {").ind(1)
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
		packageCodeMap[str.Name] = newCode("	")
		// fmt.Println("--------------------------->", moduleKind == config.ModuleKindCommonJS)
		if !(moduleKind == config.ModuleKindCommonJS) {
			packageCodeMap[str.Name].ind(1)
		}
		err = renderTypescriptStruct(str, mappings, scalarTypes, structs, packageCodeMap[str.Name])
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
			constCode := newCode("	")
			if moduleKind != config.ModuleKindCommonJS {
				constCode.ind(1)
			}
			constCode.l("// constants from " + packageName).l("export const GoConst = {").ind(1)
			// constCode.l()
			mappedTypeScript[packageName][goConstPseudoPackage] = constCode
			constPrefixParts := split(packageName, []string{"/", ".", "-"})
			constPrefix := ""
			for _, constPrefixPart := range constPrefixParts {
				constPrefix += ucFirst(constPrefixPart)
			}
			constNames := []string{}
			for constName := range packageConstants {
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

func RenderTypeScriptServices(moduleKind config.ModuleKind, tsClientFlavor config.TSClientFlavor, services ServiceList, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, target *config.Target) (typeScript string, err error) {
	ts := newCode("	")
	if !SkipGoTSRPC && tsClientFlavor == "" {

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
				} catch(e) {
	                err(request, e);
				}
				success.apply(null, data);
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
		switch tsClientFlavor {
		case config.TSClientFlavorAsync:
			err = renderTypescriptClientAsync(service, mappings, scalars, structs, ts)
		default:
			err = renderTypescriptClient(SkipGoTSRPC, moduleKind, service, mappings, scalars, structs, ts)
		}
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

func getTSHeaderComment() string {
	return "/* tslint:disable */\n"
}
