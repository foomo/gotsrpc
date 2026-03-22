package server

import (
	"net/http"
)

type Service interface {
	Error(w http.ResponseWriter, r *http.Request) (e error)
	Errors(w http.ResponseWriter, r *http.Request) (e1 error, e2 error)
	Scalar(w http.ResponseWriter, r *http.Request) (e *ScalarError)
	MultiScalar(w http.ResponseWriter, r *http.Request) (e *MultiScalar)
	Struct(w http.ResponseWriter, r *http.Request) (e *StructError)
	StructError(w http.ResponseWriter, r *http.Request) (e error)
	TypedError(w http.ResponseWriter, r *http.Request) (e error)
	ScalarError(w http.ResponseWriter, r *http.Request) (e error)
	CustomError(w http.ResponseWriter, r *http.Request) (e error)
	WrappedError(w http.ResponseWriter, r *http.Request) (e error)
	TypedWrappedError(w http.ResponseWriter, r *http.Request) (e error)
	TypedScalarError(w http.ResponseWriter, r *http.Request) (e error)
	TypedCustomError(w http.ResponseWriter, r *http.Request) (e error)
}
