package parser

import (
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v2"
)

func traceData(args ...any) {
	for _, arg := range args {
		yamlBytes, err := yaml.Marshal(arg)
		if err != nil {
			log.Debug(arg)
			continue
		}

		log.Debug(string(yamlBytes))
	}
}
