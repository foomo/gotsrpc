package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/context/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("Hello", func(t *testing.T) {
		msg, clientErr := c.Hello(t.Context(), "foomo")
		require.NoError(t, clientErr)
		assert.Equal(t, "Hello foomo", msg)
	})

	t.Run("Error", func(t *testing.T) {
		err, clientErr := c.Error(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.ErrorIs(t, err, server.ErrGo)
	})

	t.Run("PkgError", func(t *testing.T) {
		err, clientErr := c.PkgError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.ErrorIs(t, err, server.ErrPkg)
	})

	t.Run("JoinedError", func(t *testing.T) {
		err, clientErr := c.JoinedError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		require.ErrorIs(t, err, server.ErrGo)
		require.ErrorIs(t, err, server.ErrPkg)
	})

	t.Run("WrappedError", func(t *testing.T) {
		err, clientErr := c.WrappedError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		require.ErrorIs(t, err, server.ErrGo)
		require.ErrorIs(t, err, server.ErrPkg)
	})

	t.Run("CustomError", func(t *testing.T) {
		err, clientErr := c.CustomError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.Equal(t, "hello World", err.Error())

		var myErr *server.MyError
		require.ErrorAs(t, err, &myErr)
		assert.Equal(t, "hello World", myErr.Error())
	})
}
