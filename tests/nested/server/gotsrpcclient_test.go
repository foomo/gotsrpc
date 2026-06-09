package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nested/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultExtendedServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultExtendedServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultExtendedServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("GetFirstName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetFirstName(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "John", ret)
	})

	t.Run("GetMiddleName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetMiddleName(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "Michael", ret)
	})

	t.Run("GetAge", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetAge(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, 30, ret)
	})

	t.Run("GetLastName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetLastName(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "Doe", ret)
	})

	t.Run("GetPerson", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPerson(t.Context())
		require.NoError(t, clientErr)
		assert.Equal(t, "John", ret.FirstName)
		assert.Equal(t, "Michael", ret.MiddleName)
		assert.Equal(t, "Doe", ret.LastName)
		assert.Equal(t, 30, ret.Age)
	})
}
