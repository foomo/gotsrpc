package gotsrpc //nolint:testpackage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientError(t *testing.T) {
	t.Parallel()

	t.Run("wraps and unwraps the underlying error", func(t *testing.T) {
		t.Parallel()

		inner := errors.New("boom")
		ce := NewClientError(inner)
		require.NotNil(t, ce)
		assert.Equal(t, "boom", ce.Error())
		assert.Same(t, inner, ce.Unwrap())
		assert.ErrorIs(t, ce, inner)
	})

	t.Run("unwrap returns nil when no inner error", func(t *testing.T) {
		t.Parallel()

		ce := &ClientError{}
		assert.NoError(t, ce.Unwrap())
	})
}
