package model

type ScalarType string

const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeAny    ScalarType = "any"
	ScalarTypeByte   ScalarType = "byte"
	ScalarTypeError  ScalarType = "error"
	ScalarTypeNumber ScalarType = "number"
	ScalarTypeBool   ScalarType = "bool"
	ScalarTypeNone   ScalarType = ""
)
