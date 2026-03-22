package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nullable/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("VariantA", func(t *testing.T) {
		t.Parallel()
		v := server.Base{B1: "hello", D1: server.ACustomTypeOne}
		ret, clientErr := c.VariantA(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantB", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomType("test")
		ret, clientErr := c.VariantB(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantC", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomTypes{"a", "b"}
		ret, clientErr := c.VariantC(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantD", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomTypesMap{"x": "y"}
		ret, clientErr := c.VariantD(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("VariantE", func(t *testing.T) {
		t.Parallel()
		v := &server.Base{B1: "ptr"}
		ret, clientErr := c.VariantE(t.Context(), v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v.B1, ret.B1)
	})

	t.Run("VariantE_nil", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.VariantE(t.Context(), nil)
		require.NoError(t, clientErr)
		assert.Nil(t, ret)
	})

	t.Run("VariantF", func(t *testing.T) {
		t.Parallel()
		v := []*server.Base{{B1: "one"}, {B1: "two"}}
		ret, clientErr := c.VariantF(t.Context(), v)
		require.NoError(t, clientErr)
		require.Len(t, ret, 2)
		assert.Equal(t, "one", ret[0].B1)
		assert.Equal(t, "two", ret[1].B1)
	})

	t.Run("VariantG", func(t *testing.T) {
		t.Parallel()
		v := map[string]*server.Base{"k": {B1: "val"}}
		ret, clientErr := c.VariantG(t.Context(), v)
		require.NoError(t, clientErr)
		require.Contains(t, ret, "k")
		require.NotNil(t, ret["k"])
		assert.Equal(t, "val", ret["k"].B1)
	})

	t.Run("VariantH", func(t *testing.T) {
		t.Parallel()
		i1 := server.Base{B1: "one"}
		i2 := &server.Base{B1: "two"}
		i3 := []*server.Base{{B1: "three"}}
		i4 := map[string]server.Base{"k": {B1: "four"}}
		r1, r2, r3, r4, clientErr := c.VariantH(t.Context(), i1, i2, i3, i4)
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
