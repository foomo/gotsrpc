package tests_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypecheck(t *testing.T) {
	cmd := exec.CommandContext(t.Context(), "bun", "run", "typecheck")
	cmd.Stdout = t.Output()
	cmd.Stderr = t.Output()
	require.NoError(t, cmd.Run())
}
