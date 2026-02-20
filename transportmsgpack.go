package gotsrpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

var msgpackClientHandle = &clientHandle{
	contentType: "application/msgpack; charset=utf-8",
	handle:      &codec.MsgpackHandle{},
	// transform the error type to sth that is transportable
	beforeEncodeReply: defaultBeforeEncodeReply,
	beforeDecodeReply: defaultBeforeDecodeReply,
	afterDecodeReply:  defaultAfterDecodeReply,
}

func init() {
	mh := new(codec.MsgpackHandle)
	// use map[string]interface{} instead of map[interface{}]interface{}
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
	// attempting to set the promoted field in literal will cause a compiler error
	mh.RawToString = true
	msgpackClientHandle.handle = mh
	// WriteExt is not being called
	// if err := SetJSONExt(time.Time{}, 2, timeExt); err != nil {
	// 	 panic(err)
	// }
}

func NewMSGPackEncoderBytes(b *[]byte) *codec.Encoder {
	return codec.NewEncoderBytes(b, msgpackClientHandle.handle)
}

func NewMSGPackDecoderBytes(b []byte) *codec.Decoder {
	return codec.NewDecoderBytes(b, msgpackClientHandle.handle)
}

func SetMSGPackExt(rt interface{}, tag uint64, ext codec.BytesExt) error {
	if value, ok := msgpackClientHandle.handle.(*codec.MsgpackHandle); ok {
		return value.SetBytesExt(reflect.TypeOf(rt), tag, ext)
	}
	return errors.New("invalid handle type")
}
