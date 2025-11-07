package gotsrpc

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

type JSONInfo struct {
	Name      string
	Type      string
	Union     bool
	Inline    bool
	OmitEmpty bool
	Ignore    bool
}

type StructType struct {
	Name    string
	Package string
}

type Value struct {
	JSONInfo     *JSONInfo   `json:",omitempty"`
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
	Key       *Value
	Value     *Value
	KeyType   string
	KeyGoType string
}

type Field struct {
	Value    *Value
	Name     string    `json:",omitempty"`
	JSONInfo *JSONInfo `json:",omitempty"`
}

type Service struct {
	Name        string
	Methods     ServiceMethods
	Endpoint    string
	IsInterface bool
}

type ServiceMethods []*Method

type ServiceList []*Service

func (sm ServiceMethods) Len() int           { return len(sm) }
func (sm ServiceMethods) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm ServiceMethods) Less(i, j int) bool { return sm[i].Name < sm[j].Name }

type Method struct {
	Name   string
	Args   []*Field
	Return []*Field
}

type Struct struct {
	IsError      bool
	Package      string
	Name         string
	Fields       []*Field
	UnionFields  []*Field
	InlineFields []*Field
	Map          *Map
	Array        *Array
}

type Scalar struct {
	Name    string
	Package string
	Type    ScalarType
}

func (st *Scalar) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}
	return fullName
}
