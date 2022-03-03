package service

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
		*InlineStructB `json:",inline,omitempty"` // TODO this should have nil for InlineStructB as for Bug
		Bug            *InlineStructB             `json:"bug,omitempty"`
		Value          string                     `json:"value"`
	}
)

type (
	UnionString struct {
		A *UnionStringA `json:"a,omitempty,union"`

		B *UnionStringB `json:"b,omitempty,union"`
		c string
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
		Kind  string             `json:"kind,type:'UnionStructA'"`
		Value UnionStructAValueA `json:"value"`
		Bar   string             `json:"bar"`
	}
	UnionStructAValueA string

	UnionStructB struct {
		Kind  string             `json:"kind,type:'UnionStructB'"`
		Value UnionStructAValueB `json:"value"`
		Foo   string             `json:"foo"`
	}
	UnionStructAValueB string

	UnionStruct struct {
		A *UnionStructA `json:"a,omitempty,union"`
		B *UnionStructB `json:"b,omitempty,union"`
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