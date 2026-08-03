package gotsrpc_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v3"
	"github.com/stretchr/testify/assert"
)

type sampleUnion struct {
	A *string
	B *int
}

// Note: RegisterUnionExt / MustRegisterUnionExt are not unit-tested here because
// they mutate the process-global JSON codec handle, which the codec library
// forbids ("cannot modify initialized Handle") once any other test has encoded
// through it. Registration is an init-time operation exercised by the e2e suite.
func TestUnionExt_ConvertExt(t *testing.T) {
	t.Parallel()

	ext := &gotsrpc.UnionExt{}

	t.Run("returns first non-zero field", func(t *testing.T) {
		t.Parallel()

		n := 5
		got := ext.ConvertExt(sampleUnion{B: &n})
		assert.Equal(t, &n, got)
	})

	t.Run("dereferences pointer receiver", func(t *testing.T) {
		t.Parallel()

		s := "hello"
		got := ext.ConvertExt(&sampleUnion{A: &s})
		assert.Equal(t, &s, got)
	})

	t.Run("returns nil when all fields are zero", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, ext.ConvertExt(sampleUnion{}))
	})
}
