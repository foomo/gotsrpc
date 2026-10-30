package model

// ScalarType is the TypeScript category of a Go scalar type.
type ScalarType string

// Scalar types. ScalarTypeNone marks a value that is not a scalar.
const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeAny    ScalarType = "any"
	ScalarTypeByte   ScalarType = "byte"
	ScalarTypeError  ScalarType = "error"
	ScalarTypeNumber ScalarType = "number"
	ScalarTypeBool   ScalarType = "bool"
	ScalarTypeNone   ScalarType = ""
)
