package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nested-generics/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServiceGoRPCClient(t *testing.T) {
	t.Parallel()

	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	require.NoError(t, err)
	require.NoError(t, l.Close())

	s := server.NewServiceGoRPCProxy(l.Addr().String(), &server.Handler{}, nil)
	require.NoError(t, s.Start())
	t.Cleanup(s.Stop)

	c := server.NewServiceGoRPCClient(l.Addr().String(), nil)
	c.Start()
	t.Cleanup(c.Stop)

	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetValue()
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.ID)
		assert.Equal(t, "test", ret.Name)
	})

	t.Run("GetWrapped", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetWrapped()
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.Data.ID)
		assert.Equal(t, "test", ret.Data.Name)
	})

	t.Run("GetByKey", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetByKey("hello")
		require.NoError(t, clientErr)
		assert.Equal(t, 5, ret)
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetName()
		require.NoError(t, clientErr)
		assert.Equal(t, "service", ret)
	})
}
