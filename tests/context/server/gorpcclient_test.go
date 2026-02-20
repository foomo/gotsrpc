package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/context/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServiceGoRPCClient(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	require.NoError(t, err)
	require.NoError(t, l.Close())

	s := server.NewServiceGoRPCProxy(l.Addr().String(), &server.Handler{}, nil)
	require.NoError(t, s.Start())
	t.Cleanup(s.Stop)

	c := server.NewServiceGoRPCClient(l.Addr().String(), nil)
	c.Start()
	t.Cleanup(c.Stop)

	t.Run("Hello", func(t *testing.T) {
		ret, clientErr := c.Hello("foomo")
		require.NoError(t, clientErr)
		assert.Equal(t, "Hello foomo", ret)
	})
}
