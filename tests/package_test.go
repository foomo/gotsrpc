package tests_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypecheck(t *testing.T) {
	install := exec.CommandContext(t.Context(), "bun", "install")
	install.Stdout = t.Output()
	install.Stderr = t.Output()
	require.NoError(t, install.Run())

	cmd := exec.CommandContext(t.Context(), "bun", "run", "-b", "typecheck")
	cmd.Stdout = t.Output()
	cmd.Stderr = t.Output()
	require.NoError(t, cmd.Run())
}
