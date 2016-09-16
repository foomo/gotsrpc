package gotsrpc

type ScalarType string

const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeNumber ScalarType = "number"
	ScalarTypeBool   ScalarType = "bool"
	ScalarTypeNone   ScalarType = ""
)

type JSONInfo struct {
	Name            string
	OmitEmpty       bool
	ForceStringType bool
	Ignore          bool
}

type StructType struct {
	Name    string
	Package string
}

type Value struct {
	Scalar       *Scalar     `json:",omitempty"`
	ScalarType   ScalarType  `json:",omitempty"`
	GoScalarType string      `json:",omitempty"`
	StructType   *StructType `json:",omitempty"`
	Struct       *Struct     `json:",omitempty"`
	Map          *Map        `json:",omitempty"`
	Array        *Array      `json:",omitempty"`
	IsPtr        bool        `json:",omitempty"`
}

type Array struct {
	Value *Value
}

type Map struct {
	Value   *Value
	KeyType string
}

type Field struct {
	Value    *Value
	Name     string    `json:",omitempty"`
	JSONInfo *JSONInfo `json:",omitempty"`
}

type Service struct {
	Name    string
	Methods []*Method
}

type Method struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	Package string
	Name    string
	Fields  []*Field
}

type Scalar struct {
	Name    string
	Package string
	Type    ScalarType
}
