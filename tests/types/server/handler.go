package server

import (
	"context"
)

type Handler struct{}

func (h *Handler) Bool(_ context.Context, v bool) bool          { return v }
func (h *Handler) Int(_ context.Context, v int) int             { return v }
func (h *Handler) Int64(_ context.Context, v int64) int64       { return v }
func (h *Handler) Float64(_ context.Context, v float64) float64 { return v }
func (h *Handler) String(_ context.Context, v string) string    { return v }

func (h *Handler) StringPtr(_ context.Context, v *string) *string { return v }
func (h *Handler) Int64Ptr(_ context.Context, v *int64) *int64    { return v }
func (h *Handler) BoolPtr(_ context.Context, v *bool) *bool       { return v }

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

func (h *Handler) StringStringMap(_ context.Context, v map[string]string) map[string]string { return v }
func (h *Handler) StringInt64Map(_ context.Context, v map[string]int64) map[string]int64    { return v }
func (h *Handler) StringSimpleMap(_ context.Context, v map[string]Simple) map[string]Simple { return v }
func (h *Handler) StringSimplePtrMap(_ context.Context, v map[string]*Simple) map[string]*Simple {
	return v
}
func (h *Handler) StringStringSliceMap(_ context.Context, v map[string][]string) map[string][]string {
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

func (h *Handler) Empty(_ context.Context) bool                    { return true }
func (h *Handler) ByteSlice(_ context.Context, v []byte) []byte    { return v }
func (h *Handler) ObjectID(_ context.Context, v ObjectID) ObjectID { return v }
