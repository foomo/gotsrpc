package gotsrpc

import (
	"reflect"
	"time"

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
	jh.TimeNotBuiltin = true
	if err := jh.SetInterfaceExt(reflect.TypeOf(time.Time{}), 1, timeExt); err != nil {
		panic(err)
	}
	jsonHandle.handle = jh

	registerTransportHandle(EncodingJson, jsonHandle)
	setDefaultTransportHandle(jsonHandle)
}

func SetJSONExt(rt interface{}, tag uint64, ext codec.InterfaceExt) error {
	if value, ok := jsonHandle.handle.(*codec.JsonHandle); ok {
		return value.SetInterfaceExt(reflect.TypeOf(rt), tag, ext)
	}
	return errors.New("invalid handle type")
}
