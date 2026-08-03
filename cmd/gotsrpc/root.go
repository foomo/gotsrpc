package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logger is initialised in the root command's PersistentPreRunE and shared by
// all subcommands.
var logger *log.Logger

var rootCmd = &cobra.Command{
	Use:           "gotsrpc [config-file]",
	Short:         "Generate type-safe RPC bindings between Go services and TypeScript clients",
	Version:       version,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		level := log.InfoLevel

		if v := viper.GetString("log-level"); v != "" {
			parsed, err := log.ParseLevel(v)
			if err != nil {
				return fmt.Errorf("invalid log level %q: %w", v, err)
			}

			level = parsed
		}

		if viper.GetBool("debug") {
			level = log.DebugLevel
		}

		logger = newLogger(level)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug output")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")

	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))

	viper.SetEnvPrefix("GOTSRPC")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	rootCmd.SetVersionTemplate("{{.Version}}\n")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute runs the root command. For backward compatibility a bare invocation
// with a config file (e.g. `gotsrpc gotsrpc.yml`) is routed to the `generate`
// subcommand.
func Execute() error {
	args := os.Args[1:]
	if shouldDefaultToGenerate(args) {
		args = append([]string{"generate"}, args...)
	}

	rootCmd.SetArgs(args)

	return rootCmd.Execute()
}

// valueFlags lists persistent flags that consume the following argument, so the
// default-subcommand detection can skip their values.
var valueFlags = map[string]bool{
	"--log-level": true,
}

// knownCommands are the top-level command tokens that should not be rewritten to
// the default `generate` subcommand.
var knownCommands = map[string]bool{
	"generate":   true,
	"gen":        true,
	"version":    true,
	"help":       true,
	"completion": true,
}

func shouldDefaultToGenerate(args []string) bool {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			if valueFlags[arg] {
				i++ // skip this flag's value (the `--flag=value` form has no separate token)
			}

			continue
		}

		// first positional token: default to generate unless it is a known command
		return !knownCommands[arg]
	}

	return false
}
