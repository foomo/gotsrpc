package gotsrpc_test

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v2"
	"github.com/foomo/gotsrpc/v2/tests/common"
	"github.com/foomo/gotsrpc/v2/tests/types/server"
)

func setupGoRPC(b *testing.B) *server.ServiceGoRPCClient {
	b.Helper()

	gotsrpc.SetDefaultHttpClientFactory(func() *http.Client {
		return &http.Client{Timeout: 5 * time.Second} // or reusable client
	})

	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	if err != nil {
		b.Fatal(err)
	}
	addr := l.Addr().String()
	if err := l.Close(); err != nil {
		b.Fatal(err)
	}

	s := server.NewServiceGoRPCProxy(addr, &server.Handler{}, nil)
	if err := s.Start(); err != nil {
		b.Fatal(err)
	}
	b.Cleanup(s.Stop)

	c := server.NewServiceGoRPCClient(addr, nil)
	c.Start()
	b.Cleanup(c.Stop)

	return c
}

func setupGoTSRPC(b *testing.B) *server.HTTPServiceGoTSRPCClient {
	b.Helper()
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	b.Cleanup(s.Close)
	return server.NewDefaultServiceGoTSRPCClient(s.URL)
}

func Benchmark_Empty(b *testing.B) {
	b.ReportAllocs()

	b.Run("GoRPC", func(b *testing.B) {
		c := setupGoRPC(b)
		for b.Loop() {
			if _, err := c.Empty(); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoTSRPC", func(b *testing.B) {
		c := setupGoTSRPC(b)
		for b.Loop() {
			if _, err := c.Empty(b.Context()); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_String(b *testing.B) {
	b.ReportAllocs()
	v := "hello world"

	b.Run("GoRPC", func(b *testing.B) {
		c := setupGoRPC(b)
		for b.Loop() {
			if _, err := c.String(v); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoTSRPC", func(b *testing.B) {
		c := setupGoTSRPC(b)
		for b.Loop() {
			if _, err := c.String(b.Context(), v); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_SimpleStruct(b *testing.B) {
	b.ReportAllocs()
	v := common.Simple{
		Bool:    true,
		Int:     42,
		Int64:   123456789,
		Float64: 3.14159,
		String:  "benchmark",
	}

	b.Run("GoRPC", func(b *testing.B) {
		c := setupGoRPC(b)
		for b.Loop() {
			if _, err := c.SimpleStruct(v); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoTSRPC", func(b *testing.B) {
		c := setupGoTSRPC(b)
		for b.Loop() {
			if _, err := c.SimpleStruct(b.Context(), v); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_NestedStruct(b *testing.B) {
	b.ReportAllocs()
	v := common.Nested{
		Name: "parent",
		Child: common.Simple{
			Bool:    true,
			Int:     42,
			Int64:   123456789,
			Float64: 3.14159,
			String:  "child",
		},
	}

	b.Run("GoRPC", func(b *testing.B) {
		c := setupGoRPC(b)
		for b.Loop() {
			if _, err := c.NestedStruct(v); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoTSRPC", func(b *testing.B) {
		c := setupGoTSRPC(b)
		for b.Loop() {
			if _, err := c.NestedStruct(b.Context(), v); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_StructWithCollections(b *testing.B) {
	b.ReportAllocs()
	v := server.WithCollections{
		Strings: []string{"a", "b", "c"},
		Int64s:  []int64{1, 2, 3},
		Items: []common.Simple{
			{Bool: true, Int: 1, Int64: 1, Float64: 1.0, String: "one"},
		},
		ItemPtrs: []*common.Simple{
			{Bool: false, Int: 2, Int64: 2, Float64: 2.0, String: "two"},
		},
		StringMap: map[string]string{"key": "value"},
		StructMap: map[string]common.Simple{
			"item": {Bool: true, Int: 3, Int64: 3, Float64: 3.0, String: "three"},
		},
	}

	b.Run("GoRPC", func(b *testing.B) {
		c := setupGoRPC(b)
		for b.Loop() {
			if _, err := c.StructWithCollections(v); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoTSRPC", func(b *testing.B) {
		c := setupGoTSRPC(b)
		for b.Loop() {
			if _, err := c.StructWithCollections(b.Context(), v); err != nil {
				b.Fatal(err)
			}
		}
	})
}
