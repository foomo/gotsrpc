package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/types/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("Bool", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Bool(t.Context(), true)
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int(t.Context(), 42)
		require.NoError(t, clientErr)
		assert.Equal(t, 42, ret)
	})

	t.Run("Int64", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int64(t.Context(), int64(9876543210))
		require.NoError(t, clientErr)
		assert.Equal(t, int64(9876543210), ret)
	})

	t.Run("Float64", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Float64(t.Context(), 3.14159)
		require.NoError(t, clientErr)
		assert.InDelta(t, 3.14159, ret, 1e-10)
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.String(t.Context(), "hello world")
		require.NoError(t, clientErr)
		assert.Equal(t, "hello world", ret)
	})

	t.Run("StringPtr", func(t *testing.T) {
		t.Parallel()
		v := "test"
		ret, clientErr := c.StringPtr(t.Context(), &v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Int64Ptr", func(t *testing.T) {
		t.Parallel()
		v := int64(42)
		ret, clientErr := c.Int64Ptr(t.Context(), &v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("BoolPtr", func(t *testing.T) {
		t.Parallel()
		v := true
		ret, clientErr := c.BoolPtr(t.Context(), &v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("SimpleStruct", func(t *testing.T) {
		t.Parallel()
		v := server.Simple{
			Bool: true, Int: 42, Int64: 100, Float64: 2.718, String: "test",
		}
		ret, clientErr := c.SimpleStruct(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("NestedStruct", func(t *testing.T) {
		t.Parallel()
		v := server.Nested{
			Name:  "parent",
			Child: server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "child"},
		}
		ret, clientErr := c.NestedStruct(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StructWithPointers", func(t *testing.T) {
		t.Parallel()
		str := "hello"
		i := int64(42)
		b := true
		child := server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "child"}
		v := server.WithPointers{
			StrPtr: &str, Int64Ptr: &i, BoolPtr: &b, Child: &child,
		}
		ret, clientErr := c.StructWithPointers(t.Context(), v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret.StrPtr)
		assert.Equal(t, str, *ret.StrPtr)
		require.NotNil(t, ret.Int64Ptr)
		assert.Equal(t, i, *ret.Int64Ptr)
		require.NotNil(t, ret.BoolPtr)
		assert.Equal(t, b, *ret.BoolPtr)
		require.NotNil(t, ret.Child)
		assert.Equal(t, child, *ret.Child)
	})

	t.Run("StructWithCollections", func(t *testing.T) {
		t.Parallel()
		v := server.WithCollections{
			Strings:   []string{"a", "b"},
			Int64s:    []int64{1, 2, 3},
			Items:     []server.Simple{{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "item"}},
			ItemPtrs:  []*server.Simple{{Bool: false, Int: 10, Int64: 20, Float64: 30.0, String: "ptr"}},
			StringMap: map[string]string{"key": "val"},
			StructMap: map[string]server.Simple{"s": {Bool: true, Int: 5, Int64: 6, Float64: 7.0, String: "map"}},
		}
		ret, clientErr := c.StructWithCollections(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.Strings, ret.Strings)
		assert.Equal(t, v.Int64s, ret.Int64s)
		assert.Equal(t, v.Items, ret.Items)
		require.Len(t, ret.ItemPtrs, 1)
		require.NotNil(t, ret.ItemPtrs[0])
		assert.Equal(t, *v.ItemPtrs[0], *ret.ItemPtrs[0])
		assert.Equal(t, v.StringMap, ret.StringMap)
		assert.Equal(t, v.StructMap, ret.StructMap)
	})

	t.Run("StringSlice", func(t *testing.T) {
		t.Parallel()
		v := []string{"a", "b", "c"}
		ret, clientErr := c.StringSlice(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Int64Slice", func(t *testing.T) {
		t.Parallel()
		v := []int64{10, 20, 30}
		ret, clientErr := c.Int64Slice(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("SimpleSlice", func(t *testing.T) {
		t.Parallel()
		v := []server.Simple{
			{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "one"},
			{Bool: false, Int: 4, Int64: 5, Float64: 6.0, String: "two"},
		}
		ret, clientErr := c.SimpleSlice(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("SimplePtrSlice", func(t *testing.T) {
		t.Parallel()
		s1 := server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "one"}
		s2 := server.Simple{Bool: false, Int: 4, Int64: 5, Float64: 6.0, String: "two"}
		v := []*server.Simple{&s1, &s2}
		ret, clientErr := c.SimplePtrSlice(t.Context(), v)
		require.NoError(t, clientErr)
		require.Len(t, ret, 2)
		require.NotNil(t, ret[0])
		require.NotNil(t, ret[1])
		assert.Equal(t, s1, *ret[0])
		assert.Equal(t, s2, *ret[1])
	})

	t.Run("StringSlice2D", func(t *testing.T) {
		t.Parallel()
		v := [][]string{{"a", "b"}, {"c", "d"}}
		ret, clientErr := c.StringSlice2D(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringStringMap", func(t *testing.T) {
		t.Parallel()
		v := map[string]string{"a": "1", "b": "2"}
		ret, clientErr := c.StringStringMap(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringInt64Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]int64{"x": 10, "y": 20}
		ret, clientErr := c.StringInt64Map(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringSimpleMap", func(t *testing.T) {
		t.Parallel()
		v := map[string]server.Simple{
			"one": {Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "one"},
		}
		ret, clientErr := c.StringSimpleMap(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringSimplePtrMap", func(t *testing.T) {
		t.Parallel()
		s := server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "val"}
		v := map[string]*server.Simple{"k": &s}
		ret, clientErr := c.StringSimplePtrMap(t.Context(), v)
		require.NoError(t, clientErr)
		require.Contains(t, ret, "k")
		require.NotNil(t, ret["k"])
		assert.Equal(t, s, *ret["k"])
	})

	t.Run("StringStringSliceMap", func(t *testing.T) {
		t.Parallel()
		v := map[string][]string{"colors": {"red", "blue"}, "sizes": {"s", "m"}}
		ret, clientErr := c.StringStringSliceMap(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("MapOfMaps", func(t *testing.T) {
		t.Parallel()
		v := map[string]map[string]string{"outer": {"inner": "val"}}
		ret, clientErr := c.MapOfMaps(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("MapOfSimpleSlice", func(t *testing.T) {
		t.Parallel()
		v := map[string][]server.Simple{
			"group": {{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "item"}},
		}
		ret, clientErr := c.MapOfSimpleSlice(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("SliceOfMaps", func(t *testing.T) {
		t.Parallel()
		v := []map[string]string{{"a": "1"}, {"b": "2"}}
		ret, clientErr := c.SliceOfMaps(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("MultiArgs", func(t *testing.T) {
		t.Parallel()
		retA, retB, retC, clientErr := c.MultiArgs(t.Context(), "hello", int64(42), true)
		require.NoError(t, clientErr)
		assert.Equal(t, "hello", retA)
		assert.Equal(t, int64(42), retB)
		assert.True(t, retC)
	})

	t.Run("MixedArgs", func(t *testing.T) {
		t.Parallel()
		s := server.Simple{Bool: true, Int: 1, Int64: 2, Float64: 3.0, String: "mix"}
		items := []string{"a", "b"}
		m := map[string]int64{"x": 10}
		retS, retItems, retM, clientErr := c.MixedArgs(t.Context(), s, items, m)
		require.NoError(t, clientErr)
		assert.Equal(t, s, retS)
		assert.Equal(t, items, retItems)
		assert.Equal(t, m, retM)
	})

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Empty(t.Context())
		require.NoError(t, clientErr)
		assert.True(t, ret)
	})

	t.Run("ByteSlice", func(t *testing.T) {
		t.Parallel()
		v := []byte("hello world")
		ret, clientErr := c.ByteSlice(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})
}
