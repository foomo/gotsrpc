package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Target struct {
	Package          string
	Services         map[string]string
	TypeScriptModule string `yaml:"module"`
	Out              string
}

type Mapping struct {
	GoPackage        string `yaml:"-"`
	Out              string `yaml:"out"`
	TypeScriptModule string `yaml:"module"`
}

type TypeScriptMappings map[string]*Mapping

type Config struct {
	Targets  map[string]*Target
	Mappings TypeScriptMappings
}

func LoadConfigFile(file string) (conf *Config, err error) {
	yamlBytes, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		err = readErr
		return
	}
	return loadConfig(yamlBytes)
}

func loadConfig(yamlBytes []byte) (conf *Config, err error) {
	conf = &Config{}
	yamlErr := yaml.Unmarshal(yamlBytes, conf)
	if yamlErr != nil {
		err = yamlErr
		return
	}
	for goPackage, mapping := range conf.Mappings {
		mapping.GoPackage = goPackage
	}
	return
}
