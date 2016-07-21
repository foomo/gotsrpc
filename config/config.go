package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Target struct {
	Name     string
	Package  string
	Services []string
}

type Mapping struct {
	GoPackage        string `yaml:"-"`
	Out              string `yaml:"out"`
	TypeScriptModule string `yaml:"module"`
}

type Config struct {
	Targets  map[string]*Target
	Mappings map[string]*Mapping
}

func LoadConfigFile(file string) (conf *Config, err error) {
	yamlBytes, readErr := ioutil.ReadFile(file)
	if err != nil {
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
