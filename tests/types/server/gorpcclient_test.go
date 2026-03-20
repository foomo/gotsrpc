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

	t.Run("Int8", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int8(127)
		require.NoError(t, clientErr)
		assert.Equal(t, int8(127), ret)
	})

	t.Run("Int16", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int16(32767)
		require.NoError(t, clientErr)
		assert.Equal(t, int16(32767), ret)
	})

	t.Run("Int32", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Int32(2147483647)
		require.NoError(t, clientErr)
		assert.Equal(t, int32(2147483647), ret)
	})

	t.Run("Uint", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Uint(42)
		require.NoError(t, clientErr)
		assert.Equal(t, uint(42), ret)
	})

	t.Run("Uint8", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Uint8(255)
		require.NoError(t, clientErr)
		assert.Equal(t, uint8(255), ret)
	})

	t.Run("Uint16", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Uint16(65535)
		require.NoError(t, clientErr)
		assert.Equal(t, uint16(65535), ret)
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Uint32(4294967295)
		require.NoError(t, clientErr)
		assert.Equal(t, uint32(4294967295), ret)
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Uint64(9876543210)
		require.NoError(t, clientErr)
		assert.Equal(t, uint64(9876543210), ret)
	})

	t.Run("Float32", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.Float32(3.14)
		require.NoError(t, clientErr)
		assert.InDelta(t, float32(3.14), ret, 1e-5)
	})

	t.Run("AllScalarsStruct", func(t *testing.T) {
		t.Parallel()
		v := server.AllScalars{
			Int8: 127, Int16: 32767, Int32: 2147483647,
			Uint: 42, Uint8: 255, Uint16: 65535, Uint32: 4294967295, Uint64: 9876543210,
			Float32: 3.14,
		}
		ret, clientErr := c.AllScalarsStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.Int8, ret.Int8)
		assert.Equal(t, v.Int16, ret.Int16)
		assert.Equal(t, v.Int32, ret.Int32)
		assert.Equal(t, v.Uint, ret.Uint)
		assert.Equal(t, v.Uint8, ret.Uint8)
		assert.Equal(t, v.Uint16, ret.Uint16)
		assert.Equal(t, v.Uint32, ret.Uint32)
		assert.Equal(t, v.Uint64, ret.Uint64)
		assert.InDelta(t, v.Float32, ret.Float32, 1e-5)
	})

	t.Run("StringPtr", func(t *testing.T) {
		t.Parallel()
		v := "test"
		ret, clientErr := c.StringPtr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Int8Ptr", func(t *testing.T) {
		t.Parallel()
		v := int8(127)
		ret, clientErr := c.Int8Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Int16Ptr", func(t *testing.T) {
		t.Parallel()
		v := int16(32767)
		ret, clientErr := c.Int16Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Int32Ptr", func(t *testing.T) {
		t.Parallel()
		v := int32(2147483647)
		ret, clientErr := c.Int32Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("UintPtr", func(t *testing.T) {
		t.Parallel()
		v := uint(42)
		ret, clientErr := c.UintPtr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Uint8Ptr", func(t *testing.T) {
		t.Parallel()
		v := uint8(255)
		ret, clientErr := c.Uint8Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Uint16Ptr", func(t *testing.T) {
		t.Parallel()
		v := uint16(65535)
		ret, clientErr := c.Uint16Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Uint32Ptr", func(t *testing.T) {
		t.Parallel()
		v := uint32(4294967295)
		ret, clientErr := c.Uint32Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Uint64Ptr", func(t *testing.T) {
		t.Parallel()
		v := uint64(9876543210)
		ret, clientErr := c.Uint64Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.Equal(t, v, *ret)
	})

	t.Run("Float32Ptr", func(t *testing.T) {
		t.Parallel()
		v := float32(3.14)
		ret, clientErr := c.Float32Ptr(&v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret)
		assert.InDelta(t, v, *ret, 1e-5)
	})

	t.Run("AllScalarPointersStruct", func(t *testing.T) {
		t.Parallel()
		i8, i16, i32 := int8(127), int16(32767), int32(2147483647)
		u, u8, u16, u32, u64 := uint(42), uint8(255), uint16(65535), uint32(4294967295), uint64(9876543210)
		f32 := float32(3.14)
		v := server.AllScalarPointers{
			Int8Ptr: &i8, Int16Ptr: &i16, Int32Ptr: &i32,
			UintPtr: &u, Uint8Ptr: &u8, Uint16Ptr: &u16, Uint32Ptr: &u32, Uint64Ptr: &u64,
			Float32Ptr: &f32,
		}
		ret, clientErr := c.AllScalarPointersStruct(v)
		require.NoError(t, clientErr)
		require.NotNil(t, ret.Int8Ptr)
		assert.Equal(t, i8, *ret.Int8Ptr)
		require.NotNil(t, ret.Int16Ptr)
		assert.Equal(t, i16, *ret.Int16Ptr)
		require.NotNil(t, ret.Int32Ptr)
		assert.Equal(t, i32, *ret.Int32Ptr)
		require.NotNil(t, ret.UintPtr)
		assert.Equal(t, u, *ret.UintPtr)
		require.NotNil(t, ret.Uint8Ptr)
		assert.Equal(t, u8, *ret.Uint8Ptr)
		require.NotNil(t, ret.Uint16Ptr)
		assert.Equal(t, u16, *ret.Uint16Ptr)
		require.NotNil(t, ret.Uint32Ptr)
		assert.Equal(t, u32, *ret.Uint32Ptr)
		require.NotNil(t, ret.Uint64Ptr)
		assert.Equal(t, u64, *ret.Uint64Ptr)
		require.NotNil(t, ret.Float32Ptr)
		assert.InDelta(t, f32, *ret.Float32Ptr, 1e-5)
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

	t.Run("Int8Slice", func(t *testing.T) {
		t.Parallel()
		v := []int8{-128, 0, 127}
		ret, clientErr := c.Int8Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Int16Slice", func(t *testing.T) {
		t.Parallel()
		v := []int16{-32768, 0, 32767}
		ret, clientErr := c.Int16Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Int32Slice", func(t *testing.T) {
		t.Parallel()
		v := []int32{-2147483648, 0, 2147483647}
		ret, clientErr := c.Int32Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("UintSlice", func(t *testing.T) {
		t.Parallel()
		v := []uint{0, 42, 100}
		ret, clientErr := c.UintSlice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Uint16Slice", func(t *testing.T) {
		t.Parallel()
		v := []uint16{0, 1000, 65535}
		ret, clientErr := c.Uint16Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Uint32Slice", func(t *testing.T) {
		t.Parallel()
		v := []uint32{0, 1000, 4294967295}
		ret, clientErr := c.Uint32Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Uint64Slice", func(t *testing.T) {
		t.Parallel()
		v := []uint64{0, 1000, 9876543210}
		ret, clientErr := c.Uint64Slice(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("Float32Slice", func(t *testing.T) {
		t.Parallel()
		v := []float32{1.1, 2.2, 3.3}
		ret, clientErr := c.Float32Slice(v)
		require.NoError(t, clientErr)
		require.Len(t, ret, 3)
		assert.InDelta(t, v[0], ret[0], 1e-5)
		assert.InDelta(t, v[1], ret[1], 1e-5)
		assert.InDelta(t, v[2], ret[2], 1e-5)
	})

	t.Run("AllScalarSlicesStruct", func(t *testing.T) {
		t.Parallel()
		v := server.AllScalarSlices{
			Int8s: []int8{-1, 0, 1}, Int16s: []int16{-1, 0, 1}, Int32s: []int32{-1, 0, 1},
			Uints: []uint{0, 1, 2}, Uint16s: []uint16{0, 1, 2},
			Uint32s: []uint32{0, 1, 2}, Uint64s: []uint64{0, 1, 2},
			Float32s: []float32{1.1, 2.2},
		}
		ret, clientErr := c.AllScalarSlicesStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.Int8s, ret.Int8s)
		assert.Equal(t, v.Int16s, ret.Int16s)
		assert.Equal(t, v.Int32s, ret.Int32s)
		assert.Equal(t, v.Uints, ret.Uints)
		assert.Equal(t, v.Uint16s, ret.Uint16s)
		assert.Equal(t, v.Uint32s, ret.Uint32s)
		assert.Equal(t, v.Uint64s, ret.Uint64s)
		require.Len(t, ret.Float32s, 2)
		assert.InDelta(t, v.Float32s[0], ret.Float32s[0], 1e-5)
		assert.InDelta(t, v.Float32s[1], ret.Float32s[1], 1e-5)
	})

	t.Run("StringStringMap", func(t *testing.T) {
		t.Parallel()
		v := map[string]string{"a": "1", "b": "2"}
		ret, clientErr := c.StringStringMap(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringInt8Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]int8{"a": 1, "b": -1}
		ret, clientErr := c.StringInt8Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringInt16Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]int16{"a": 100, "b": -100}
		ret, clientErr := c.StringInt16Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringInt32Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]int32{"a": 100000, "b": -100000}
		ret, clientErr := c.StringInt32Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringUintMap", func(t *testing.T) {
		t.Parallel()
		v := map[string]uint{"a": 0, "b": 42}
		ret, clientErr := c.StringUintMap(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringUint8Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]uint8{"a": 0, "b": 255}
		ret, clientErr := c.StringUint8Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringUint16Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]uint16{"a": 0, "b": 65535}
		ret, clientErr := c.StringUint16Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringUint32Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]uint32{"a": 0, "b": 4294967295}
		ret, clientErr := c.StringUint32Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringUint64Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]uint64{"a": 0, "b": 9876543210}
		ret, clientErr := c.StringUint64Map(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("StringFloat32Map", func(t *testing.T) {
		t.Parallel()
		v := map[string]float32{"a": 1.1, "b": 2.2}
		ret, clientErr := c.StringFloat32Map(v)
		require.NoError(t, clientErr)
		require.Contains(t, ret, "a")
		require.Contains(t, ret, "b")
		assert.InDelta(t, v["a"], ret["a"], 1e-5)
		assert.InDelta(t, v["b"], ret["b"], 1e-5)
	})

	t.Run("AllScalarMapsStruct", func(t *testing.T) {
		t.Parallel()
		v := server.AllScalarMaps{
			Int8Map: map[string]int8{"x": 1}, Int16Map: map[string]int16{"x": 1},
			Int32Map: map[string]int32{"x": 1}, UintMap: map[string]uint{"x": 1},
			Uint8Map: map[string]uint8{"x": 1}, Uint16Map: map[string]uint16{"x": 1},
			Uint32Map: map[string]uint32{"x": 1}, Uint64Map: map[string]uint64{"x": 1},
			Float32Map: map[string]float32{"x": 1.5},
		}
		ret, clientErr := c.AllScalarMapsStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.Int8Map, ret.Int8Map)
		assert.Equal(t, v.Int16Map, ret.Int16Map)
		assert.Equal(t, v.Int32Map, ret.Int32Map)
		assert.Equal(t, v.UintMap, ret.UintMap)
		assert.Equal(t, v.Uint8Map, ret.Uint8Map)
		assert.Equal(t, v.Uint16Map, ret.Uint16Map)
		assert.Equal(t, v.Uint32Map, ret.Uint32Map)
		assert.Equal(t, v.Uint64Map, ret.Uint64Map)
		require.Contains(t, ret.Float32Map, "x")
		assert.InDelta(t, v.Float32Map["x"], ret.Float32Map["x"], 1e-5)
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

	t.Run("StringObjectID", func(t *testing.T) {
		t.Parallel()
		v := server.StringObjectID{}
		copy(v.ObjectID[:], "hello123456")
		ret, clientErr := c.StringObjectID(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.ObjectID, ret.ObjectID)
	})
}
