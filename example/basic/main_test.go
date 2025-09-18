package main_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v2/example/basic/service"
	"github.com/stretchr/testify/assert"
)

func TestContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	svr := httptest.NewServer(service.NewDefaultServiceGoTSRPCProxy(&service.Handler{}))
	defer svr.Close()

	client := service.NewDefaultServiceGoTSRPCClient(svr.URL)

	clientErr := client.Context(ctx)
	assert.ErrorIs(t, clientErr, context.Canceled)
}
