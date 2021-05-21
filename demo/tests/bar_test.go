package tests

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foomo/gotsrpc/v2/demo"
)

func TestHello(t *testing.T) {
	proxy := demo.NewDefaultBarGoTSRPCProxy(&Bar{})
	server := httptest.NewServer(proxy)
	defer server.Close()
	ctx := context.Background()

	client := demo.NewDefaultBarGoTSRPCClient(server.URL)

	if res, clientErr := client.Hello(ctx, 10); assert.NoError(t, clientErr) {
		assert.Equal(t, 10, res)
	}
}

func TestCustomError(t *testing.T) {
	proxy := demo.NewDefaultBarGoTSRPCProxy(&Bar{})
	server := httptest.NewServer(proxy)
	defer server.Close()
	ctx := context.Background()

	client := demo.NewDefaultBarGoTSRPCClient(server.URL)

	if three, four, clientErr := client.CustomError(ctx, "", nil); assert.NoError(t, clientErr) {
		assert.Equal(t, demo.CustomError(""), three)
		assert.Nil(t, four)
	}

	one := demo.CustomErrorDemo
	two := demo.ErrCustomDemo
	if three, four, clientErr := client.CustomError(ctx, one, two); assert.NoError(t, clientErr) {
		assert.NotNil(t, three)
		assert.NotNil(t, four)
		assert.ErrorIs(t, demo.ErrCustomDemo, two)
	}
}
