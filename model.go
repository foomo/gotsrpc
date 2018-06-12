package gotsrpc

type ScalarType string

const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeByte   ScalarType = "byte"
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
	IsError      bool        `json:",omitempty"`
	IsInterface  bool        `json:",omitempty"`
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
	Name     string
	Methods  ServiceMethods
	Endpoint string
}

type ServiceMethods []*Method

type ServiceList []*Service

func (sm ServiceMethods) Len() int           { return len(sm) }
func (sm ServiceMethods) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm ServiceMethods) Less(i, j int) bool { return sm[i].Name < sm[j].Name }

func (sm ServiceMethods) getMethodByName(name string) *Method {
	for _, m := range sm {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (sl ServiceList) getServiceByName(name string) *Service {
	for _, s := range sl {
		if s.Name == name {
			return s
		}
	}
	return nil
}

type Method struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	IsError bool
	Package string
	Name    string
	Fields  []*Field
}

type Scalar struct {
	Name    string
	Package string
	Type    ScalarType
}
