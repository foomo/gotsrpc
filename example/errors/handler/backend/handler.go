package backend

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/foomo/gotsrpc/v2/example/errors/service/backend"
)

type (
	ScalarError string
	StructError struct {
		Msg    string
		Map    map[string]string
		Slice  []string
		Struct struct {
			A string
		}
	}
	CustomError struct {
		Msg    string
		Map    map[string]string
		Slice  []string
		Struct struct {
			A string
		}
	}
)

const (
	ScalarErrorOne ScalarError = "scalar error one"
	ScalarErrorTwo ScalarError = "scalar error two"
)

func NewScalarError(e ScalarError) *ScalarError {
	return &e
}

func (e *ScalarError) Error() string {
	return string(*e)
}

func NewStructError(msg string) StructError {
	return StructError{
		Msg:    msg,
		Map:    map[string]string{"a": "b"},
		Slice:  []string{"a", "b"},
		Struct: struct{ A string }{A: "b"},
	}
}

func (e StructError) Error() string {
	return e.Msg
}

func NewCustomError(msg string) *CustomError {
	return &CustomError{
		Msg:    msg,
		Map:    map[string]string{"a": "b"},
		Slice:  []string{"a", "b"},
		Struct: struct{ A string }{A: "b"},
	}
}

func (e *CustomError) Error() string {
	return e.Msg
}

var (
	ErrTyped     = errors.New("typed error")
	ErrCustom    = NewCustomError("typed custom error")
	ErrScalarOne = NewScalarError(ScalarErrorOne)
	ErrScalarTwo = NewScalarError(ScalarErrorTwo)
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Error(w http.ResponseWriter, r *http.Request) (e error) {
	return errors.New("error")
}

func (h *Handler) Scalar(w http.ResponseWriter, r *http.Request) (e *backend.ScalarError) {
	s := backend.ScalarOne
	return &s
}

func (h *Handler) MultiScalar(w http.ResponseWriter, r *http.Request) (e *backend.MultiScalar) {
	return &backend.MultiScalar{
		ScalarA: backend.ScalarAOne,
	}
}

func (h *Handler) Struct(w http.ResponseWriter, r *http.Request) (e *backend.StructError) {
	return &backend.StructError{
		Message: "my custom scalar",
		Data:    "hello world",
	}
}

func (h *Handler) TypedError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrTyped
}

func (h *Handler) StructError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewStructError("struct error")
}

func (h *Handler) ScalarError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewScalarError(ScalarErrorOne)
}

func (h *Handler) CustomError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewCustomError("custom error")
}

func (h *Handler) WrappedError(w http.ResponseWriter, r *http.Request) (e error) {
	return errors.Wrap(errors.New("error"), "wrapped")
}

func (h *Handler) TypedWrappedError(w http.ResponseWriter, r *http.Request) (e error) {
	return errors.Wrap(ErrTyped, "wrapped")
}

func (h *Handler) TypedScalarError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrScalarTwo
}

func (h *Handler) TypedCustomError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrCustom
}
