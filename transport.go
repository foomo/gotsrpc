package gotsrpc

import (
	"github.com/ugorji/go/codec"
	"net/http"
)

var (
	msgpackHandle      = &codec.MsgpackHandle{}
	msgpackContentType = "application/msgpack; charset=utf-8"
)

var (
	jsonHandle      = &codec.JsonHandle{}
	jsonContentType = "application/json; charset=utf-8"
)


type responseWriterWithLength struct {
	http.ResponseWriter
	length int
}

func newResponseWriterWithLength(w http.ResponseWriter) *responseWriterWithLength {
	return &responseWriterWithLength{w, 0}
}

func (w *responseWriterWithLength) Write(b []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(b)
	w.length += n
	return
}

func (w *responseWriterWithLength) Length() int {
	return w.length
}
