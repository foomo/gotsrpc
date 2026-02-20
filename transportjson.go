package gotsrpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

func init() {
	jh := new(codec.JsonHandle)
	jh.MapKeyAsString = true
	jh.TimeFormat = []string{"unixmilli"}
	jsonClientHandle.handle = jh
}

var jsonClientHandle = &clientHandle{
	handle:      &codec.JsonHandle{},
	contentType: "application/json; charset=utf-8",
	// transform the error type to sth that is transportable
	beforeEncodeReply: defaultBeforeEncodeReply,
	beforeDecodeReply: defaultBeforeDecodeReply,
	afterDecodeReply:  defaultAfterDecodeReply,
}

func SetJSONExt(rt interface{}, tag uint64, ext codec.InterfaceExt) error {
	if value, ok := jsonClientHandle.handle.(*codec.JsonHandle); ok {
		return value.SetInterfaceExt(reflect.TypeOf(rt), tag, ext)
	}
	return errors.New("invalid handle type")
}
