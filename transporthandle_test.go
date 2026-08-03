package gotsrpc //nolint:testpackage

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHandleForEncoding(t *testing.T) {
	t.Parallel()

	json := getHandleForEncoding(EncodingJson)
	require.NotNil(t, json)
	assert.Equal(t, "application/json; charset=utf-8", json.contentType)

	msgpack := getHandleForEncoding(EncodingMsgpack)
	require.NotNil(t, msgpack)
	assert.Equal(t, "application/msgpack; charset=utf-8", msgpack.contentType)

	// unknown encoding falls back to the default handle
	assert.Same(t, defaultTransportHandle, getHandleForEncoding(ClientEncoding(99)))
}

func TestGetHandlerForContentType(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "application/json; charset=utf-8",
		getHandlerForContentType("application/json; charset=utf-8").contentType)
	assert.Equal(t, "application/msgpack; charset=utf-8",
		getHandlerForContentType("application/msgpack; charset=utf-8").contentType)

	// unknown content type falls back to the default handle
	assert.Same(t, defaultTransportHandle, getHandlerForContentType("text/plain"))
}

func TestTransportHandle_EncoderDecoderPool(t *testing.T) {
	t.Parallel()

	for _, encoding := range []ClientEncoding{EncodingJson, EncodingMsgpack} {
		h := getHandleForEncoding(encoding)
		require.NotNil(t, h)

		var buf bytes.Buffer

		enc := h.getEncoder(&buf)
		require.NoError(t, enc.Encode("hello"))
		h.putEncoder(enc)

		// re-acquiring exercises the pooled path (Reset on the reused encoder)
		enc2 := h.getEncoder(&buf)
		require.NotNil(t, enc2)
		h.putEncoder(enc2)

		dec := h.getDecoder(&buf)

		var out string
		require.NoError(t, dec.Decode(&out))
		h.putDecoder(dec)
		assert.Equal(t, "hello", out)
	}
}

func TestNewErrorEncodeHook(t *testing.T) {
	t.Parallel()

	hook := newErrorEncodeHook()
	resp := []any{errors.New("boom"), "not-an-error"}
	require.NoError(t, hook(&resp, []int{0}))

	wrapped, ok := resp[0].(*Error)
	require.True(t, ok)
	assert.Equal(t, "boom", wrapped.Msg)
	// non-indexed entries are left untouched
	assert.Equal(t, "not-an-error", resp[1])
}

func TestNewErrorDecodeHook(t *testing.T) {
	t.Parallel()

	hook := newErrorDecodeHook()

	t.Run("empty indices passes through", func(t *testing.T) {
		t.Parallel()

		reply := []any{"a", "b"}
		out, err := hook(reply, nil)
		require.NoError(t, err)
		assert.Equal(t, reply, out)
	})

	t.Run("nils out error indices", func(t *testing.T) {
		t.Parallel()

		out, err := hook([]any{"a", "b"}, []int{1})
		require.NoError(t, err)
		require.Len(t, out, 2)
		assert.Equal(t, "a", out[0])
		assert.Nil(t, out[1])
	})
}

func TestNewErrorAfterDecodeHook(t *testing.T) {
	t.Parallel()

	hook := newErrorAfterDecodeHook()

	var target error

	reply := []any{&target}
	wrapped := []any{NewError(errors.New("boom"))}

	require.NoError(t, hook(&reply, wrapped, []int{0}))
	require.Error(t, target)
	assert.Equal(t, "boom", target.Error())
}
