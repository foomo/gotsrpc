package server_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/context/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	h := &server.Handler{}
	msg := "foomo"

	t.Run("Hello", func(t *testing.T) {
		assert.Equal(t, "Hello foomo", h.Hello(t.Context(), msg))
	})

	t.Run("Error", func(t *testing.T) {
		err := h.Error(t.Context(), msg)
		require.Error(t, err)
	})

	t.Run("PkgError", func(t *testing.T) {
		err := h.PkgError(t.Context(), msg)
		require.Error(t, err)
	})

	t.Run("JoinedError", func(t *testing.T) {
		err := h.JoinedError(t.Context(), msg)
		require.Error(t, err)
	})

	t.Run("WrappedError", func(t *testing.T) {
		err := h.WrappedError(t.Context(), msg)
		require.Error(t, err)
	})

	t.Run("CustomError", func(t *testing.T) {
		err := h.CustomError(t.Context(), msg)
		require.Error(t, err)
	})
}
