package main_test

import (
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/context/server"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	cmd := exec.CommandContext(t.Context(), "bun", "test", "./client/client.test.ts")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "GOTSRPC_SERVER_URL="+s.URL)
	require.NoError(t, cmd.Run())
}
