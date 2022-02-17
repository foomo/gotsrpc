package service

type IntType int

const (
	IntAType IntType = 1
	IntBType IntType = 2
)

type Int32Type int32

const (
	Int32AType Int32Type = 1
	Int32BType Int32Type = 2
)

type Int64Type int64

const (
	Int64AType Int64Type = 1
	Int64BType Int64Type = 2
)

type UIntType int

const (
	UIntAType UIntType = 1
	UIntBType UIntType = 2
)

type UInt32Type uint32

const (
	UInt32AType UInt32Type = 1
	UInt32BType UInt32Type = 2
)

type UInt64Type uint64

const (
	UInt64AType UInt64Type = 1
	UInt64BType UInt64Type = 2
)

type Float32Type float32

const (
	Float32AType Float32Type = 1
	Float32BType Float32Type = 2
)

type Float64Type float64

const (
	Float64AType Float64Type = 1
	Float64BType Float64Type = 2
)

type StringType string

const (
	StringAType StringType = "A"
	StringBType StringType = "B"
)

type IntTypeMapKey int

const (
	IntATypeMapKey IntTypeMapKey = 1
	IntBTypeMapKey IntTypeMapKey = 2
)

type Int32TypeMapKey int32

const (
	Int32ATypeMapKey Int32TypeMapKey = 1
	Int32BTypeMapKey Int32TypeMapKey = 2
)

type Int64TypeMapKey int64

const (
	Int64ATypeMapKey Int64TypeMapKey = 1
	Int64BTypeMapKey Int64TypeMapKey = 2
)

type UIntTypeMapKey int

const (
	UIntATypeMapKey UIntTypeMapKey = 1
	UIntBTypeMapKey UIntTypeMapKey = 2
)

type UInt32TypeMapKey uint32

const (
	UInt32ATypeMapKey UInt32TypeMapKey = 1
	UInt32BTypeMapKey UInt32TypeMapKey = 2
)

type UInt64TypeMapKey uint64

const (
	UInt64ATypeMapKey UInt64TypeMapKey = 1
	UInt64BTypeMapKey UInt64TypeMapKey = 2
)

type Float32TypeMapKey float32

const (
	Float32ATypeMapKey Float32TypeMapKey = 1
	Float32BTypeMapKey Float32TypeMapKey = 2
)

type Float64TypeMapKey float64

const (
	Float64ATypeMapKey Float64TypeMapKey = 1
	Float64BTypeMapKey Float64TypeMapKey = 2
)

type StringTypeMapKey string

const (
	StringATypeMapKey StringTypeMapKey = "A"
	StringBTypeMapKey StringTypeMapKey = "B"
)

type IntTypeMapValue int

const (
	IntATypeMapValue IntTypeMapValue = 1
	IntBTypeMapValue IntTypeMapValue = 2
)

type Int32TypeMapValue int32

const (
	Int32ATypeMapValue Int32TypeMapValue = 1
	Int32BTypeMapValue Int32TypeMapValue = 2
)

type Int64TypeMapValue int64

const (
	Int64ATypeMapValue Int64TypeMapValue = 1
	Int64BTypeMapValue Int64TypeMapValue = 2
)

type UIntTypeMapValue int

const (
	UIntATypeMapValue UIntTypeMapValue = 1
	UIntBTypeMapValue UIntTypeMapValue = 2
)

type UInt32TypeMapValue uint32

const (
	UInt32ATypeMapValue UInt32TypeMapValue = 1
	UInt32BTypeMapValue UInt32TypeMapValue = 2
)

type UInt64TypeMapValue uint64

const (
	UInt64ATypeMapValue UInt64TypeMapValue = 1
	UInt64BTypeMapValue UInt64TypeMapValue = 2
)

type Float32TypeMapValue float32

const (
	Float32ATypeMapValue Float32TypeMapValue = 1
	Float32BTypeMapValue Float32TypeMapValue = 2
)

type Float64TypeMapValue float64

const (
	Float64ATypeMapValue Float64TypeMapValue = 1
	Float64BTypeMapValue Float64TypeMapValue = 2
)

type StringTypeMapValue string

const (
	StringATypeMapValue StringTypeMapValue = "A"
	StringBTypeMapValue StringTypeMapValue = "B"
)

type (
	IntTypeMapTyped     map[IntTypeMapKey]IntTypeMapValue
	Int32TypeMapTyped   map[Int32TypeMapKey]Int32TypeMapValue
	Int64TypeMapTyped   map[Int64TypeMapKey]Int64TypeMapValue
	UIntTypeMapTyped    map[UIntTypeMapKey]UIntTypeMapValue
	UInt32TypeMapTyped  map[UInt32TypeMapKey]UInt32TypeMapValue
	UInt64TypeMapTyped  map[UInt64TypeMapKey]UInt64TypeMapValue
	Float32TypeMapTyped map[Float32TypeMapKey]Float32TypeMapValue
	Float64TypeMapTyped map[Float64TypeMapKey]Float64TypeMapValue
	StringTypeMapTyped  map[StringTypeMapKey]StringTypeMapValue
)

type Struct struct {
	Int                 int
	Int32               int32
	Int64               int64
	UInt                uint
	UInt32              uint32
	UInt64              uint64
	Float32             float32
	Float64             float64
	String              string
	Interface           interface{}
	IntTypeMapTyped     map[IntTypeMapKey]IntTypeMapValue
	Int32TypeMapTyped   map[Int32TypeMapKey]Int32TypeMapValue
	Int64TypeMapTyped   map[Int64TypeMapKey]Int64TypeMapValue
	UIntTypeMapTyped    map[UIntTypeMapKey]UIntTypeMapValue
	UInt32TypeMapTyped  map[UInt32TypeMapKey]UInt32TypeMapValue
	UInt64TypeMapTyped  map[UInt64TypeMapKey]UInt64TypeMapValue
	Float32TypeMapTyped map[Float32TypeMapKey]Float32TypeMapValue
	Float64TypeMapTyped map[Float64TypeMapKey]Float64TypeMapValue
	StringTypeMapTyped  map[StringTypeMapKey]StringTypeMapValue
}
