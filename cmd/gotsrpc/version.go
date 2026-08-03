package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build-time metadata, injected via -ldflags (see Makefile).
var (
	version        = "v3-dev"
	commitHash     = "n/a"
	buildTimestamp = "n/a"
)

var labelStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()

		if !viper.GetBool("debug") {
			_, _ = fmt.Fprintln(out, version)

			return nil
		}

		buildTime := buildTimestamp
		if value, err := strconv.ParseInt(buildTimestamp, 10, 64); err == nil {
			buildTime = time.Unix(value, 0).String()
		}

		_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("Version:"), version)
		_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("Commit:"), commitHash)
		_, _ = fmt.Fprintf(out, "%s %s\n", labelStyle.Render("BuildTime:"), buildTime)

		return nil
	},
}
