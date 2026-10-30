package gotsrpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

var msgpackHandle = &transportHandle{
	contentType:       "application/msgpack; charset=utf-8",
	handle:            &codec.MsgpackHandle{},
	beforeEncodeReply: newErrorEncodeHook(),
	beforeDecodeReply: newErrorDecodeHook(),
	afterDecodeReply:  newErrorAfterDecodeHook(),
}

func init() {
	mh := new(codec.MsgpackHandle)
	// use map[string]interface{} instead of map[interface{}]interface{}
	mh.MapType = reflect.TypeFor[map[string]any]()
	// attempting to set the promoted field in literal will cause a compiler error
	mh.RawToString = true
	mh.ReaderBufferSize = 4096
	mh.WriterBufferSize = 4096
	msgpackHandle.handle = mh

	registerTransportHandle(EncodingMsgpack, msgpackHandle)
}

// NewMSGPackEncoderBytes returns an encoder that writes msgpack to *b with
// the same settings as the msgpack transport.
func NewMSGPackEncoderBytes(b *[]byte) *codec.Encoder {
	return codec.NewEncoderBytes(b, msgpackHandle.handle)
}

// NewMSGPackDecoderBytes returns a decoder that reads msgpack from b with the
// same settings as the msgpack transport.
func NewMSGPackDecoderBytes(b []byte) *codec.Decoder {
	return codec.NewDecoderBytes(b, msgpackHandle.handle)
}

// SetMSGPackExt registers ext for the type of rt with the msgpack codec,
// using tag as the extension tag. It is not safe for concurrent use with
// encoding; call it during program initialization.
func SetMSGPackExt(rt any, tag uint64, ext codec.BytesExt) error {
	if value, ok := msgpackHandle.handle.(*codec.MsgpackHandle); ok {
		return value.SetBytesExt(reflect.TypeOf(rt), tag, ext)
	}

	return errors.New("invalid handle type")
}
