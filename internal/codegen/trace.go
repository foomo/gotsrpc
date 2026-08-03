package codegen

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/log"
)

// Logger receives debug output for code generation. It defaults to a no-op
// logger and can be replaced via SetLogger.
var Logger = log.New(io.Discard)

// SetLogger sets the logger used for trace output.
func SetLogger(l *log.Logger) {
	if l != nil {
		Logger = l
	}
}

func trace(args ...any) {
	if Logger.GetLevel() > log.DebugLevel {
		return
	}

	Logger.Debug(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}
