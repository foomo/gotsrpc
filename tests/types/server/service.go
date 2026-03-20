package server

import (
	"context"
)

type Service interface {
	// Scalars
	Bool(ctx context.Context, v bool) bool
	Int(ctx context.Context, v int) int
	Int8(ctx context.Context, v int8) int8
	Int16(ctx context.Context, v int16) int16
	Int32(ctx context.Context, v int32) int32
	Int64(ctx context.Context, v int64) int64
	Uint(ctx context.Context, v uint) uint
	Uint8(ctx context.Context, v uint8) uint8
	Uint16(ctx context.Context, v uint16) uint16
	Uint32(ctx context.Context, v uint32) uint32
	Uint64(ctx context.Context, v uint64) uint64
	Float32(ctx context.Context, v float32) float32
	Float64(ctx context.Context, v float64) float64
	String(ctx context.Context, v string) string

	// Pointers
	StringPtr(ctx context.Context, v *string) *string
	Int64Ptr(ctx context.Context, v *int64) *int64
	BoolPtr(ctx context.Context, v *bool) *bool
	Int8Ptr(ctx context.Context, v *int8) *int8
	Int16Ptr(ctx context.Context, v *int16) *int16
	Int32Ptr(ctx context.Context, v *int32) *int32
	UintPtr(ctx context.Context, v *uint) *uint
	Uint8Ptr(ctx context.Context, v *uint8) *uint8
	Uint16Ptr(ctx context.Context, v *uint16) *uint16
	Uint32Ptr(ctx context.Context, v *uint32) *uint32
	Uint64Ptr(ctx context.Context, v *uint64) *uint64
	Float32Ptr(ctx context.Context, v *float32) *float32

	// Structs
	SimpleStruct(ctx context.Context, v Simple) Simple
	NestedStruct(ctx context.Context, v Nested) Nested
	StructWithPointers(ctx context.Context, v WithPointers) WithPointers
	StructWithCollections(ctx context.Context, v WithCollections) WithCollections

	// Slices
	StringSlice(ctx context.Context, v []string) []string
	Int64Slice(ctx context.Context, v []int64) []int64
	SimpleSlice(ctx context.Context, v []Simple) []Simple
	SimplePtrSlice(ctx context.Context, v []*Simple) []*Simple
	StringSlice2D(ctx context.Context, v [][]string) [][]string
	Int8Slice(ctx context.Context, v []int8) []int8
	Int16Slice(ctx context.Context, v []int16) []int16
	Int32Slice(ctx context.Context, v []int32) []int32
	UintSlice(ctx context.Context, v []uint) []uint
	Uint16Slice(ctx context.Context, v []uint16) []uint16
	Uint32Slice(ctx context.Context, v []uint32) []uint32
	Uint64Slice(ctx context.Context, v []uint64) []uint64
	Float32Slice(ctx context.Context, v []float32) []float32

	// Maps
	StringStringMap(ctx context.Context, v map[string]string) map[string]string
	StringInt64Map(ctx context.Context, v map[string]int64) map[string]int64
	StringSimpleMap(ctx context.Context, v map[string]Simple) map[string]Simple
	StringSimplePtrMap(ctx context.Context, v map[string]*Simple) map[string]*Simple
	StringStringSliceMap(ctx context.Context, v map[string][]string) map[string][]string
	StringInt8Map(ctx context.Context, v map[string]int8) map[string]int8
	StringInt16Map(ctx context.Context, v map[string]int16) map[string]int16
	StringInt32Map(ctx context.Context, v map[string]int32) map[string]int32
	StringUintMap(ctx context.Context, v map[string]uint) map[string]uint
	StringUint8Map(ctx context.Context, v map[string]uint8) map[string]uint8
	StringUint16Map(ctx context.Context, v map[string]uint16) map[string]uint16
	StringUint32Map(ctx context.Context, v map[string]uint32) map[string]uint32
	StringUint64Map(ctx context.Context, v map[string]uint64) map[string]uint64
	StringFloat32Map(ctx context.Context, v map[string]float32) map[string]float32

	// Complex nested
	MapOfMaps(ctx context.Context, v map[string]map[string]string) map[string]map[string]string
	MapOfSimpleSlice(ctx context.Context, v map[string][]Simple) map[string][]Simple
	SliceOfMaps(ctx context.Context, v []map[string]string) []map[string]string

	// Multi-args
	MultiArgs(ctx context.Context, a string, b int64, c bool) (string, int64, bool)
	MixedArgs(ctx context.Context, s Simple, items []string, m map[string]int64) (Simple, []string, map[string]int64)

	// Struct with all scalars
	AllScalarsStruct(ctx context.Context, v AllScalars) AllScalars
	AllScalarPointersStruct(ctx context.Context, v AllScalarPointers) AllScalarPointers
	AllScalarSlicesStruct(ctx context.Context, v AllScalarSlices) AllScalarSlices
	AllScalarMapsStruct(ctx context.Context, v AllScalarMaps) AllScalarMaps

	// Edge cases
	Empty(ctx context.Context) bool
	ByteSlice(ctx context.Context, v []byte) []byte
	ObjectID(ctx context.Context, v ObjectID) ObjectID
	StringObjectID(ctx context.Context, v StringObjectID) StringObjectID
}
