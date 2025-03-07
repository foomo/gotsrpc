package service

import (
	"net/http"
)

type Service interface {
	Context(w http.ResponseWriter, r *http.Request)
	Empty()
	Bool(v bool) bool
	BoolPtr(v bool) *bool
	Int(v int) int
	Int32(v int32) int32
	Int64(v int64) int64
	UInt(v uint) uint
	UInt32(v uint32) uint32
	UInt64(v uint64) uint64
	Float32(v float32) float32
	Float64(v float64) float64
	String(v string) string
	Struct(v Struct) Struct
	Interface(v interface{}) interface{}
	BoolSlice(v []bool) []bool
	IntSlice(v []int) []int
	Int32Slice(v []int32) []int32
	Int64Slice(v []int64) []int64
	UIntSlice(v []uint) []uint
	UInt32Slice(v []uint32) []uint32
	UInt64Slice(v []uint64) []uint64
	Float32Slice(v []float32) []float32
	Float64Slice(v []float64) []float64
	StringSlice(v []string) []string
	IntMap(v map[int]interface{}) map[int]interface{}
	Int32Map(v map[int32]interface{}) map[int32]interface{}
	Int64Map(v map[int64]interface{}) map[int64]interface{}
	UIntMap(v map[uint]interface{}) map[uint]interface{}
	UInt32Map(v map[uint32]interface{}) map[uint32]interface{}
	UInt64Map(v map[uint64]interface{}) map[uint64]interface{}
	Float32Map(v map[float32]interface{}) map[float32]interface{}
	Float64Map(v map[float64]interface{}) map[float64]interface{}
	StringMap(v map[string]interface{}) map[string]interface{}
	IntTypeMap(v map[IntTypeMapKey]IntTypeMapValue) map[IntTypeMapKey]IntTypeMapValue
	Int32TypeMap(v map[Int32TypeMapKey]Int32TypeMapValue) map[Int32TypeMapKey]Int32TypeMapValue
	Int64TypeMap(v map[Int64TypeMapKey]Int64TypeMapValue) map[Int64TypeMapKey]Int64TypeMapValue
	UIntTypeMap(v map[UIntTypeMapKey]UIntTypeMapValue) map[UIntTypeMapKey]UIntTypeMapValue
	UInt32TypeMap(v map[UInt32TypeMapKey]UInt32TypeMapValue) map[UInt32TypeMapKey]UInt32TypeMapValue
	UInt64TypeMap(v map[UInt64TypeMapKey]UInt64TypeMapValue) map[UInt64TypeMapKey]UInt64TypeMapValue
	Float32TypeMap(v map[Float32TypeMapKey]Float32TypeMapValue) map[Float32TypeMapKey]Float32TypeMapValue
	Float64TypeMap(v map[Float64TypeMapKey]Float64TypeMapValue) map[Float64TypeMapKey]Float64TypeMapValue
	StringTypeMap(v map[StringTypeMapKey]StringTypeMapValue) map[StringTypeMapKey]StringTypeMapValue
	IntTypeMapTyped(v IntTypeMapTyped) IntTypeMapTyped
	Int32TypeMapTyped(v Int32TypeMapTyped) Int32TypeMapTyped
	Int64TypeMapTyped(v Int64TypeMapTyped) Int64TypeMapTyped
	UIntTypeMapTyped(v UIntTypeMapTyped) UIntTypeMapTyped
	UInt32TypeMapTyped(v UInt32TypeMapTyped) UInt32TypeMapTyped
	UInt64TypeMapTyped(v UInt64TypeMapTyped) UInt64TypeMapTyped
	Float32TypeMapTyped(v Float32TypeMapTyped) Float32TypeMapTyped
	Float64TypeMapTyped(v Float64TypeMapTyped) Float64TypeMapTyped
	StringTypeMapTyped(v StringTypeMapTyped) StringTypeMapTyped
	InterfaceSlice(v []interface{}) []interface{}
	IntType(v IntType) IntType
	Int32Type(v Int32Type) Int32Type
	Int64Type(v Int64Type) Int64Type
	UIntType(v UIntType) UIntType
	UInt32Type(v UInt32Type) UInt32Type
	UInt64Type(v UInt64Type) UInt64Type
	Float32Type(v Float32Type) Float32Type
	Float64Type(v Float64Type) Float64Type
	StringType(v StringType) StringType
	NestedType(v NestedType) NestedType
}
