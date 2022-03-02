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

//type TimeExt struct{}
//
//func (x TimeExt) WriteExt(v interface{}) []byte {
//	b := make([]byte, binary.MaxVarintLen64)
//	switch t := v.(type) {
//	case time.Time:
//		binary.PutVarint(b, t.UnixNano())
//		return b
//	case *time.Time:
//		binary.PutVarint(b, t.UnixNano())
//		return b
//	default:
//		panic("Bug")
//	}
//}
//func (x TimeExt) ReadExt(dest interface{}, src []byte) {
//	tt := dest.(*time.Time)
//	r := bytes.NewBuffer(src)
//	v, err := binary.ReadVarint(r)
//	if err != nil {
//		panic("BUG")
//	}
//	*tt = time.Unix(0, v).UTC()
//}

func init() {
	mh := new(codec.MsgpackHandle)
	// use map[string]interface{} instead of map[interface{}]interface{}
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
	//if err := mh.SetBytesExt(reflect.TypeOf(time.Time{}), 1, TimeExt{}); err != nil {
	//	panic("2")
	//}
	//mh.TimeNotBuiltin = true
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

func NewMSGPackEncoderBytes(b *[]byte) *codec.Encoder {
	return codec.NewEncoderBytes(b, msgpackClientHandle.handle)
}

func NewMSGPackDecoderBytes(b []byte) *codec.Decoder {
	return codec.NewDecoderBytes(b, msgpackClientHandle.handle)
}

func SetJSONExt(rt interface{}, tag uint64, ext codec.InterfaceExt) error {
	return jsonClientHandle.handle.(*codec.JsonHandle).SetInterfaceExt(reflect.TypeOf(rt), tag, ext)
}
func SetMSGPackExt(rt interface{}, tag uint64, ext codec.BytesExt) error {
	return msgpackClientHandle.handle.(*codec.MsgpackHandle).SetBytesExt(reflect.TypeOf(rt), tag, ext)
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
