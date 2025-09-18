package service

import (
	"net/http"
)

type Service interface {
	String(w http.ResponseWriter, r *http.Request, a string)
	Strings(w http.ResponseWriter, r *http.Request, a, b string)
}
