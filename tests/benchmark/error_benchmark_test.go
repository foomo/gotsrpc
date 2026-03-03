package benchmark_test

import (
	"errors"
	"testing"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	server "github.com/foomo/gotsrpc/v2/tests/errors/server"
	pkgerrors "github.com/pkg/errors"
)

func BenchmarkNewError_Simple(b *testing.B) {
	err := errors.New("simple error")
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = gotsrpc.NewError(err)
	}
}

func BenchmarkNewError_Wrapped(b *testing.B) {
	base := errors.New("base error")
	wrapped := pkgerrors.Wrap(base, "wrapped")
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = gotsrpc.NewError(wrapped)
	}
}

func BenchmarkNewError_Struct(b *testing.B) {
	err := server.NewMyStructError("struct error")
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = gotsrpc.NewError(err)
	}
}

func BenchmarkNewError_Custom(b *testing.B) {
	err := server.NewMyCustomError("custom error")
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = gotsrpc.NewError(err)
	}
}

func BenchmarkError_Is(b *testing.B) {
	e := gotsrpc.NewError(server.ErrTyped)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = e.Is(server.ErrTyped)
	}
}

func BenchmarkError_As(b *testing.B) {
	e := gotsrpc.NewError(server.NewMyStructError("struct error"))
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		var target server.MyStructError
		_ = e.As(&target)
	}
}
