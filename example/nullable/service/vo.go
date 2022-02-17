package service

type (
	ACustomType     string
	ACustomTypes    []ACustomType
	ACustomTypesMap map[ACustomType]ACustomType
	BCustomType     string
	BCustomTypes    []BCustomType
	BCustomTypesMap map[BCustomType]BCustomType
)

const (
	ACustomTypeOne ACustomType = "one"
	ACustomTypeTwo ACustomType = "two"
)

type Base struct {
	A1 Nested  `json:"a1"`
	A2 *Nested `json:"a2,omitempty"`
	A3 *Nested `json:"a3"`

	B1 string  `json:"b1"`
	B2 *string `json:"b2,omitempty"`
	B3 *string `json:"b3"`

	C1 interface{}  `json:"c1"`
	C2 *interface{} `json:"c2,omitempty"`
	C3 *interface{} `json:"c3"`

	D1 ACustomType  `json:"d1"`
	D2 *ACustomType `json:"d2,omitempty"`
	D3 *ACustomType `json:"d3"`

	E1 ACustomTypes  `json:"e1"`
	E2 *ACustomTypes `json:"e2,omitempty"`
	E3 *ACustomTypes `json:"e3"`

	F1 ACustomTypesMap  `json:"f1"`
	F2 *ACustomTypesMap `json:"f2,omitempty"`
	F3 *ACustomTypesMap `json:"f3"`

	Two    []Nested
	Two1   [][]Nested
	Two2   []map[string]Nested
	Three  []*Nested
	Three1 []*string
	Four   map[string]Nested
	Five   map[string]*Nested
	Six    *struct {
		Foo string
	}
}

type Nested struct {
	Foo string
}
