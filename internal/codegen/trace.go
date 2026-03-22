package codegen

import (
	"fmt"
	"os"
)

// Trace enables debug output for code generation.
var Trace = false

func trace(args ...interface{}) {
	if Trace {
		_, _ = fmt.Fprintln(os.Stderr, args...)
	}
}
