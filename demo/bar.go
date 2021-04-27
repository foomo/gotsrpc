package demo

import (
	"github.com/foomo/gotsrpc/demo/nested"
)

type (
	CustomTypeInt    int
	CustomTypeString string
	CustomTypeFoo    string
)

const (
	CustomTypeIntOne   CustomTypeInt = 1
	CustomTypeIntTwo   CustomTypeInt = 2
	CustomTypeIntThree CustomTypeInt = 3
)

const (
	CustomTypeStringOne   CustomTypeString = "one"
	CustomTypeStringTwo   CustomTypeString = "two"
	CustomTypeStringThree CustomTypeString = "three"
)

type (
	Inner struct {
		One string `json:"one"`
	}
	OuterNested struct {
		Inner Inner  `json:"inner"`
		Two   string `json:"two"`
	}
	OuterInline struct {
		Inner `json:",inline"`
		Two   string `json:"two"`
	}
	CustomTypeStruct struct {
		CustomTypeFoo    CustomTypeFoo
		CustomTypeInt    CustomTypeInt
		CustomTypeString CustomTypeString
		CustomTypeNested nested.CustomTypeNested
		Check            Check
	}
	//Check struct {
	//	Foo string
	//}
)

type Bar interface {
	//Hello(w http.ResponseWriter, r *http.Request, number int64) int
	//Repeat(one, two string) (three, four bool)
	//Inheritance(inner Inner, nested OuterNested, inline OuterInline) (Inner, OuterNested, OuterInline)
	CustomType(customTypeInt CustomTypeInt, customTypeString CustomTypeString, CustomTypeStruct CustomTypeStruct) (*CustomTypeInt, *CustomTypeString, CustomTypeStruct)
}
