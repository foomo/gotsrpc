package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nested/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExtendedServiceGoRPCClient(t *testing.T) {
	t.Parallel()

	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	require.NoError(t, err)
	require.NoError(t, l.Close())

	s := server.NewExtendedServiceGoRPCProxy(l.Addr().String(), &server.Handler{}, nil)
	require.NoError(t, s.Start())
	t.Cleanup(s.Stop)

	c := server.NewExtendedServiceGoRPCClient(l.Addr().String(), nil)
	c.Start()
	t.Cleanup(c.Stop)

	t.Run("GetFirstName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetFirstName()
		require.NoError(t, clientErr)
		assert.Equal(t, "John", ret)
	})

	t.Run("GetMiddleName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetMiddleName()
		require.NoError(t, clientErr)
		assert.Equal(t, "Michael", ret)
	})

	t.Run("GetAge", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetAge()
		require.NoError(t, clientErr)
		assert.Equal(t, 30, ret)
	})

	t.Run("GetLastName", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetLastName()
		require.NoError(t, clientErr)
		assert.Equal(t, "Doe", ret)
	})

	t.Run("GetPerson", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.GetPerson()
		require.NoError(t, clientErr)
		assert.Equal(t, "John", ret.FirstName)
		assert.Equal(t, "Michael", ret.MiddleName)
		assert.Equal(t, "Doe", ret.LastName)
		assert.Equal(t, 30, ret.Age)
	})
}
