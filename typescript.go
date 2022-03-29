package gotsrpc

import (
	"errors"
	"go/ast"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/foomo/gotsrpc/v2/config"
)

func (f *Field) tsName() string {
	n := f.Name
	if f.JSONInfo != nil && len(f.JSONInfo.Name) > 0 {
		n = f.JSONInfo.Name
	}
	return n
}

var tsTypeAliases = map[string]string{
	"time.Time": "number",
}

func (v *Value) tsType(mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code, jsonInfo *JSONInfo) {
	switch {
	case jsonInfo != nil && len(jsonInfo.Type) > 0:
		ts.app(jsonInfo.Type)
	case v.Map != nil:
		ts.app("Record<")
		if v.Map.Key != nil {
			v.Map.Key.tsType(mappings, scalars, structs, ts, nil)
		} else {
			ts.app(v.Map.KeyType)
		}
		ts.app(",")
		v.Map.Value.tsType(mappings, scalars, structs, ts, nil)
		ts.app(">")
		if jsonInfo == nil || !jsonInfo.OmitEmpty {
			ts.app("|null")
		}
	case v.Array != nil:
		if v.Array.Value.ScalarType != ScalarTypeByte {
			ts.app("Array<")
		}
		v.Array.Value.tsType(mappings, scalars, structs, ts, nil)
		if v.Array.Value.ScalarType != ScalarTypeByte {
			ts.app(">")
		}
		if jsonInfo == nil || !jsonInfo.OmitEmpty {
			ts.app("|null")
		}
	case v.Scalar != nil:
		if v.Scalar.Package != "" {
			mapping, ok := mappings[v.Scalar.Package]
			var tsModule string
			if ok {
				tsModule = mapping.TypeScriptModule
			}
			tsType := tsModule + "." + tsTypeFromScalarType(v.ScalarType)
			if value, ok := tsTypeAliases[tsType]; ok {
				tsType = value
			}
			ts.app(tsType)
			if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.app("|null")
			}
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
			ts.app(tsModule + "." + v.StructType.Name)
			hiddenStruct, isHiddenStruct := structs[v.StructType.FullName()]
			if isHiddenStruct && (hiddenStruct.Array != nil || hiddenStruct.Map != nil) && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.app("|null")
			} else if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.app("|null")
			}
			return
		}
		ts.app(v.StructType.Name)
	case v.Struct != nil:
		// v.Struct.Value.tsType(mappings, ts)
		ts.l("{").ind(1)
		renderStructFields(v.Struct.Fields, mappings, scalars, structs, ts)
		ts.ind(-1).app("}")
		if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
			ts.app("|null")
		}
	case len(v.ScalarType) > 0:
		ts.app(tsTypeFromScalarType(v.ScalarType))
		if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
			ts.app("|null")
		}
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
		if len(f.Name) == 0 {
			continue
		} else if f.JSONInfo != nil && f.JSONInfo.Ignore {
			continue
		}
		ts.app(f.tsName())
		if f.JSONInfo != nil && f.JSONInfo.OmitEmpty {
			ts.app("?")
		}
		ts.app(":")
		f.Value.tsType(mappings, scalars, structs, ts, f.JSONInfo)
		ts.app(";")
		ts.nl()
	}
}

func renderTypescriptStruct(str *Struct, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) error {
	ts.l("// " + str.FullName())
	switch {
	case str.Array != nil:
		ts.app("export type " + str.Name + " = Array<")
		str.Array.Value.tsType(mappings, scalars, structs, ts, nil)
		ts.app(">")
		ts.nl()
	case str.Map != nil:
		ts.app("export type " + str.Name + " = Record<")
		if str.Map.Key != nil {
			str.Map.Key.tsType(mappings, scalars, structs, ts, nil)
		} else {
			ts.app(str.Map.KeyType)
		}
		ts.app(",")
		str.Map.Value.tsType(mappings, scalars, structs, ts, nil)
		ts.app(">")
		ts.nl()
	// special handling of inline only structs
	case len(str.UnionFields) > 0:
		if len(str.Fields) > 0 || len(str.InlineFields) > 0 {
			return errors.New("no fields or inline fields are supported when using union")
		}
		switch {
		case str.UnionFields[0].Value.StructType != nil:
			ts.app("export type " + str.Name + " = ")
			var isUndefined bool
			for i, unionField := range str.UnionFields {
				if i > 0 {
					ts.app(" | ")
				}
				unionField.Value.tsType(mappings, scalars, structs, ts, &JSONInfo{OmitEmpty: true})
				if unionField.Value.IsPtr {
					isUndefined = true
				}
			}
			if isUndefined {
				ts.app(" | undefined")
			}
			ts.nl()
		case str.UnionFields[0].Value.Scalar != nil:
			ts.app("export type " + str.Name + " = ")
			for i, field := range str.UnionFields {
				if i > 0 {
					ts.app(" & ")
				}
				ts.app("(typeof ")
				field.Value.tsType(mappings, scalars, structs, ts, &JSONInfo{OmitEmpty: true})
				ts.app(")")
			}
			ts.nl()
		default:
			return errors.New("could not resolve this union type")
		}
	case len(str.InlineFields) > 0:
		ts.app("export interface " + str.Name + " extends ")
		for i, inlineField := range str.InlineFields {
			if i > 0 {
				ts.app(", ")
			}
			if inlineField.Value.IsPtr {
				ts.app("Partial<")
			}
			inlineField.Value.tsType(mappings, scalars, structs, ts, &JSONInfo{OmitEmpty: true})
			if inlineField.Value.IsPtr {
				ts.app(">")
			}
			ts.app(" ")
		}
		ts.app("{")
		ts.nl()
		ts.ind(1)
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
		ts.ind(-1).l("}")
	default:
		ts.l("export interface " + str.Name + " {").ind(1)
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
		ts.ind(-1).l("}")
	}
	return nil
}

func renderTypescriptStructsToPackages(
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
	var res []string
	strs := []string{str}
	for _, sep := range seps {
		var nextStrs []string
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

func RenderTypeScriptServices(services ServiceList, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, target *config.Target) (typeScript string, err error) {
	ts := newCode("	")
	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !target.IsTSRPC(service.Name) {
			continue
		}
		err = renderTypescriptClient(service, mappings, scalars, structs, ts)
		if err != nil {
			return
		}
	}
	typeScript = ts.string()
	return
}

func getTSHeaderComment() string {
	return "/* eslint:disable */\n"
}
