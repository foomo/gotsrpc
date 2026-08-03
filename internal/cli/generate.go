package cli

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/foomo/gotsrpc/v3/config"
	"github.com/foomo/gotsrpc/v3/internal/build"
)

func newGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "generate [config-file]",
		Aliases: []string{"gen"},
		Short:   "Generate RPC bindings from a gotsrpc config file",
		Long: "Generate RPC bindings from a gotsrpc config file.\n\n" +
			"If no config file is given, gotsrpc.yml in the current directory is used.\n\n" +
			"The referenced Go code has to compile and mappings for all used packages must\n" +
			"be configured. Previously generated files are overwritten; obsolete files are\n" +
			"not removed, so add a clean step to your build if needed.",
		Example: "  # Use gotsrpc.yml in the current directory\n" +
			"  gotsrpc generate\n\n" +
			"  # Use an explicit config file\n" +
			"  gotsrpc generate path/to/gotsrpc.yml\n\n" +
			"  # Increase log verbosity\n" +
			"  gotsrpc generate --log-level debug",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := "gotsrpc.yml"
			if len(args) == 1 {
				configFile = args[0]
			}

			ctx := cmd.Context()

			goRoot, err := goEnv(ctx, "GOROOT")
			if err != nil {
				return err
			}

			goPath, err := goEnv(ctx, "GOPATH")
			if err != nil {
				return err
			}

			conf, err := config.LoadConfigFile(configFile)
			if err != nil {
				return fmt.Errorf("could not load config from %s: %w", configFile, err)
			}

			return build.Build(logger, conf, goPath, goRoot)
		},
	}
}

func goEnv(ctx context.Context, name string) (string, error) {
	out, err := exec.CommandContext(ctx, "go", "env", name).Output()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve %s: %w", name, err)
	}

	return string(bytes.TrimSpace(out)), nil
}
