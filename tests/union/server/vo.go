package server

import (
	"github.com/foomo/gotsrpc/v2"
)

func init() {
	gotsrpc.MustRegisterUnionExt(UnionString{}, UnionStruct{})
}

type (
	InlineStructA struct {
		ValueA string `json:"valueA"`
	}
	InlineStructB struct {
		ValueB string `json:"valueB"`
	}
	InlineStruct struct {
		InlineStructA `json:",inline"`
		InlineStructB `json:",inline"`
		Value         string `json:"value"`
	}
	InlineStructPtr struct {
		*InlineStructA `json:",inline,omitempty"`
		*InlineStructB `json:",inline,omitempty"`
		Bug            *InlineStructB `json:"bug,omitempty"`
		Value          string         `json:"value"`
	}
)

type (
	UnionString struct {
		A *UnionStringA `json:"a,omitempty" gotsrpc:"union"`
		B *UnionStringB `json:"b,omitempty" gotsrpc:"union"`
	}
	UnionStringA string
	UnionStringB string
)

const (
	UnionStringAOne   UnionStringA = "one"
	UnionStringATwo   UnionStringA = "two"
	UnionStringBThree UnionStringB = "three"
	UnionStringBFour  UnionStringB = "four"
)

type (
	UnionStructA struct {
		Kind  string             `json:"kind" gotsrpc:"type:'UnionStructA'"`
		Value UnionStructAValueA `json:"value"`
		Bar   string             `json:"bar"`
	}
	UnionStructAValueA string

	UnionStructB struct {
		Kind  string             `json:"kind" gotsrpc:"type:'UnionStructB'"`
		Value UnionStructAValueB `json:"value"`
		Foo   string             `json:"foo"`
	}
	UnionStructAValueB string

	UnionStruct struct {
		A *UnionStructA `json:"a,omitempty" gotsrpc:"union"`
		B *UnionStructB `json:"b,omitempty" gotsrpc:"union"`
	}
)

const (
	UnionStructAValueAOne   UnionStructAValueA = "one"
	UnionStructAValueATwo   UnionStructAValueA = "two"
	UnionStructAValueAThree UnionStructAValueA = "three"

	UnionStructAValueBOne   UnionStructAValueB = "one"
	UnionStructAValueBTwo   UnionStructAValueB = "two"
	UnionStructAValueBThree UnionStructAValueB = "three"
)
