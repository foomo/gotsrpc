package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/generics/server"
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

	t.Run("GetStringResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetStringResponse()
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", ret.Data)
	})

	t.Run("GetItemResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetItemResponse()
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.Data.ID)
		assert.Equal(t, "test", ret.Data.Name)
	})

	t.Run("SetItemResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.SetItemResponse(server.Response[server.Item]{Data: server.Item{ID: "1", Name: "x"}})
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("GetPair", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPair()
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", ret.First)
		assert.Equal(t, 42, ret.Second)
	})

	t.Run("GetPagedItems", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPagedItems(1)
		require.NoError(t, clientErr)
		require.Len(t, ret.Items, 1)
		assert.Equal(t, "1", ret.Items[0].ID)
		assert.Equal(t, 1, ret.Total)
	})

	t.Run("GetNestedGeneric", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetNestedGeneric()
		require.NoError(t, clientErr)
		require.Len(t, ret.Items, 1)
		assert.Equal(t, "key", ret.Items[0].First)
		assert.Equal(t, "1", ret.Items[0].Second.ID)
	})

	t.Run("GetResult", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetResult()
		require.NoError(t, clientErr)
		require.NotNil(t, ret.Value)
		assert.Equal(t, "1", ret.Value.ID)
	})

	t.Run("GetContainer", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetContainer()
		require.NoError(t, clientErr)
		require.Contains(t, ret.Data, "key")
		assert.Equal(t, "1", ret.Data["key"].ID)
		assert.Equal(t, "default", ret.Default.Name)
	})
}
