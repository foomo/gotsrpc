package gotsrpc

import (
	"fmt"
	"reflect"
)

type UnionExt struct{}

var unionExt = &UnionExt{}

func RegisterUnionExt(v ...any) error {
	for _, i := range v {
		if err := SetJSONExt(i, 1, unionExt); err != nil {
			return err
		}
	}

	return nil
}

func MustRegisterUnionExt(v ...any) {
	if err := RegisterUnionExt(v...); err != nil {
		panic(err)
	}
}

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

func (x *UnionExt) UpdateExt(dst any, src any) {
	fmt.Println("")
}
