package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v2"
)

// Logger receives trace output. It defaults to a no-op logger and can be
// replaced via SetLogger.
var Logger = log.New(io.Discard)

// SetLogger sets the logger used for trace output.
func SetLogger(l *log.Logger) {
	if l != nil {
		Logger = l
	}
}

// traceEnabled reports whether debug-level trace output would be emitted.
func traceEnabled() bool {
	return Logger.GetLevel() <= log.DebugLevel
}

func trace(args ...any) {
	if !traceEnabled() {
		return
	}

	Logger.Debug(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}

func traceData(args ...any) {
	if !traceEnabled() {
		return
	}

	for _, arg := range args {
		yamlBytes, errMarshal := yaml.Marshal(arg)
		if errMarshal != nil {
			trace(arg)
			continue
		}

		trace(string(yamlBytes))
	}
}
