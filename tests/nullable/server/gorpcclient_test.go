package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nullable/server"
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

	t.Run("VariantA", func(t *testing.T) {
		t.Parallel()
		v := server.Base{B1: "hello", D1: server.ACustomTypeOne}
		ret, clientErr := c.VariantA(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantB", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomType("test")
		ret, clientErr := c.VariantB(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantE", func(t *testing.T) {
		t.Parallel()
		v := &server.Base{B1: "ptr"}
		ret, clientErr := c.VariantE(v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v.B1, ret.B1)
	})

	t.Run("VariantH", func(t *testing.T) {
		t.Parallel()
		i1 := server.Base{B1: "one"}
		i2 := &server.Base{B1: "two"}
		i3 := []*server.Base{{B1: "three"}}
		i4 := map[string]server.Base{"k": {B1: "four"}}
		r1, r2, r3, r4, clientErr := c.VariantH(i1, i2, i3, i4)
		require.NoError(t, clientErr)
		assert.Equal(t, "one", r1.B1)
		require.NotNil(t, r2)
		assert.Equal(t, "two", r2.B1)
		require.Len(t, r3, 1)
		assert.Equal(t, "three", r3[0].B1)
		require.Contains(t, r4, "k")
		assert.Equal(t, "four", r4["k"].B1)
	})
}
