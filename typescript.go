package gotsrpc

import (
	"errors"
	"go/ast"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/foomo/gotsrpc/v2/config"
)

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
		ts.app("Record<")
		if v.Map.Key != nil {
			v.Map.Key.tsType(mappings, scalars, structs, ts)
		} else {
			ts.app(v.Map.KeyType)
		}
		ts.app(",")
		v.Map.Value.tsType(mappings, scalars, structs, ts)
		ts.app(">")
	case v.Array != nil:
		v.Array.Value.tsType(mappings, scalars, structs, ts)
		if v.Array.Value.ScalarType != ScalarTypeByte {
			ts.app("[]")
		}
	case v.Scalar != nil:
		if v.Scalar.Package != "" {
			mapping, ok := mappings[v.Scalar.Package]
			var tsModule string
			if ok {
				tsModule = mapping.TypeScriptModule
			}
			ts.app(tsModule + "." + tsTypeFromScalarType(v.ScalarType))
			return
		}
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
				} else if hiddenStruct.Array.Value.Scalar != nil {
					var tsModule string
					if value, ok := mappings[hiddenStruct.Array.Value.Scalar.Package]; ok {
						tsModule = value.TypeScriptModule
					}
					ts.app(tsModule + "." + tsTypeFromScalarType(hiddenStruct.Array.Value.Scalar.Type))
					return
				} else if hiddenStruct.Array.Value.GoScalarType == "byte" { // this fixes types like primitive.ID [12]byte
					ts.app(tsTypeFromScalarType(hiddenStruct.Array.Value.ScalarType))
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
	switch {
	case str.Map != nil:
		ts.app("export type " + str.Name + " = Record<")
		if str.Map.Key != nil {
			str.Map.Key.tsType(mappings, scalars, structs, ts)
		} else {
			ts.app(str.Map.KeyType)
		}
		ts.app(",")
		str.Map.Value.tsType(mappings, scalars, structs, ts)
		ts.app(">")
		ts.nl()
	default:
		ts.l("export interface " + str.Name + " {").ind(1)
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
		ts.ind(-1).l("}")
	}
	return nil
}

func renderTypescriptStructsToPackages(
	moduleKind config.ModuleKind,
	structs map[string]*Struct,
	mappings config.TypeScriptMappings,
	constantTypes map[string]map[string]interface{},
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

	for packageName, packageConstantTypes := range constantTypes {
		if len(packageConstantTypes) > 0 {
			packageCodeMap, ok := codeMap[packageName]
			if !ok {
				err = errors.New("missing code mapping for go package : " + packageName + " => you have to add a mapping from this go package to a TypeScript module in your build-config.yml in the mappings section")
				return
			}
			for packageConstantTypeName, packageConstantTypeValues := range packageConstantTypes {
				packageCodeMap[packageConstantTypeName] = newCode("	")
				if !(moduleKind == config.ModuleKindCommonJS) {
					packageCodeMap[packageConstantTypeName].ind(1)
				}
				packageCodeMap[packageConstantTypeName].l("// " + packageName + "." + packageConstantTypeName)

				if packageConstantTypeValuesList, ok := packageConstantTypeValues.(map[string]*ast.BasicLit); ok {
					keys := make([]string, 0, len(packageConstantTypeValuesList))
					for k := range packageConstantTypeValuesList {
						keys = append(keys, k)
					}
					sort.Strings(keys)
					packageCodeMap[packageConstantTypeName].l("export enum " + packageConstantTypeName + " {").ind(1)
					for _, k := range keys {
						enum := strings.TrimPrefix(strcase.ToCamel(k), packageConstantTypeName)
						packageCodeMap[packageConstantTypeName].l(enum + " = " + packageConstantTypeValuesList[k].Value + ",")
					}
					packageCodeMap[packageConstantTypeName].ind(-1).l("}")

				} else if packageConstantTypeValuesString, ok := packageConstantTypeValues.(string); ok {
					packageCodeMap[packageConstantTypeName].l("export type " + packageConstantTypeName + " = " + packageConstantTypeValuesString)
				}
			}
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
	return "/* eslint:disable */\n"
}
