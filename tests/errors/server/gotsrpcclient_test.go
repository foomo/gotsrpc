package server_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/tests/errors/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(server.New()))
	t.Cleanup(s.Close)

	c := server.NewDefaultServiceGoTSRPCClient(s.URL)

	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.Error(t.Context())
		require.NoError(t, clientErr)

		var expectedErr *gotsrpc.Error
		require.ErrorAs(t, serviceErr, &expectedErr)
	})

	t.Run("Errors", func(t *testing.T) {
		t.Parallel()
		serviceErr1, serviceErr2, clientErr := c.Errors(t.Context())
		require.NoError(t, clientErr)

		var expectedErr *gotsrpc.Error
		require.ErrorAs(t, serviceErr1, &expectedErr)

		require.ErrorAs(t, serviceErr2, &expectedErr)
	})

	t.Run("Scalar", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.Scalar(t.Context())
		require.NoError(t, clientErr)
		var err *server.ScalarError
		require.ErrorAs(t, serviceErr, &err)
		assert.Equal(t, "one", err.String())
	})

	t.Run("MultiScalar", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.MultiScalar(t.Context())
		require.NoError(t, clientErr)
		assert.NotNil(t, serviceErr)
	})

	t.Run("Struct", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.Struct(t.Context())
		require.NoError(t, clientErr)
		require.NotNil(t, serviceErr)
		assert.Equal(t, "my custom scalar", serviceErr.Message)
	})

	t.Run("TypedError", func(t *testing.T) {
		t.Parallel()
		err, clientErr := c.TypedError(t.Context())
		require.NoError(t, clientErr)
		assert.ErrorIs(t, err, server.ErrTyped)
	})

	t.Run("StructError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.StructError(t.Context())
		require.NoError(t, clientErr)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}

		var structErr server.MyStructError
		if assert.ErrorAs(t, serviceErr, &structErr) {
			t.Log(structErr.Error())
		}
	})

	t.Run("TypedError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.TypedError(t.Context())
		require.NoError(t, clientErr)

		assert.ErrorIs(t, serviceErr, server.ErrTyped)
	})

	t.Run("ScalarError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.ScalarError(t.Context())
		require.NoError(t, clientErr)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}

		var scalarErr *server.MyScalarError
		if assert.ErrorAs(t, serviceErr, &scalarErr) {
			t.Log(scalarErr.Error())
		}
	})

	t.Run("CustomError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.CustomError(t.Context())
		require.NoError(t, clientErr)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			fmt.Println(gotsrpcErr)
		}

		var customErr *server.MyCustomError
		if assert.ErrorAs(t, serviceErr, &customErr) {
			fmt.Println(customErr)
		}
	})

	t.Run("WrappedError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.WrappedError(t.Context())
		require.NoError(t, clientErr)

		var gotsrpcErr *gotsrpc.Error
		if assert.ErrorAs(t, serviceErr, &gotsrpcErr) {
			t.Log(gotsrpcErr.Error())
		}
	})

	t.Run("TypedWrappedError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.TypedWrappedError(t.Context())
		require.NoError(t, clientErr)

		assert.ErrorIs(t, serviceErr, server.ErrTyped)
	})

	t.Run("TypedScalarError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.TypedScalarError(t.Context())
		require.NoError(t, clientErr)

		assert.ErrorIs(t, serviceErr, server.ErrScalarTwo)
	})

	t.Run("TypedCustomError", func(t *testing.T) {
		t.Parallel()
		serviceErr, clientErr := c.TypedCustomError(t.Context())
		require.NoError(t, clientErr)

		assert.ErrorIs(t, serviceErr, server.ErrCustom)
	})
}
