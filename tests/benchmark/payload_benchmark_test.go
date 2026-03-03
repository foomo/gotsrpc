package benchmark_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	server "github.com/foomo/gotsrpc/v2/tests/types/server"
)

func BenchmarkGoTSRPC_StringSlice_Scale(b *testing.B) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	b.Cleanup(s.Close)
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	ctx := context.Background()

	for _, n := range []int{10, 100, 1000} {
		v := make([]string, n)
		for i := range v {
			v[i] = fmt.Sprintf("item-%d", i)
		}
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for b.Loop() {
				if _, err := c.StringSlice(ctx, v); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkGoTSRPC_SimpleSlice_Scale(b *testing.B) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	b.Cleanup(s.Close)
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	ctx := context.Background()

	for _, n := range []int{10, 100, 1000} {
		v := make([]server.Simple, n)
		for i := range v {
			v[i] = server.Simple{
				Bool:    i%2 == 0,
				Int:     i,
				Int64:   int64(i),
				Float64: float64(i),
				String:  fmt.Sprintf("item-%d", i),
			}
		}
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for b.Loop() {
				if _, err := c.SimpleSlice(ctx, v); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkGoTSRPC_StringStringMap_Scale(b *testing.B) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	b.Cleanup(s.Close)
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	ctx := context.Background()

	for _, n := range []int{10, 100, 1000} {
		v := make(map[string]string, n)
		for i := range n {
			v[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
		}
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for b.Loop() {
				if _, err := c.StringStringMap(ctx, v); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
