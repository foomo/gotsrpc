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

type AllScalars struct {
	Int8    int8    `json:"int8"`
	Int16   int16   `json:"int16"`
	Int32   int32   `json:"int32"`
	Uint    uint    `json:"uint"`
	Uint8   uint8   `json:"uint8"`
	Uint16  uint16  `json:"uint16"`
	Uint32  uint32  `json:"uint32"`
	Uint64  uint64  `json:"uint64"`
	Float32 float32 `json:"float32"`
}

type AllScalarPointers struct {
	Int8Ptr    *int8    `json:"int8Ptr"`
	Int16Ptr   *int16   `json:"int16Ptr"`
	Int32Ptr   *int32   `json:"int32Ptr"`
	UintPtr    *uint    `json:"uintPtr"`
	Uint8Ptr   *uint8   `json:"uint8Ptr"`
	Uint16Ptr  *uint16  `json:"uint16Ptr"`
	Uint32Ptr  *uint32  `json:"uint32Ptr"`
	Uint64Ptr  *uint64  `json:"uint64Ptr"`
	Float32Ptr *float32 `json:"float32Ptr"`
}

type AllScalarSlices struct {
	Int8s    []int8    `json:"int8s"`
	Int16s   []int16   `json:"int16s"`
	Int32s   []int32   `json:"int32s"`
	Uints    []uint    `json:"uints"`
	Uint16s  []uint16  `json:"uint16s"`
	Uint32s  []uint32  `json:"uint32s"`
	Uint64s  []uint64  `json:"uint64s"`
	Float32s []float32 `json:"float32s"`
}

type AllScalarMaps struct {
	Int8Map    map[string]int8    `json:"int8Map"`
	Int16Map   map[string]int16   `json:"int16Map"`
	Int32Map   map[string]int32   `json:"int32Map"`
	UintMap    map[string]uint    `json:"uintMap"`
	Uint8Map   map[string]uint8   `json:"uint8Map"`
	Uint16Map  map[string]uint16  `json:"uint16Map"`
	Uint32Map  map[string]uint32  `json:"uint32Map"`
	Uint64Map  map[string]uint64  `json:"uint64Map"`
	Float32Map map[string]float32 `json:"float32Map"`
}

type (
	ObjectID       [12]byte
	StringObjectID struct {
		ObjectID ObjectID `json:"objectId" gotsrpc:"type:string"`
	}
)

type WithCollections struct {
	Strings   []string          `json:"strings"`
	Int64s    []int64           `json:"int64s"`
	Items     []Simple          `json:"items"`
	ItemPtrs  []*Simple         `json:"itemPtrs"`
	StringMap map[string]string `json:"stringMap"`
	StructMap map[string]Simple `json:"structMap"`
}
