package backend_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/example/errors/handler/backend"
	backendsvs "github.com/foomo/gotsrpc/v2/example/errors/service/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	s := httptest.NewServer(backendsvs.NewDefaultServiceGoTSRPCProxy(backend.New()))
	c := backendsvs.NewDefaultServiceGoTSRPCClient(s.URL)

	t.Run("Error", func(t *testing.T) {
		serviceErr, err := c.Error(t.Context())
		require.NoError(t, err)

		var expectedErr *gotsrpc.Error
		assert.ErrorAs(t, serviceErr, &expectedErr)
	})

	t.Run("Scalar", func(t *testing.T) {
		serviceErr, err := c.Scalar(t.Context())
		require.NoError(t, err)
		assert.NotNil(t, serviceErr)
	})

	t.Run("MultiScalar", func(t *testing.T) {
		serviceErr, err := c.MultiScalar(t.Context())
		require.NoError(t, err)
		assert.NotNil(t, serviceErr)
	})

	t.Run("WrappedError", func(t *testing.T) {
		serviceErr, err := c.WrappedError(t.Context())
		require.NoError(t, err)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}
	})

	t.Run("ScalarError", func(t *testing.T) {
		serviceErr, err := c.ScalarError(t.Context())
		require.NoError(t, err)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}

		var scalarErr *backend.ScalarError
		if assert.ErrorAs(t, serviceErr, &scalarErr) {
			t.Log(scalarErr.Error())
		}
	})

	t.Run("StructError", func(t *testing.T) {
		serviceErr, err := c.StructError(t.Context())
		require.NoError(t, err)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}

		var structErr backend.StructError
		if assert.ErrorAs(t, serviceErr, &structErr) {
			t.Log(structErr.Error())
		}
	})

	t.Run("CustomError", func(t *testing.T) {
		serviceErr, err := c.CustomError(t.Context())
		require.NoError(t, err)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			fmt.Println(gotsrpcErr)
		}

		var customErr *backend.CustomError
		if assert.ErrorAs(t, serviceErr, &customErr) {
			fmt.Println(customErr)
		}
	})

	t.Run("TypedError", func(t *testing.T) {
		serviceErr, err := c.TypedError(t.Context())
		require.NoError(t, err)

		assert.ErrorIs(t, serviceErr, backend.ErrTyped)
	})

	t.Run("TypedWrappedError", func(t *testing.T) {
		serviceErr, err := c.TypedWrappedError(t.Context())
		require.NoError(t, err)

		assert.ErrorIs(t, serviceErr, backend.ErrTyped)
	})

	t.Run("TypedCustomError", func(t *testing.T) {
		serviceErr, err := c.TypedCustomError(t.Context())
		require.NoError(t, err)

		assert.ErrorIs(t, serviceErr, backend.ErrCustom)
	})
}
