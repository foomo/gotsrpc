package server

import (
	"context"
)

type Service interface {
	// Scalars
	Bool(ctx context.Context, v bool) bool
	Int(ctx context.Context, v int) int
	Int64(ctx context.Context, v int64) int64
	Float64(ctx context.Context, v float64) float64
	String(ctx context.Context, v string) string

	// Pointers
	StringPtr(ctx context.Context, v *string) *string
	Int64Ptr(ctx context.Context, v *int64) *int64
	BoolPtr(ctx context.Context, v *bool) *bool

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

	// Maps
	StringStringMap(ctx context.Context, v map[string]string) map[string]string
	StringInt64Map(ctx context.Context, v map[string]int64) map[string]int64
	StringSimpleMap(ctx context.Context, v map[string]Simple) map[string]Simple
	StringSimplePtrMap(ctx context.Context, v map[string]*Simple) map[string]*Simple
	StringStringSliceMap(ctx context.Context, v map[string][]string) map[string][]string

	// Complex nested
	MapOfMaps(ctx context.Context, v map[string]map[string]string) map[string]map[string]string
	MapOfSimpleSlice(ctx context.Context, v map[string][]Simple) map[string][]Simple
	SliceOfMaps(ctx context.Context, v []map[string]string) []map[string]string

	// Multi-args
	MultiArgs(ctx context.Context, a string, b int64, c bool) (string, int64, bool)
	MixedArgs(ctx context.Context, s Simple, items []string, m map[string]int64) (Simple, []string, map[string]int64)

	// Edge cases
	Empty(ctx context.Context) bool
	ByteSlice(ctx context.Context, v []byte) []byte
	ObjectID(ctx context.Context, v ObjectID) ObjectID
}
