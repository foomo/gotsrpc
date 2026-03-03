package main_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/example/context/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	svr := httptest.NewServer(service.NewDefaultServiceGoTSRPCProxy(&service.Handler{}))
	defer svr.Close()

	client := service.NewDefaultServiceGoTSRPCClient(svr.URL)

	t.Run("Hello", func(t *testing.T) {
		msg, clientErr := client.Hello(t.Context(), "foomo")
		require.NoError(t, clientErr)
		assert.Equal(t, "Hello foomo", msg)
	})

	t.Run("StandardError", func(t *testing.T) {
		err, clientErr := client.StandardError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.Equal(t, "something went wrong: hello World", err.Error())
	})

	t.Run("WrappedError", func(t *testing.T) {
		err, clientErr := client.WrappedError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.Equal(t, "hello World: something", err.Error())
	})

	t.Run("TypedError", func(t *testing.T) {
		err, clientErr := client.TypedError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		require.ErrorIs(t, err, service.ErrSomething)
		assert.Equal(t, "something", err.Error())
	})

	t.Run("CustomError", func(t *testing.T) {
		err, clientErr := client.CustomError(t.Context(), "hello World")
		require.NoError(t, clientErr)
		var myErr *service.MyError
		require.ErrorAs(t, err, &myErr)
		assert.Equal(t, "hello World: something", err.Error())
	})
}
