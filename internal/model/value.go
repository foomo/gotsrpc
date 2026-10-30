package model

// Value describes the Go type of a field, argument or return value.
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
	TypeParam    string      `json:",omitempty"`
	TypeArgs     []*Value    `json:",omitempty"`
}
