package gotsrpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

var jsonHandle = &transportHandle{
	handle:            &codec.JsonHandle{},
	contentType:       "application/json; charset=utf-8",
	beforeEncodeReply: newErrorEncodeHook(),
	beforeDecodeReply: newErrorDecodeHook(),
	afterDecodeReply:  newErrorAfterDecodeHook(),
}

func init() {
	jh := new(codec.JsonHandle)
	jh.MapKeyAsString = true
	jh.TimeFormat = []string{"unixmilli"}
	jh.ReaderBufferSize = 4096
	jh.WriterBufferSize = 4096
	jsonHandle.handle = jh

	registerTransportHandle(EncodingJson, jsonHandle)
	setDefaultTransportHandle(jsonHandle)
}

// SetJSONExt registers ext for the type of rt with the JSON codec, using tag
// as the extension tag. It is not safe for concurrent use with encoding; call
// it during program initialization.
func SetJSONExt(rt any, tag uint64, ext codec.InterfaceExt) error {
	if value, ok := jsonHandle.handle.(*codec.JsonHandle); ok {
		return value.SetInterfaceExt(reflect.TypeOf(rt), tag, ext)
	}

	return errors.New("invalid handle type")
}
