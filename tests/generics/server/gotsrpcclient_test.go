package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/generics/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("GetStringResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetStringResponse(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", ret.Data)
	})

	t.Run("GetItemResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetItemResponse(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.Data.ID)
		assert.Equal(t, "test", ret.Data.Name)
	})

	t.Run("SetItemResponse", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.SetItemResponse(t.Context(), server.Response[server.Item]{Data: server.Item{ID: "1", Name: "x"}})
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("GetPair", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPair(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", ret.First)
		assert.Equal(t, 42, ret.Second)
	})

	t.Run("GetPagedItems", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPagedItems(t.Context(), 1)
		require.NoError(t, clientErr)
		require.Len(t, ret.Items, 1)
		assert.Equal(t, "1", ret.Items[0].ID)
		assert.Equal(t, 1, ret.Total)
	})

	t.Run("GetNestedGeneric", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetNestedGeneric(t.Context())
		require.NoError(t, clientErr)
		require.Len(t, ret.Items, 1)
		assert.Equal(t, "key", ret.Items[0].First)
		assert.Equal(t, "1", ret.Items[0].Second.ID)
	})

	t.Run("GetResult", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetResult(t.Context())
		require.NoError(t, clientErr)
		require.NotNil(t, ret.Value)
		assert.Equal(t, "1", ret.Value.ID)
	})

	t.Run("GetContainer", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetContainer(t.Context())
		require.NoError(t, clientErr)
		require.Contains(t, ret.Data, "key")
		assert.Equal(t, "1", ret.Data["key"].ID)
		assert.Equal(t, "default", ret.Default.Name)
	})
}
