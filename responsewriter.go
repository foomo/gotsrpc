package gotsrpc

import (
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	wroteHeader bool
	status      int
}

func (r *ResponseWriter) WriteHeader(status int) {
	r.status = status
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(status)
}

func (r *ResponseWriter) Status() int {
	if !r.wroteHeader {
		return http.StatusOK
	}
	return r.status
}
