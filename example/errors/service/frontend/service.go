package frontend

import (
	"net/http"
)

const (
	ErrSimpleOne ErrSimple = "one"
	ErrSimpleTwo ErrSimple = "two"

	ErrMultiAOne ErrMultiA = "one"
	ErrMultiATwo ErrMultiA = "two"

	ErrMultiBThree ErrMultiB = "three"
	ErrMultiBFour  ErrMultiB = "four"
)

type Service interface {
	Simple(w http.ResponseWriter, r *http.Request) (e *ErrSimple)
	Multiple(w http.ResponseWriter, r *http.Request) (e *ErrMulti)
}
