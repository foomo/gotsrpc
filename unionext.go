package gotsrpc

import (
	"fmt"
	"reflect"
)

// UnionExt is a JSON codec extension that encodes a union struct as its
// first non-zero field. Register union types with [RegisterUnionExt].
type UnionExt struct{}

var unionExt = &UnionExt{}

// RegisterUnionExt registers [UnionExt] with the JSON codec for the type of
// each value in v. Generated proxies call it for every union struct.
func RegisterUnionExt(v ...any) error {
	for _, i := range v {
		if err := SetJSONExt(i, 1, unionExt); err != nil {
			return err
		}
	}

	return nil
}

// MustRegisterUnionExt is like [RegisterUnionExt] but panics on error.
func MustRegisterUnionExt(v ...any) {
	if err := RegisterUnionExt(v...); err != nil {
		panic(err)
	}
}

// ConvertExt returns the first non-zero field of the union struct v, or nil
// if all fields are zero.
func (x *UnionExt) ConvertExt(v any) any {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for _, field := range val.Fields() {
		if !field.IsZero() {
			return field.Interface()
		}
	}

	return nil
}

// UpdateExt does not decode src into dst, so union structs cannot be decoded
// through the JSON extension.
func (x *UnionExt) UpdateExt(dst any, src any) {
	fmt.Println("")
}
