package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/types/server"
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

	t.Run("Bool", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Bool(true)
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int(42)
		require.NoError(t, clientErr)
		assert.Equal(t, 42, ret)
	})

	t.Run("Int64", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int64(int64(9876543210))
		require.NoError(t, clientErr)
		assert.Equal(t, int64(9876543210), ret)
	})

	t.Run("Float64", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Float64(3.14159)
		require.NoError(t, clientErr)
		assert.InDelta(t, 3.14159, ret, 1e-10)
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.String("hello world")
		require.NoError(t, clientErr)
		assert.Equal(t, "hello world", ret)
	})

	t.Run("StringPtr", func(t *testing.T) {
		t.Parallel()
		v := "test"
		ret, clientErr := c.StringPtr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("SimpleStruct", func(t *testing.T) {
		t.Parallel()
		v := server.Simple{Bool: true, Int: 42, Int64: 100, Float64: 2.718, String: "test"}
		ret, clientErr := c.SimpleStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("NestedStruct", func(t *testing.T) {
		t.Parallel()
		v := server.Nested{
			Name:  "parent",
			Child: server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "child"},
		}
		ret, clientErr := c.NestedStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringSlice", func(t *testing.T) {
		t.Parallel()
		v := []string{"a", "b", "c"}
		ret, clientErr := c.StringSlice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringStringMap", func(t *testing.T) {
		t.Parallel()
		v := map[string]string{"a": "1", "b": "2"}
		ret, clientErr := c.StringStringMap(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("MapOfMaps", func(t *testing.T) {
		t.Parallel()
		v := map[string]map[string]string{"outer": {"inner": "val"}}
		ret, clientErr := c.MapOfMaps(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("MultiArgs", func(t *testing.T) {
		t.Parallel()
		retA, retB, retC, clientErr := c.MultiArgs("hello", int64(42), true)
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", retA)
		assert.Equal(t, int64(42), retB)
		assert.True(t, retC)
	})

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Empty()
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("ByteSlice", func(t *testing.T) {
		t.Parallel()
		v := []byte("hello world")
		ret, clientErr := c.ByteSlice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("ObjectID", func(t *testing.T) {
		t.Parallel()
		var v server.ObjectID
		copy(v[:], "hello123456")
		ret, clientErr := c.ObjectID(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})
}
