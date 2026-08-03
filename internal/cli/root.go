package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BuildInfo carries build-time metadata (injected via -ldflags in the main
// package) into the command tree.
type BuildInfo struct {
	Version   string
	Commit    string
	BuildTime string
}

// logger is initialised in the root command's PersistentPreRunE and shared by
// all subcommands.
var logger *log.Logger

// buildInfo holds the metadata passed to the most recently constructed root
// command, consumed by the version command.
var buildInfo BuildInfo

// Logger returns the shared CLI logger. It is nil until the root command's
// PersistentPreRunE has run.
func Logger() *log.Logger {
	return logger
}

// NewRootCmd builds the gotsrpc command tree.
func NewRootCmd(info BuildInfo) *cobra.Command {
	buildInfo = info

	rootCmd := &cobra.Command{
		Use:   "gotsrpc [config-file]",
		Short: "Generate type-safe RPC bindings between Go services and TypeScript clients",
		Long: "gotsrpc generates type-safe RPC bindings between Go services and TypeScript\n" +
			"clients, with optional Go-to-Go RPC support. It parses your Go source via the\n" +
			"AST, reads a YAML config (gotsrpc.yml) and writes Go proxies/clients and\n" +
			"TypeScript clients and type definitions.\n\n" +
			"Running gotsrpc with a config file and no subcommand is shorthand for\n" +
			"`gotsrpc generate <config-file>`.\n\n" +
			"Flags can also be set via environment variables prefixed with GOTSRPC_, e.g.\n" +
			"GOTSRPC_LOG_LEVEL=debug or GOTSRPC_DEBUG=true.",
		Example: "  # Generate from gotsrpc.yml in the current directory\n" +
			"  gotsrpc gotsrpc.yml\n\n" +
			"  # Equivalent, using the explicit subcommand\n" +
			"  gotsrpc generate path/to/gotsrpc.yml\n\n" +
			"  # Print the version\n" +
			"  gotsrpc version",
		Version:           info.Version,
		SilenceUsage:      true,
		SilenceErrors:     true,
		DisableAutoGenTag: true,
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

	rootCmd.PersistentFlags().Bool("debug", false, "enable debug output")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")

	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))

	viper.SetEnvPrefix("GOTSRPC")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	rootCmd.SetVersionTemplate("{{.Version}}\n")

	rootCmd.AddCommand(newGenerateCmd())
	rootCmd.AddCommand(newVersionCmd())

	return rootCmd
}

// Execute runs the root command. For backward compatibility a bare invocation
// with a config file (e.g. `gotsrpc gotsrpc.yml`) is routed to the `generate`
// subcommand.
func Execute(info BuildInfo) error {
	args := os.Args[1:]
	if shouldDefaultToGenerate(args) {
		args = append([]string{"generate"}, args...)
	}

	rootCmd := NewRootCmd(info)
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
