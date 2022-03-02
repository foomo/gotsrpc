package service

import (
	"net/http"
)

type Service interface {
	InlineStruct(w http.ResponseWriter, r *http.Request) (e InlineStruct)
	InlineStructPtr(w http.ResponseWriter, r *http.Request) (e InlineStructPtr)
	UnionString(w http.ResponseWriter, r *http.Request) (e UnionString)
	UnionStruct(w http.ResponseWriter, r *http.Request) (e UnionStruct)
}
