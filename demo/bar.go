package demo

import (
	"net/http"

	"github.com/foomo/gotsrpc/v2/demo/nested"
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
	CustomTypeStringOne   CustomTypeString = "regular"
	CustomTypeStringTwo   CustomTypeString = "camelCase"
	CustomTypeStringThree CustomTypeString = "snake_case"
	CustomTypeStringFour CustomTypeString = "slug-case"
	CustomTypeStringFive CustomTypeString = "CONST_CASE"
	CustomTypeStringSix CustomTypeString = "SLUG-CASE-UPPER"
	CustomTypeStringSeven CustomTypeString = "dot.case"
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
)

type CustomError string

const (
	CustomErrorDemo CustomError = "demo"
)

var (
	ErrCustomDemo = NewCustomError(CustomErrorDemo)
)

func NewCustomError(e CustomError) *CustomError {
	return &e
}

func (e *CustomError) Error() string {
	return string(*e)
}

type Bar interface {
	Hello(w http.ResponseWriter, r *http.Request, number int64) int
	Repeat(one, two string) (three, four bool)
	Inheritance(inner Inner, nested OuterNested, inline OuterInline) (Inner, OuterNested, OuterInline)
	CustomType(customTypeInt CustomTypeInt, customTypeString CustomTypeString, CustomTypeStruct CustomTypeStruct) (*CustomTypeInt, *CustomTypeString, CustomTypeStruct)
	CustomError(one CustomError, two *CustomError) (three CustomError, four *CustomError)
}
