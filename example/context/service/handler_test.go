package service_test

import (
	"context"
	"testing"

	"github.com/foomo/gotsrpc/v2/example/context/service"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	h := &service.Handler{}
	msg := "foomo"

	assert.Equal(t, "Hello foomo", h.Hello(context.Background(), msg))

	assert.Error(t, h.TypedError(context.Background(), msg))
	assert.Error(t, h.CustomError(context.Background(), msg))
	assert.Error(t, h.WrappedError(context.Background(), msg))
	assert.Error(t, h.StandardError(context.Background(), msg))
}
