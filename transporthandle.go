package gotsrpc

import (
	"io"
	"sync"

	"github.com/mitchellh/mapstructure"
	pkgerrors "github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

type transportHandle struct {
	handle            codec.Handle
	contentType       string
	encoderPool       sync.Pool
	decoderPool       sync.Pool
	beforeEncodeReply func(*[]any, []int) error
	beforeDecodeReply func([]any, []int) ([]any, error)
	afterDecodeReply  func(*[]any, []any, []int) error
}

func (ch *transportHandle) getEncoder(w io.Writer) *codec.Encoder {
	if enc, ok := ch.encoderPool.Get().(*codec.Encoder); ok {
		enc.Reset(w)
		return enc
	}

	return codec.NewEncoder(w, ch.handle)
}

func (ch *transportHandle) putEncoder(enc *codec.Encoder) {
	ch.encoderPool.Put(enc)
}

func (ch *transportHandle) getDecoder(r io.Reader) *codec.Decoder {
	if dec, ok := ch.decoderPool.Get().(*codec.Decoder); ok {
		dec.Reset(r)
		return dec
	}

	return codec.NewDecoder(r, ch.handle)
}

func (ch *transportHandle) putDecoder(dec *codec.Decoder) {
	ch.decoderPool.Put(dec)
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

func newErrorEncodeHook() func(*[]any, []int) error {
	return func(resp *[]any, errorIndices []int) error {
		for _, i := range errorIndices {
			if e, ok := (*resp)[i].(error); ok {
				(*resp)[i] = NewError(e)
			}
		}

		return nil
	}
}

func newErrorDecodeHook() func([]any, []int) ([]any, error) {
	return func(reply []any, errorIndices []int) ([]any, error) {
		if len(errorIndices) == 0 {
			return reply, nil
		}

		ret := make([]any, len(reply))
		copy(ret, reply)

		for _, i := range errorIndices {
			var e *Error

			ret[i] = e
		}

		return ret, nil
	}
}

func newErrorAfterDecodeHook() func(*[]any, []any, []int) error {
	return func(reply *[]any, wrappedReply []any, errorIndices []int) error {
		for _, i := range errorIndices {
			if e, ok := wrappedReply[i].(*Error); ok && e != nil {
				if y, ok := (*reply)[i].(*error); ok {
					*y = e
				} else if err := mapstructure.Decode(e.Data, (*reply)[i]); err != nil {
					return pkgerrors.Wrap(err, "failed to decode wrapped error")
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
