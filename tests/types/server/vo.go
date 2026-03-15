package server

type Simple struct {
	Bool    bool    `json:"bool"`
	Int     int     `json:"int"`
	Int64   int64   `json:"int64"`
	Float64 float64 `json:"float64"`
	String  string  `json:"string"`
}

type Nested struct {
	Name  string `json:"name"`
	Child Simple `json:"child"`
}

type WithPointers struct {
	StrPtr   *string `json:"strPtr"`
	Int64Ptr *int64  `json:"int64Ptr"`
	BoolPtr  *bool   `json:"boolPtr"`
	Child    *Simple `json:"child"`
}

type ObjectID [12]byte

type WithCollections struct {
	Strings   []string          `json:"strings"`
	Int64s    []int64           `json:"int64s"`
	Items     []Simple          `json:"items"`
	ItemPtrs  []*Simple         `json:"itemPtrs"`
	StringMap map[string]string `json:"stringMap"`
	StructMap map[string]Simple `json:"structMap"`
}
