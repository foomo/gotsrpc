package server_test

import (
	"net"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/types/server"
)

func setupGoRPC(b *testing.B) *server.ServiceGoRPCClient {
	b.Helper()

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

func BenchmarkGoRPC_Empty(b *testing.B) {
	c := setupGoRPC(b)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.Empty(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoRPC_String(b *testing.B) {
	c := setupGoRPC(b)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := c.String("hello world"); err != nil {
			b.Fatal(err)
		}
	}
}
