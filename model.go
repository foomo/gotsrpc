package gotsrpc

type ScalarType string

const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeNumber            = "number"
	ScalarTypeBool              = "bool"
	ScalarTypeNone              = ""
)

type JSONInfo struct {
	Name            string
	OmitEmpty       bool
	ForceStringType bool
	Ignore          bool
}

type Value struct {
	ScalarType ScalarType `json:",omitempty"`
	StructType string     `json:",omitempty"`
	Struct     *Struct    `json:",omitempty"`
	Map        *Map       `json:",omitempty"`
	Array      *Array     `json:",omitempty"`
	IsPtr      bool       `json:",omitempty"`
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

type Func struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	Name   string
	Fields map[string]*Field
}
