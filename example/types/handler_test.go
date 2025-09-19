package service_test

import (
	"net/http/httptest"
	"testing"

	service "github.com/foomo/gotsrpc/v2/example/types"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	s := httptest.NewServer(service.NewDefaultServiceGoTSRPCProxy(&service.Handler{}))
	c := service.NewDefaultServiceGoTSRPCClient(s.URL)
	c.Client.SetTransportHttpClient(s.Client())

	require.NoError(t, c.String(t.Context(), "s"))

	require.NoError(t, c.Strings(t.Context(), "a", "b"))
}
