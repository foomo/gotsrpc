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

	t.Run("Error", func(t *testing.T) {
		err, clientErr := client.Error(t.Context(), "hello World")
		require.NoError(t, clientErr)
		assert.Equal(t, "hello World", err.Error())
	})
}

func TestHandler(t *testing.T) {
	client := &service.Handler{}

	t.Run("Hello", func(t *testing.T) {
		msg := client.Hello(t.Context(), "foomo")
		assert.Equal(t, "Hello foomo", msg)
	})

	t.Run("Error", func(t *testing.T) {
		err := client.Error(t.Context(), "hello World")
		assert.Equal(t, "hello World: something", err.Error())
	})
}
