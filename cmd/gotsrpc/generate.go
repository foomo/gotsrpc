package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/foomo/gotsrpc/v3/config"
	"github.com/foomo/gotsrpc/v3/internal/build"
)

var generateCmd = &cobra.Command{
	Use:     "generate [config-file]",
	Aliases: []string{"gen"},
	Short:   "Generate RPC bindings from a gotsrpc config file",
	Long: "Generate RPC bindings from a gotsrpc config file.\n\n" +
		"If no config file is given, gotsrpc.yml in the current directory is used.",
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

func goEnv(ctx context.Context, name string) (string, error) {
	out, err := exec.CommandContext(ctx, "go", "env", name).Output()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve %s: %w", name, err)
	}

	return string(bytes.TrimSpace(out)), nil
}
