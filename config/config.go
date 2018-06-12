package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PHPTarget struct {
	Out       string `yaml:"out"`
	Namespace string `yaml:"namespace"`
}

type Target struct {
	Package          string                `yaml:"package"`
	Services         map[string]string     `yaml:"services"`
	TypeScriptModule string                `yaml:"module"`
	Out              string                `yaml:"out"`
	GoRPC            []string              `yaml:"gorpc"`
	TSRPC            []string              `yaml:"tsrpc"`
	PHPRPC           map[string]*PHPTarget `yaml:"phprpc"`
}

func (t *Target) IsGoRPC(service string) bool {
	for _, value := range t.GoRPC {
		if value == service {
			return true
		}
	}
	return false
}

func (t *Target) IsTSRPC(service string) bool {
	if len(t.TSRPC) == 0 {
		return true
	}
	for _, value := range t.TSRPC {
		if value == service {
			return true
		}
	}
	return false
}

func (t *Target) IsPHPRPC(service string) bool {
	if len(t.PHPRPC) == 0 {
		return false
	}
	_, ok := t.PHPRPC[service]
	return ok
}

func (t *Target) GetPHPTarget(service string) *PHPTarget {
	return t.PHPRPC[service]
}

type Mapping struct {
	GoPackage        string `yaml:"-"`
	Out              string `yaml:"out"`
	TypeScriptModule string `yaml:"module"`
}

type TypeScriptMappings map[string]*Mapping

type ModuleKind string
type TSClientFlavor string

const (
	ModuleKindDefault   ModuleKind     = "default"
	ModuleKindCommonJS  ModuleKind     = "commonjs"
	TSClientFlavorAsync TSClientFlavor = "async"
)

type Config struct {
	ModuleKind     ModuleKind
	TSClientFlavor TSClientFlavor
	Targets        map[string]*Target
	Mappings       TypeScriptMappings
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
	switch conf.TSClientFlavor {
	case "", TSClientFlavorAsync:
	default:
		err = errors.New("unknown ts client flavor: " + conf.TSClientFlavor + " must be empty or " + TSClientFlavorAsync)
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
