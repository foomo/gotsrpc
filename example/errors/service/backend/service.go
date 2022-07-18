package backend

import (
	"net/http"
)

type ErrService string

const (
	ErrUnauthorized ErrService = "unauthorized"
)

type Scalar string

func (s Scalar) String() string {
	return string(s)
}

func (s *Scalar) Error() string {
	return s.String()
}

type Struct struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (s *Struct) Error() string {
	return s.Message
}

type (
	ScalarA     string
	ScalarB     string
	MultiScalar struct {
		ScalarA `json:",inline"`
		ScalarB `json:",inline"`
	}
)

const (
	ScalarOne Scalar = "one"
	ScalarTwo Scalar = "two"

	ScalarAOne ScalarA = "one"
	ScalarATwo ScalarA = "two"

	ScalarBThree ScalarB = "three"
	ScalarBFour  ScalarB = "four"
)

type Service interface {
	Error(w http.ResponseWriter, r *http.Request) (e error)
	Scalar(w http.ResponseWriter, r *http.Request) (e *Scalar)
	MultiScalar(w http.ResponseWriter, r *http.Request) (e *MultiScalar)
	Struct(w http.ResponseWriter, r *http.Request) (e *Struct)
	TypedError(w http.ResponseWriter, r *http.Request) (e error)
	ScalarError(w http.ResponseWriter, r *http.Request) (e error)
	CustomError(w http.ResponseWriter, r *http.Request) (e error)
	WrappedError(w http.ResponseWriter, r *http.Request) (e error)
	TypedWrappedError(w http.ResponseWriter, r *http.Request) (e error)
	TypedScalarError(w http.ResponseWriter, r *http.Request) (e error)
	TypedCustomError(w http.ResponseWriter, r *http.Request) (e error)
}
