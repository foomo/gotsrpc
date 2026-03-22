package gotsrpc

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

type transportHandle struct {
	handle            codec.Handle
	contentType       string
	beforeEncodeReply func(*[]any) error
	beforeDecodeReply func([]any) ([]any, error)
	afterDecodeReply  func(*[]any, []any) error
}

// --------------------------------
// Registry
// --------------------------------

var (
	handlesByEncoding      = map[ClientEncoding]*transportHandle{}
	handlesByContentType   = map[string]*transportHandle{}
	defaultTransportHandle *transportHandle
)

func registerTransportHandle(encoding ClientEncoding, h *transportHandle) {
	handlesByEncoding[encoding] = h
	handlesByContentType[h.contentType] = h
	if defaultTransportHandle == nil {
		defaultTransportHandle = h
	}
}

func setDefaultTransportHandle(h *transportHandle) {
	defaultTransportHandle = h
}

// --------------------------------
// Shared hook constructors
// --------------------------------

func newErrorEncodeHook() func(*[]any) error {
	return func(resp *[]any) error {
		for k, v := range *resp {
			if e, ok := v.(error); ok {
				if r := reflect.ValueOf(e); !r.IsZero() {
					(*resp)[k] = NewError(e)
				}
			}
		}
		return nil
	}
}

func newErrorDecodeHook() func([]any) ([]any, error) {
	return func(reply []any) ([]any, error) {
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
}

func newErrorAfterDecodeHook() func(*[]any, []any) error {
	return func(reply *[]any, wrappedReply []any) error {
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
}

// --------------------------------
// Lookup functions
// --------------------------------

func getHandleForEncoding(encoding ClientEncoding) *transportHandle {
	if h, ok := handlesByEncoding[encoding]; ok {
		return h
	}
	return defaultTransportHandle
}

func getHandlerForContentType(contentType string) *transportHandle {
	if h, ok := handlesByContentType[contentType]; ok {
		return h
	}
	return defaultTransportHandle
}
