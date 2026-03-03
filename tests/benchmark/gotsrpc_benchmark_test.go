package benchmark_test

import (
	"context"
	"net/http/httptest"
	"testing"

	server "github.com/foomo/gotsrpc/v2/tests/types/server"
)

func setupGoTSRPC(b *testing.B) *server.HTTPServiceGoTSRPCClient {
	b.Helper()
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	b.Cleanup(s.Close)
	return server.NewDefaultServiceGoTSRPCClient(s.URL)
}

func BenchmarkGoTSRPC_Empty(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.Empty(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_String(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.String(ctx, "hello world"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_SimpleStruct(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := server.Simple{
		Bool:    true,
		Int:     42,
		Int64:   123456789,
		Float64: 3.14159,
		String:  "benchmark",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.SimpleStruct(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_NestedStruct(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := server.Nested{
		Name: "parent",
		Child: server.Simple{
			Bool:    true,
			Int:     42,
			Int64:   123456789,
			Float64: 3.14159,
			String:  "child",
		},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.NestedStruct(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_StructWithCollections(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := server.WithCollections{
		Strings: []string{"a", "b", "c"},
		Int64s:  []int64{1, 2, 3},
		Items: []server.Simple{
			{Bool: true, Int: 1, Int64: 1, Float64: 1.0, String: "one"},
		},
		ItemPtrs: []*server.Simple{
			{Bool: false, Int: 2, Int64: 2, Float64: 2.0, String: "two"},
		},
		StringMap: map[string]string{"key": "value"},
		StructMap: map[string]server.Simple{
			"item": {Bool: true, Int: 3, Int64: 3, Float64: 3.0, String: "three"},
		},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.StructWithCollections(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_MultiArgs(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, _, _, err := c.MultiArgs(ctx, "hello", 42, true); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_MixedArgs(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	s := server.Simple{Bool: true, Int: 1, Int64: 1, Float64: 1.0, String: "one"}
	items := []string{"a", "b", "c"}
	m := map[string]int64{"x": 1, "y": 2}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, _, _, err := c.MixedArgs(ctx, s, items, m); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_StringSlice(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.StringSlice(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_StringStringMap(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.StringStringMap(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_MapOfMaps(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := map[string]map[string]string{
		"outer1": {"a": "1", "b": "2"},
		"outer2": {"c": "3", "d": "4"},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.MapOfMaps(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoTSRPC_ByteSlice(b *testing.B) {
	c := setupGoTSRPC(b)
	ctx := context.Background()
	v := make([]byte, 1024)
	for i := range v {
		v[i] = byte(i % 256)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.ByteSlice(ctx, v); err != nil {
			b.Fatal(err)
		}
	}
}
