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
	// Go package name
	Package string `json:"package" yaml:"package"`
	// Map of default routes to service names
	Services map[string]string `json:"services" yaml:"services"`
	// TypeScript module name
	TypeScriptModule string `json:"module" yaml:"module"`
	// TypeScript output filename
	Out string `json:"out" yaml:"out"`
	// List of go rpc services to generate
	GoRPC []string `json:"gorpc" yaml:"gorpc"`
	// List of ts rpc services to generate
	TSRPC []string `json:"tsrpc" yaml:"tsrpc"`
	// Skip generating go rpc client
	SkipTSRPCClient bool `json:"skipTSRPCClient" yaml:"skipTSRPCClient"`
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
	// Internal go package name
	GoPackage string `json:"-" yaml:"-"`
	// TypeScript output filename
	Out string `json:"out" yaml:"out"`
	// List of go types to generate
	Structs []string `json:"structs" yaml:"structs"`
	// List of go types to generate
	Scalars []string `json:"scalars" yaml:"scalars"`
	// Optional TypeScript module name
	TypeScriptModule string `json:"module" yaml:"module"`
}

type TypeScriptMappings map[string]*Mapping

type Namespace struct {
	// Go module name
	Name string `json:"name" yaml:"name"`
	// Go module path
	Path string `json:"path" yaml:"path"`
	// Internally loaded mod file
	ModFile *modfile.File `json:"-" yaml:"-"`
}

type Config struct {
	// Go module settings
	Module Namespace `json:"module" yaml:"module"`
	// Map of target names to target settings
	Targets map[string]*Target `json:"targets" yaml:"targets"`
	// Map of go module names to TypeScript mapping settings
	Mappings TypeScriptMappings `json:"mappings" yaml:"mappings"`
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
