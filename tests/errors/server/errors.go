package server

import (
	"github.com/pkg/errors"
)

type MyScalarError string

const (
	MyScalarErrorOne MyScalarError = "scalar error one"
	MyScalarErrorTwo MyScalarError = "scalar error two"
)

func NewMyScalarError(e MyScalarError) *MyScalarError {
	return &e
}

func (e *MyScalarError) Error() string {
	return string(*e)
}

type MyStructError struct {
	Msg    string
	Map    map[string]string
	Slice  []string
	Struct struct {
		A string
	}
}

func NewMyStructError(msg string) MyStructError {
	return MyStructError{
		Msg:    msg,
		Map:    map[string]string{"a": "b"},
		Slice:  []string{"a", "b"},
		Struct: struct{ A string }{A: "b"},
	}
}

func (e MyStructError) Error() string {
	return e.Msg
}

type MyCustomError struct {
	Msg    string
	Map    map[string]string
	Slice  []string
	Struct struct {
		A string
	}
}

func NewMyCustomError(msg string) *MyCustomError {
	return &MyCustomError{
		Msg:    msg,
		Map:    map[string]string{"a": "b"},
		Slice:  []string{"a", "b"},
		Struct: struct{ A string }{A: "b"},
	}
}

func (e *MyCustomError) Error() string {
	return e.Msg
}

var (
	ErrTyped     = errors.New("typed error")
	ErrCustom    = NewMyCustomError("typed custom error")
	ErrScalarOne = NewMyScalarError(MyScalarErrorOne)
	ErrScalarTwo = NewMyScalarError(MyScalarErrorTwo)
)
