package gotsrpc

import (
	"net/http"
)

// ResponseWriter wraps an [http.ResponseWriter] and records the status code
// written. Set the embedded ResponseWriter before use.
type ResponseWriter struct {
	http.ResponseWriter
	wroteHeader bool
	status      int
}

// WriteHeader records status and forwards it to the wrapped writer.
func (r *ResponseWriter) WriteHeader(status int) {
	r.status = status
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(status)
}

// Status returns the status code written, or 200 OK if WriteHeader was not
// called.
func (r *ResponseWriter) Status() int {
	if !r.wroteHeader {
		return http.StatusOK
	}

	return r.status
}
