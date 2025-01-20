package service

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	service := NewDefaultServiceGoTSRPCProxy(&Handler{})

	server := httptest.NewServer(service)
	client := NewDefaultServiceGoTSRPCClient(server.URL,
		gotsrpc.WithHTTPClient(server.Client()),
		gotsrpc.WithSnappyCompression())
	response, err := client.String(context.Background(), "test")
	require.NoError(t, err)
	require.Equal(t, "test", response)
	fmt.Println(response)
}
