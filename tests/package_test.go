package tests_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypecheck(t *testing.T) {
	t.Skip()
	install := exec.CommandContext(t.Context(), "bun", "install")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	require.NoError(t, install.Run())

	cmd := exec.CommandContext(t.Context(), "bun", "run", "-b", "typecheck")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Run())
}
