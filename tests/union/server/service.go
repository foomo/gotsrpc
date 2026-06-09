package server

import (
	"net/http"

	"github.com/foomo/gotsrpc/v2/tests/union/private"
)

type Service interface {
	InlineStruct(w http.ResponseWriter, r *http.Request) (e InlineStruct)
	InlineStructPtr(w http.ResponseWriter, r *http.Request) (e InlineStructPtr)
	UnionString(w http.ResponseWriter, r *http.Request) (e UnionString)
	UnionStruct(w http.ResponseWriter, r *http.Request) (e UnionStruct)
	PrivateInlineStruct(w http.ResponseWriter, r *http.Request) (e *private.InlineStruct)
	PrivateInlineStructPtr(w http.ResponseWriter, r *http.Request) (e *private.InlineStructPtr)
	PrivateUnionString(w http.ResponseWriter, r *http.Request) (e *PrivateUnionString)
	PrivateUnionStruct(w http.ResponseWriter, r *http.Request) (e *PrivateUnionStruct)
}
