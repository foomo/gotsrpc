package gotsrpc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type GoTypeScriptMapping struct {
	GoPackage        string `yaml:"-"`
	TypeScriptDir    string `yaml:"dir"`
	TypeScriptModule string `yaml:"module"`
}

type Config struct {
	Mappings map[string]*GoTypeScriptMapping
}

func loadConfigfile(file string) (conf *Config, err error) {
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
