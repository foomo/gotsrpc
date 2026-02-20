package gotsrpc

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

type ClientEncoding int

const (
	EncodingMsgpack = ClientEncoding(0)
	EncodingJson    = ClientEncoding(1) //nolint:staticcheck
)

type clientHandle struct {
	handle            codec.Handle
	contentType       string
	beforeEncodeReply func(*[]any) error
	beforeDecodeReply func([]any) ([]any, error)
	afterDecodeReply  func(*[]any, []any) error
}

var (
	errorType                = reflect.TypeOf((*error)(nil)).Elem()
	defaultBeforeEncodeReply = func(resp *[]any) error {
		for k, v := range *resp {
			if e, ok := v.(error); ok {
				if r := reflect.ValueOf(e); !r.IsZero() {
					(*resp)[k] = NewError(e)
				}
			}
		}
		return nil
	}
	defaultBeforeDecodeReply = func(reply []any) ([]any, error) {
		ret := make([]any, len(reply))
		for k, v := range reply {
			val := reflect.TypeOf(v)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}
			if val.Implements(errorType) {
				var e *Error
				ret[k] = e
			} else {
				ret[k] = v
			}
		}
		return ret, nil
	}
	defaultAfterDecodeReply = func(reply *[]any, wrappedReply []any) error {
		for k, v := range wrappedReply {
			if e, ok := v.(*Error); ok && e != nil {
				if y, ok := (*reply)[k].(*error); ok {
					*y = e
				} else if err := mapstructure.Decode(e.Data, (*reply)[k]); err != nil {
					return errors.Wrap(err, "failed to decode wrapped error")
				}
			}
		}
		return nil
	}
)

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
