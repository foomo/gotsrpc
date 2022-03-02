package gotsrpc

import (
	"fmt"
	"reflect"
)

type UnionExt struct{}

var unionExt = &UnionExt{}

func RegisterUnionExt(v ...interface{}) error {
	for _, i := range v {
		if err := SetJSONExt(i, 1, unionExt); err != nil {
			return err
		}
	}
	return nil
}

func MustRegisterUnionExt(v ...interface{}) {
	if err := RegisterUnionExt(v...); err != nil {
		panic(err)
	}
}

func (x *UnionExt) ConvertExt(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		if field := val.Field(i); !field.IsZero() {
			return field.Interface()
		}
	}
	return nil
}

func (x *UnionExt) UpdateExt(dst interface{}, src interface{}) {
	fmt.Println("")
}
