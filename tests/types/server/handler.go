package server

import (
	"context"
)

type Handler struct{}

func (h *Handler) Bool(_ context.Context, v bool) bool          { return v }
func (h *Handler) Int(_ context.Context, v int) int             { return v }
func (h *Handler) Int8(_ context.Context, v int8) int8          { return v }
func (h *Handler) Int16(_ context.Context, v int16) int16       { return v }
func (h *Handler) Int32(_ context.Context, v int32) int32       { return v }
func (h *Handler) Int64(_ context.Context, v int64) int64       { return v }
func (h *Handler) Uint(_ context.Context, v uint) uint          { return v }
func (h *Handler) Uint8(_ context.Context, v uint8) uint8       { return v }
func (h *Handler) Uint16(_ context.Context, v uint16) uint16    { return v }
func (h *Handler) Uint32(_ context.Context, v uint32) uint32    { return v }
func (h *Handler) Uint64(_ context.Context, v uint64) uint64    { return v }
func (h *Handler) Float32(_ context.Context, v float32) float32 { return v }
func (h *Handler) Float64(_ context.Context, v float64) float64 { return v }
func (h *Handler) String(_ context.Context, v string) string    { return v }

func (h *Handler) StringPtr(_ context.Context, v *string) *string    { return v }
func (h *Handler) Int64Ptr(_ context.Context, v *int64) *int64       { return v }
func (h *Handler) BoolPtr(_ context.Context, v *bool) *bool          { return v }
func (h *Handler) Int8Ptr(_ context.Context, v *int8) *int8          { return v }
func (h *Handler) Int16Ptr(_ context.Context, v *int16) *int16       { return v }
func (h *Handler) Int32Ptr(_ context.Context, v *int32) *int32       { return v }
func (h *Handler) UintPtr(_ context.Context, v *uint) *uint          { return v }
func (h *Handler) Uint8Ptr(_ context.Context, v *uint8) *uint8       { return v }
func (h *Handler) Uint16Ptr(_ context.Context, v *uint16) *uint16    { return v }
func (h *Handler) Uint32Ptr(_ context.Context, v *uint32) *uint32    { return v }
func (h *Handler) Uint64Ptr(_ context.Context, v *uint64) *uint64    { return v }
func (h *Handler) Float32Ptr(_ context.Context, v *float32) *float32 { return v }

func (h *Handler) SimpleStruct(_ context.Context, v Simple) Simple                   { return v }
func (h *Handler) NestedStruct(_ context.Context, v Nested) Nested                   { return v }
func (h *Handler) StructWithPointers(_ context.Context, v WithPointers) WithPointers { return v }
func (h *Handler) StructWithCollections(_ context.Context, v WithCollections) WithCollections {
	return v
}

func (h *Handler) StringSlice(_ context.Context, v []string) []string       { return v }
func (h *Handler) Int64Slice(_ context.Context, v []int64) []int64          { return v }
func (h *Handler) SimpleSlice(_ context.Context, v []Simple) []Simple       { return v }
func (h *Handler) SimplePtrSlice(_ context.Context, v []*Simple) []*Simple  { return v }
func (h *Handler) StringSlice2D(_ context.Context, v [][]string) [][]string { return v }
func (h *Handler) Int8Slice(_ context.Context, v []int8) []int8             { return v }
func (h *Handler) Int16Slice(_ context.Context, v []int16) []int16          { return v }
func (h *Handler) Int32Slice(_ context.Context, v []int32) []int32          { return v }
func (h *Handler) UintSlice(_ context.Context, v []uint) []uint             { return v }
func (h *Handler) Uint16Slice(_ context.Context, v []uint16) []uint16       { return v }
func (h *Handler) Uint32Slice(_ context.Context, v []uint32) []uint32       { return v }
func (h *Handler) Uint64Slice(_ context.Context, v []uint64) []uint64       { return v }
func (h *Handler) Float32Slice(_ context.Context, v []float32) []float32    { return v }

func (h *Handler) StringStringMap(_ context.Context, v map[string]string) map[string]string { return v }
func (h *Handler) StringInt64Map(_ context.Context, v map[string]int64) map[string]int64    { return v }
func (h *Handler) StringSimpleMap(_ context.Context, v map[string]Simple) map[string]Simple { return v }
func (h *Handler) StringSimplePtrMap(_ context.Context, v map[string]*Simple) map[string]*Simple {
	return v
}
func (h *Handler) StringStringSliceMap(_ context.Context, v map[string][]string) map[string][]string {
	return v
}
func (h *Handler) StringInt8Map(_ context.Context, v map[string]int8) map[string]int8       { return v }
func (h *Handler) StringInt16Map(_ context.Context, v map[string]int16) map[string]int16    { return v }
func (h *Handler) StringInt32Map(_ context.Context, v map[string]int32) map[string]int32    { return v }
func (h *Handler) StringUintMap(_ context.Context, v map[string]uint) map[string]uint       { return v }
func (h *Handler) StringUint8Map(_ context.Context, v map[string]uint8) map[string]uint8    { return v }
func (h *Handler) StringUint16Map(_ context.Context, v map[string]uint16) map[string]uint16 { return v }
func (h *Handler) StringUint32Map(_ context.Context, v map[string]uint32) map[string]uint32 { return v }
func (h *Handler) StringUint64Map(_ context.Context, v map[string]uint64) map[string]uint64 { return v }
func (h *Handler) StringFloat32Map(_ context.Context, v map[string]float32) map[string]float32 {
	return v
}

func (h *Handler) MapOfMaps(_ context.Context, v map[string]map[string]string) map[string]map[string]string {
	return v
}
func (h *Handler) MapOfSimpleSlice(_ context.Context, v map[string][]Simple) map[string][]Simple {
	return v
}
func (h *Handler) SliceOfMaps(_ context.Context, v []map[string]string) []map[string]string {
	return v
}

func (h *Handler) MultiArgs(_ context.Context, a string, b int64, c bool) (string, int64, bool) {
	return a, b, c
}
func (h *Handler) MixedArgs(_ context.Context, s Simple, items []string, m map[string]int64) (Simple, []string, map[string]int64) {
	return s, items, m
}

func (h *Handler) AllScalarsStruct(_ context.Context, v AllScalars) AllScalars { return v }
func (h *Handler) AllScalarPointersStruct(_ context.Context, v AllScalarPointers) AllScalarPointers {
	return v
}
func (h *Handler) AllScalarSlicesStruct(_ context.Context, v AllScalarSlices) AllScalarSlices {
	return v
}
func (h *Handler) AllScalarMapsStruct(_ context.Context, v AllScalarMaps) AllScalarMaps { return v }

func (h *Handler) Empty(_ context.Context) bool                                      { return true }
func (h *Handler) ByteSlice(_ context.Context, v []byte) []byte                      { return v }
func (h *Handler) ObjectID(_ context.Context, v ObjectID) ObjectID                   { return v }
func (h *Handler) StringObjectID(_ context.Context, v StringObjectID) StringObjectID { return v }
