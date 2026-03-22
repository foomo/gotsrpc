package server

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) InlineStruct(w http.ResponseWriter, r *http.Request) (e InlineStruct) {
	return InlineStruct{InlineStructA: InlineStructA{ValueA: "a"}}
}

func (h *Handler) InlineStructPtr(w http.ResponseWriter, r *http.Request) (e InlineStructPtr) {
	return InlineStructPtr{InlineStructA: &InlineStructA{ValueA: "a"}}
}

func (h *Handler) UnionString(w http.ResponseWriter, r *http.Request) (e UnionString) {
	v := UnionStringBThree
	return UnionString{B: &v}
}

func (h *Handler) UnionStruct(w http.ResponseWriter, r *http.Request) (e UnionStruct) {
	return UnionStruct{B: &UnionStructB{Kind: "UnionStructB", Value: UnionStructAValueBOne}}
}
