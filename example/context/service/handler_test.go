package service_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v2/example/context/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	h := &service.Handler{}
	msg := "foomo"

	assert.Equal(t, "Hello foomo", h.Hello(t.Context(), msg))

	require.Error(t, h.TypedError(t.Context(), msg))
	require.Error(t, h.CustomError(t.Context(), msg))
	require.Error(t, h.WrappedError(t.Context(), msg))
	require.Error(t, h.StandardError(t.Context(), msg))
}
