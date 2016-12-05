package config

import (
	"errors"
	"fmt"
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

type ModuleKind string

const (
	ModuleKindDefault  ModuleKind = "default"
	ModuleKindCommonJS ModuleKind = "commonjs"
)

type Config struct {
	ModuleKind ModuleKind
	Targets    map[string]*Target
	Mappings   TypeScriptMappings
}

func LoadConfigFile(file string) (conf *Config, err error) {
	yamlBytes, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		err = errors.New("could not read config file: " + readErr.Error())
		return
	}
	return loadConfig(yamlBytes)
}

func loadConfig(yamlBytes []byte) (conf *Config, err error) {
	conf = &Config{}
	yamlErr := yaml.Unmarshal(yamlBytes, conf)
	if yamlErr != nil {
		err = errors.New("could not parse yaml: " + yamlErr.Error())
		return
	}
	switch conf.ModuleKind {
	case ModuleKindCommonJS, ModuleKindDefault:
	case "":
		conf.ModuleKind = ModuleKindDefault

	default:
		err = errors.New(fmt.Sprintln("illegal module kind:", conf.ModuleKind, "must be in", ModuleKindDefault, ModuleKindCommonJS))
		return
	}
	for goPackage, mapping := range conf.Mappings {
		mapping.GoPackage = goPackage
	}
	return
}
