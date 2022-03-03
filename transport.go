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
	handle            codec.Handle
	contentType       string
	beforeEncodeReply func(*[]interface{}) error
	beforeDecodeReply func([]interface{}) ([]interface{}, error)
	afterDecodeReply  func(*[]interface{}, []interface{}) error
}

var msgpackClientHandle = &clientHandle{
	contentType: "application/msgpack; charset=utf-8",
	handle:      &codec.MsgpackHandle{},
	// transform error type to sth that is transportable
	beforeEncodeReply: func(resp *[]interface{}) error {
		for k, v := range *resp {
			if e, ok := v.(error); ok {
				if !reflect.ValueOf(e).IsNil() {
					(*resp)[k] = NewError(e)
				}
			}
		}
		return nil
	},
	beforeDecodeReply: func(reply []interface{}) ([]interface{}, error) {
		ret := make([]interface{}, len(reply))
		for k, v := range reply {
			if _, ok := v.(*error); ok {
				var e *Error
				ret[k] = e
			} else {
				ret[k] = v
			}
		}
		return ret, nil
	},
	afterDecodeReply: func(reply *[]interface{}, wrappedReply []interface{}) error {
		for k, v := range wrappedReply {
			if x, ok := v.(*Error); ok && x != nil {
				if y, ok := (*reply)[k].(*error); ok {
					*y = x
				}
			}
		}
		return nil
	},
}

func init() {
	mh := new(codec.MsgpackHandle)
	// use map[string]interface{} instead of map[interface{}]interface{}
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
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
