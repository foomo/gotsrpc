package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var labelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Long: "Display version information.\n\n" +
			"By default only the version string is printed. With --debug (or\n" +
			"GOTSRPC_DEBUG=true) the commit hash and build time are printed as well.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			if !viper.GetBool("debug") {
				_, _ = fmt.Fprintln(out, buildInfo.Version)

				return nil
			}

			buildTime := buildInfo.BuildTime
			if value, err := strconv.ParseInt(buildInfo.BuildTime, 10, 64); err == nil {
				buildTime = time.Unix(value, 0).String()
			}

			_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("Version:"), buildInfo.Version)
			_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("Commit:"), buildInfo.Commit)
			_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("BuildTime:"), buildTime)

			return nil
		},
	}
}
