package gotsrpc

import (
	"net/http"
	"reflect"

	"github.com/ugorji/go/codec"
)

type ClientEncoding int

const (
	EncodingMsgpack = ClientEncoding(0)
	EncodingJson    = ClientEncoding(1)
)

type clientHandle struct {
	handle      codec.Handle
	contentType string
}

var msgpackClientHandle = &clientHandle{
	handle:      &codec.MsgpackHandle{},
	contentType: "application/msgpack; charset=utf-8",
}

func init() {
	mh := new(codec.MsgpackHandle)
	// use map[string]interface{} instead of map[interface{}]interface{}
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
	msgpackClientHandle.handle = mh
	// attempting to set promoted field in literal will cause a compiler error
	mh.RawToString = true
}

var jsonClientHandle = &clientHandle{
	handle: &codec.JsonHandle{
		MapKeyAsString: true,
	},
	contentType: "application/json; charset=utf-8",
}

func getHandleForEncoding(encoding ClientEncoding) *clientHandle {
	switch encoding {
	case EncodingMsgpack:
		return msgpackClientHandle
	case EncodingJson:
		return jsonClientHandle
	default:
		return jsonClientHandle
	}
}

func getHandlerForContentType(contentType string) *clientHandle {
	switch contentType {
	case msgpackClientHandle.contentType:
		return msgpackClientHandle
	case jsonClientHandle.contentType:
		return jsonClientHandle
	default:
		return jsonClientHandle
	}
}

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
