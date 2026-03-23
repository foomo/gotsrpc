package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var ReaderTrace = false

func trace(args ...any) {
	if ReaderTrace {
		_, _ = fmt.Fprintln(os.Stderr, args...)
	}
}

func traceData(args ...any) {
	if ReaderTrace {
		for _, arg := range args {
			yamlBytes, errMarshal := yaml.Marshal(arg)
			if errMarshal != nil {
				trace(arg)
				continue
			}

			trace(string(yamlBytes))
		}
	}
}
