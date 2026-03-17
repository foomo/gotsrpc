package tests_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypecheck(t *testing.T) {
	cmd := exec.CommandContext(t.Context(), "bun", "run", "typecheck")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Run())
}
