package server

import (
	"errors"
	"net/http"

	pkgerrors "github.com/pkg/errors"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Error(w http.ResponseWriter, r *http.Request) (e error) {
	return errors.New("error")
}

func (h *Handler) Scalar(w http.ResponseWriter, r *http.Request) (e *ScalarError) {
	s := ScalarOne
	return &s
}

func (h *Handler) MultiScalar(w http.ResponseWriter, r *http.Request) (e *MultiScalar) {
	return &MultiScalar{
		ScalarA: ScalarAOne,
	}
}

func (h *Handler) Struct(w http.ResponseWriter, r *http.Request) (e *StructError) {
	return &StructError{
		Message: "my custom scalar",
		Data:    "hello world",
	}
}

func (h *Handler) TypedError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrTyped
}

func (h *Handler) StructError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewMyStructError("struct error")
}

func (h *Handler) ScalarError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewMyScalarError(MyScalarErrorOne)
}

func (h *Handler) CustomError(w http.ResponseWriter, r *http.Request) (e error) {
	return NewMyCustomError("custom error")
}

func (h *Handler) WrappedError(w http.ResponseWriter, r *http.Request) (e error) {
	return pkgerrors.Wrap(errors.New("error"), "wrapped")
}

func (h *Handler) TypedWrappedError(w http.ResponseWriter, r *http.Request) (e error) {
	return pkgerrors.Wrap(ErrTyped, "wrapped")
}

func (h *Handler) TypedScalarError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrScalarTwo
}

func (h *Handler) TypedCustomError(w http.ResponseWriter, r *http.Request) (e error) {
	return ErrCustom
}
