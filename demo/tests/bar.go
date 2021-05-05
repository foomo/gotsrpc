package tests

import (
	"net/http"

	"github.com/foomo/gotsrpc/v2/demo"
)

type Bar struct {}

func (b Bar) Hello(w http.ResponseWriter, r *http.Request, number int64) int {
	return int(number)
}

func (b Bar) Repeat(one, two string) (three, four bool) {
	return true, true
}

func (b Bar) Inheritance(inner demo.Inner, nested demo.OuterNested, inline demo.OuterInline) (demo.Inner, demo.OuterNested, demo.OuterInline) {
	return inner, nested, inline
}

func (b Bar) CustomType(customTypeInt demo.CustomTypeInt, customTypeString demo.CustomTypeString, customTypeStruct demo.CustomTypeStruct) (*demo.CustomTypeInt, *demo.CustomTypeString, demo.CustomTypeStruct) {
	a := demo.CustomTypeIntOne
	c := demo.CustomTypeStringOne
	return &a, &c, customTypeStruct
}

func (b Bar) CustomError(one demo.CustomError, two *demo.CustomError) (three demo.CustomError, four *demo.CustomError) {
	return one, two
}
