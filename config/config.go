package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v2"
)

type Target struct {
	Package          string            `yaml:"package"`
	Services         map[string]string `yaml:"services"`
	TypeScriptModule string            `yaml:"module"`
	Out              string            `yaml:"out"`
	GoRPC            []string          `yaml:"gorpc"`
	TSRPC            []string          `yaml:"tsrpc"`
	SkipTSRPCClient  bool              `yaml:"skipTSRPCClient"`
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

type Mapping struct {
	GoPackage        string   `yaml:"-"`
	Out              string   `yaml:"out"`
	Structs          []string `yaml:"structs"`
	Scalars          []string `yaml:"scalars"`
	TypeScriptModule string   `yaml:"module"`
}

type TypeScriptMappings map[string]*Mapping

type Namespace struct {
	Name    string        `yaml:"name"`
	Path    string        `yaml:"path"`
	ModFile *modfile.File `yaml:"-"`
}

type Config struct {
	Module   Namespace
	Targets  map[string]*Target
	Mappings TypeScriptMappings
}

func LoadConfigFile(file string) (conf *Config, err error) {
	yamlBytes, readErr := os.ReadFile(file)
	if readErr != nil {
		return nil, errors.New("could not read config file: " + readErr.Error())
	}
	conf, err = loadConfig(yamlBytes)
	if err != nil {
		return nil, err
	}

	if conf.Module.Path != "" && !path.IsAbs(conf.Module.Path) {
		absPath, err := filepath.Abs(filepath.Join(filepath.Dir(file), conf.Module.Path))
		if err != nil {
			return nil, err
		}
		conf.Module.Path = absPath

		if data, err := os.ReadFile(path.Join(absPath, "go.mod")); err != nil && !os.IsNotExist(err) {
			return nil, err
		} else if err == nil {
			modFile, err := modfile.Parse(path.Join(absPath, "go.mod"), data, nil)
			if err != nil {
				return nil, err
			}
			conf.Module.ModFile = modFile
		}
	}
	return conf, nil
}

func loadConfig(yamlBytes []byte) (conf *Config, err error) {
	conf = &Config{}
	yamlErr := yaml.Unmarshal(yamlBytes, conf)
	if yamlErr != nil {
		err = errors.New("could not parse yaml: " + yamlErr.Error())
		return
	}
	for goPackage, mapping := range conf.Mappings {
		mapping.GoPackage = goPackage
	}
	return
}
