package server

import (
	"net/http"

	"github.com/foomo/gotsrpc/v2/tests/union/private"
)

type Handler struct{}

func (h *Handler) InlineStruct(w http.ResponseWriter, r *http.Request) (e InlineStruct) {
	return InlineStruct{InlineStructA: InlineStructA{ValueA: "a"}}
}

func (h *Handler) InlineStructPtr(w http.ResponseWriter, r *http.Request) (e InlineStructPtr) {
	return InlineStructPtr{InlineStructA: &InlineStructA{ValueA: "a"}}
}

func (h *Handler) UnionString(w http.ResponseWriter, r *http.Request) (e UnionString) {
	return UnionString{B: new(UnionStringBThree)}
}

func (h *Handler) UnionStruct(w http.ResponseWriter, r *http.Request) (e UnionStruct) {
	return UnionStruct{B: &UnionStructB{Kind: "UnionStructB", Value: UnionStructAValueBOne}}
}

func (h *Handler) PrivateInlineStruct(w http.ResponseWriter, r *http.Request) (e *private.InlineStruct) {
	return &private.InlineStruct{InlineStructA: private.InlineStructA{ValueA: "a"}}
}

func (h *Handler) PrivateInlineStructPtr(w http.ResponseWriter, r *http.Request) (e *private.InlineStructPtr) {
	return &private.InlineStructPtr{InlineStructA: &private.InlineStructA{ValueA: "a"}}
}

func (h *Handler) PrivateUnionString(w http.ResponseWriter, r *http.Request) (e *PrivateUnionString) {
	return &PrivateUnionString{B: new(private.UnionStringBThree)}
}

func (h *Handler) PrivateUnionStruct(w http.ResponseWriter, r *http.Request) (e *PrivateUnionStruct) {
	return &PrivateUnionStruct{B: &private.UnionStructB{Kind: "UnionStructB", Value: private.UnionStructAValueBOne}}
}
