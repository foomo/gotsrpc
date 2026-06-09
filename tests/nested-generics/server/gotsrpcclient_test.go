package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nested-generics/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetValue(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.ID)
		assert.Equal(t, "test", ret.Name)
	})

	t.Run("GetWrapped", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetWrapped(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "1", ret.Data.ID)
		assert.Equal(t, "test", ret.Data.Name)
	})

	t.Run("GetByKey", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetByKey(t.Context(), "hello")
		require.NoError(t, clientErr)
		assert.Equal(t, 5, ret)
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetName(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "service", ret)
	})
}
