package gotsrpc

import (
	"io"
	"sync"

	"github.com/mitchellh/mapstructure"
	pkgerrors "github.com/pkg/errors"
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
	encoderPool       sync.Pool
	decoderPool       sync.Pool
	beforeEncodeReply func(*[]any, []int) error
	beforeDecodeReply func([]any, []int) ([]any, error)
	afterDecodeReply  func(*[]any, []any, []int) error
}

func (ch *clientHandle) getEncoder(w io.Writer) *codec.Encoder {
	if enc, ok := ch.encoderPool.Get().(*codec.Encoder); ok {
		enc.Reset(w)
		return enc
	}
	return codec.NewEncoder(w, ch.handle)
}

func (ch *clientHandle) putEncoder(enc *codec.Encoder) {
	ch.encoderPool.Put(enc)
}

func (ch *clientHandle) getDecoder(r io.Reader) *codec.Decoder {
	if dec, ok := ch.decoderPool.Get().(*codec.Decoder); ok {
		dec.Reset(r)
		return dec
	}
	return codec.NewDecoder(r, ch.handle)
}

func (ch *clientHandle) putDecoder(dec *codec.Decoder) {
	ch.decoderPool.Put(dec)
}

var (
	defaultBeforeEncodeReply = func(resp *[]any, errorIndices []int) error {
		for _, i := range errorIndices {
			if e, ok := (*resp)[i].(error); ok {
				(*resp)[i] = NewError(e)
			}
		}
		return nil
	}
	defaultBeforeDecodeReply = func(reply []any, errorIndices []int) ([]any, error) {
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
	defaultAfterDecodeReply = func(reply *[]any, wrappedReply []any, errorIndices []int) error {
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
