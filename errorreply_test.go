package gotsrpc //nolint:testpackage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorReply(t *testing.T) {
	t.Parallel()

	inner := errors.New("boom")
	reply := ErrorReply(inner)

	wrapper, ok := reply.(*errorReply)
	require.True(t, ok)
	assert.Same(t, inner, wrapper.err)
}
