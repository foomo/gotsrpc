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
	beforeEncodeReply func(*[]any, bool) error
	beforeDecodeReply func([]any, bool) ([]any, error)
	afterDecodeReply  func(*[]any, []any, bool) error
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
	defaultBeforeEncodeReply = func(resp *[]any, lastIsError bool) error {
		if !lastIsError || len(*resp) == 0 {
			return nil
		}
		last := len(*resp) - 1
		if e, ok := (*resp)[last].(error); ok {
			(*resp)[last] = NewError(e)
		}
		return nil
	}
	defaultBeforeDecodeReply = func(reply []any, lastIsError bool) ([]any, error) {
		if !lastIsError || len(reply) == 0 {
			return reply, nil
		}
		ret := make([]any, len(reply))
		copy(ret, reply)
		var e *Error
		ret[len(ret)-1] = e
		return ret, nil
	}
	defaultAfterDecodeReply = func(reply *[]any, wrappedReply []any, lastIsError bool) error {
		if !lastIsError || len(wrappedReply) == 0 {
			return nil
		}
		last := len(wrappedReply) - 1
		if e, ok := wrappedReply[last].(*Error); ok && e != nil {
			if y, ok := (*reply)[last].(*error); ok {
				*y = e
			} else if err := mapstructure.Decode(e.Data, (*reply)[last]); err != nil {
				return pkgerrors.Wrap(err, "failed to decode wrapped error")
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
