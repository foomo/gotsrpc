package service

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) String(w http.ResponseWriter, r *http.Request, a string)     {}
func (h *Handler) Strings(w http.ResponseWriter, r *http.Request, a, b string) {}
