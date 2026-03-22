package codegen

import (
	"errors"
	"go/ast"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/foomo/gotsrpc/v2/config"
	"github.com/foomo/gotsrpc/v2/internal/model"
)

func fieldTSName(f *model.Field) string {
	n := f.Name
	if f.JSONInfo != nil && len(f.JSONInfo.Name) > 0 {
		n = f.JSONInfo.Name
	}
	return n
}

var tsTypeAliases = map[string]string{
	"time.Time": "number",
}

func valueTSType(v *model.Value, mappings config.TypeScriptMappings, scalars map[string]*model.Scalar, structs map[string]*model.Struct, ts *Code, jsonInfo *model.JSONInfo) {
	switch {
	case jsonInfo != nil && len(jsonInfo.Type) > 0:
		ts.App(jsonInfo.Type)
	case v.Map != nil:
		ts.App("Record<")
		if v.Map.Key != nil {
			valueTSType(v.Map.Key, mappings, scalars, structs, ts, nil)
		} else {
			ts.App(v.Map.KeyType)
		}
		ts.App(",")
		valueTSType(v.Map.Value, mappings, scalars, structs, ts, nil)
		ts.App(">")
		if jsonInfo == nil || !jsonInfo.OmitEmpty {
			ts.App("|null")
		}
	case v.Array != nil:
		if v.Array.Value.ScalarType != model.ScalarTypeByte {
			ts.App("Array<")
		}
		valueTSType(v.Array.Value, mappings, scalars, structs, ts, nil)
		if v.Array.Value.ScalarType != model.ScalarTypeByte {
			ts.App(">")
		}
		if jsonInfo == nil || !jsonInfo.OmitEmpty {
			ts.App("|null")
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
			ts.App(tsType)
			if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.App("|null")
			}
			return
		}
		ts.App(tsTypeFromScalarType(v.Scalar.Type))
	case v.StructType != nil:
		if v.StructType.Package != "" {
			mapping, ok := mappings[v.StructType.Package]
			var tsModule string
			if ok {
				tsModule = mapping.TypeScriptModule
			}
			ts.App(tsModule + "." + v.StructType.Name)
			hiddenStruct, isHiddenStruct := structs[v.StructType.FullName()]
			if isHiddenStruct && (hiddenStruct.Array != nil || hiddenStruct.Map != nil) && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.App("|null")
			} else if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
				ts.App("|null")
			}
			return
		}
		ts.App(v.StructType.Name)
	case v.Struct != nil:
		ts.L("{").Ind(1)
		renderStructFields(v.Struct.Fields, mappings, scalars, structs, ts)
		ts.Ind(-1).App("}")
		if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
			ts.App("|null")
		}
	case len(v.ScalarType) > 0:
		ts.App(tsTypeFromScalarType(v.ScalarType))
		if v.IsPtr && (jsonInfo == nil || !jsonInfo.OmitEmpty) {
			ts.App("|null")
		}
	default:
		ts.App("any")
	}
}

func tsTypeFromScalarType(scalarType model.ScalarType) string {
	switch scalarType { //nolint:exhaustive
	case model.ScalarTypeError:
		return "github_com_foomo_gotsrpc_v2.Error"
	case model.ScalarTypeByte:
		return "string"
	case model.ScalarTypeBool:
		return "boolean"
	}
	return string(scalarType)
}

func renderStructFields(fields []*model.Field, mappings config.TypeScriptMappings, scalars map[string]*model.Scalar, structs map[string]*model.Struct, ts *Code) {
	for _, f := range fields {
		if len(f.Name) == 0 {
			continue
		} else if f.JSONInfo != nil && f.JSONInfo.Ignore {
			continue
		}
		ts.App(fieldTSName(f))
		if f.JSONInfo != nil && f.JSONInfo.OmitEmpty {
			ts.App("?")
		}
		ts.App(":")
		valueTSType(f.Value, mappings, scalars, structs, ts, f.JSONInfo)
		ts.App(";")
		ts.NL()
	}
}

func renderTypescriptStruct(str *model.Struct, mappings config.TypeScriptMappings, scalars map[string]*model.Scalar, structs map[string]*model.Struct, ts *Code) error {
	ts.L("// " + str.FullName())
	switch {
	case str.Array != nil:
		ts.App("export type " + str.Name + " = Array<")
		valueTSType(str.Array.Value, mappings, scalars, structs, ts, nil)
		ts.App(">")
		ts.NL()
	case str.Map != nil:
		ts.App("export type " + str.Name + " = Record<")
		if str.Map.Key != nil {
			valueTSType(str.Map.Key, mappings, scalars, structs, ts, nil)
		} else {
			ts.App(str.Map.KeyType)
		}
		ts.App(",")
		valueTSType(str.Map.Value, mappings, scalars, structs, ts, nil)
		ts.App(">")
		ts.NL()
	case len(str.UnionFields) > 0:
		if len(str.Fields) > 0 || len(str.InlineFields) > 0 {
			return errors.New("no fields or inline fields are supported when using union")
		}
		switch {
		case str.UnionFields[0].Value.StructType != nil:
			ts.App("export type " + str.Name + " = ")
			var isUndefined bool
			for i, unionField := range str.UnionFields {
				if i > 0 {
					ts.App(" | ")
				}
				valueTSType(unionField.Value, mappings, scalars, structs, ts, &model.JSONInfo{OmitEmpty: true})
				if unionField.Value.IsPtr {
					isUndefined = true
				}
			}
			if isUndefined {
				ts.App(" | undefined")
			}
			ts.NL()
		case str.UnionFields[0].Value.Scalar != nil:
			ts.App("export const " + str.Name + " = ")
			ts.App("{ ")
			for i, field := range str.UnionFields {
				if i > 0 {
					ts.App(", ")
				}
				ts.App("...")
				valueTSType(field.Value, mappings, scalars, structs, ts, &model.JSONInfo{OmitEmpty: true})
			}
			ts.App(" }")
			ts.NL()
			ts.App("export type " + str.Name + " = typeof " + str.Name)
			ts.NL()
		default:
			return errors.New("could not resolve this union type")
		}
	case len(str.InlineFields) > 0:
		ts.App("export interface " + str.Name + " extends ")
		for i, inlineField := range str.InlineFields {
			if i > 0 {
				ts.App(", ")
			}
			if inlineField.Value.IsPtr {
				ts.App("Partial<")
			}
			valueTSType(inlineField.Value, mappings, scalars, structs, ts, &model.JSONInfo{OmitEmpty: true})
			if inlineField.Value.IsPtr {
				ts.App(">")
			}
			ts.App(" ")
		}
		ts.App("{")
		ts.NL()
		ts.Ind(1)
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
		ts.Ind(-1).L("}")
	default:
		ts.L("export interface " + str.Name + " {").Ind(1)
		renderStructFields(str.Fields, mappings, scalars, structs, ts)
		ts.Ind(-1).L("}")
	}
	return nil
}

func RenderTypescriptStructsToPackages(
	structs map[string]*model.Struct,
	mappings config.TypeScriptMappings,
	constantTypes map[string]map[string]interface{},
	scalarTypes map[string]*model.Scalar,
	mappedTypeScript map[string]map[string]*Code,
) (err error) {
	codeMap := map[string]map[string]*Code{}
	for _, mapping := range mappings {
		codeMap[mapping.GoPackage] = map[string]*Code{}
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
		packageCodeMap[str.Name] = NewCode("	")
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
				packageCodeMap[packageConstantTypeName] = NewCode("	")
				packageCodeMap[packageConstantTypeName].L("// " + packageName + "." + packageConstantTypeName)

				if packageConstantTypeValuesList, ok := packageConstantTypeValues.(map[string]*ast.BasicLit); ok {
					keys := make([]string, 0, len(packageConstantTypeValuesList))
					for k := range packageConstantTypeValuesList {
						keys = append(keys, k)
					}
					sort.Strings(keys)
					packageCodeMap[packageConstantTypeName].L("export enum " + packageConstantTypeName + " {").Ind(1)
					for _, k := range keys {
						enum := strings.TrimPrefix(strcase.ToCamel(k), packageConstantTypeName)
						packageCodeMap[packageConstantTypeName].L(enum + " = " + packageConstantTypeValuesList[k].Value + ",")
					}
					packageCodeMap[packageConstantTypeName].Ind(-1).L("}")
				} else if packageConstantTypeValuesString, ok := packageConstantTypeValues.(string); ok {
					packageCodeMap[packageConstantTypeName].L("export type " + packageConstantTypeName + " = " + packageConstantTypeValuesString)
				}
			}
		}
	}
	ensureCodeInPackage := func(goPackage string) {
		_, ok := mappedTypeScript[goPackage]
		if !ok {
			mappedTypeScript[goPackage] = map[string]*Code{}
		}
	}
	for _, mapping := range mappings {
		for structName, structCode := range codeMap[mapping.GoPackage] {
			ensureCodeInPackage(mapping.GoPackage)
			mappedTypeScript[mapping.GoPackage][structName] = structCode
		}
	}
	return nil
}

func Split(str string, seps []string) []string {
	var res []string
	strs := []string{str}
	for _, sep := range seps {
		var nextStrs []string
		for _, str := range strs {
			nextStrs = append(nextStrs, strings.Split(str, sep)...)
		}
		strs = nextStrs
		res = nextStrs
	}
	return res
}

func RenderTypeScriptServices(services model.ServiceList, mappings config.TypeScriptMappings, scalars map[string]*model.Scalar, structs map[string]*model.Struct, target *config.Target) (typeScript string, err error) {
	ts := NewCode("	")
	for _, service := range services {
		if !target.IsTSRPC(service.Name) {
			continue
		}
		err = renderTypescriptClient(service, mappings, scalars, structs, ts)
		if err != nil {
			return
		}
	}
	typeScript = ts.String()
	return
}

func GetTSHeaderComment() string {
	return "/* eslint:disable */\n// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2 - DO NOT EDIT.\n"
}
