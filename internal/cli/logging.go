package cli

import (
	"os"

	"github.com/charmbracelet/log"
)

// newLogger constructs the CLI logger writing styled output to stderr.
func newLogger(level log.Level) *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		Level:           level,
		ReportTimestamp: false,
	})
}
