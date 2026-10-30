package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"slices"

	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v2"
)

// Target selects the services of one Go package to generate and the output
// files to write.
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
	// Skip generating the Go client for tsrpc services
	SkipTSRPCClient bool `json:"skipTSRPCClient" yaml:"skipTSRPCClient"`
	// Optional override of the service name used in metrics/spans (span name and
	// rpc.method). Keyed by the service name from `services`. Falls back to the
	// service name when absent. Does NOT affect generated Go type names or routing.
	ServiceNames map[string]string `json:"serviceNames" yaml:"serviceNames"`
}

// IsGoRPC reports whether gorpc code is generated for service.
func (t *Target) IsGoRPC(service string) bool {
	return slices.Contains(t.GoRPC, service)
}

// ServiceName returns the telemetry display name for a service, falling back to
// the service name itself when no override is configured.
func (t *Target) ServiceName(service string) string {
	if name, ok := t.ServiceNames[service]; ok && name != "" {
		return name
	}

	return service
}

// IsTSRPC reports whether tsrpc code is generated for service. An empty
// TSRPC list selects all services.
func (t *Target) IsTSRPC(service string) bool {
	if len(t.TSRPC) == 0 {
		return true
	}

	return slices.Contains(t.TSRPC, service)
}

// Mapping maps one Go package to a TypeScript output file.
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

// TypeScriptMappings maps Go package paths to their TypeScript [Mapping].
type TypeScriptMappings map[string]*Mapping

// Namespace describes the Go module that contains the target packages.
type Namespace struct {
	// Go module name
	Name string `json:"name" yaml:"name"`
	// Go module path
	Path string `json:"path" yaml:"path"`
	// Internally loaded mod file
	ModFile *modfile.File `json:"-" yaml:"-"`
}

// Config is the parsed content of a gotsrpc.yml file.
type Config struct {
	// Go module settings
	Module Namespace `json:"module" yaml:"module"`
	// Map of target names to target settings
	Targets map[string]*Target `json:"targets" yaml:"targets"`
	// Map of go module names to TypeScript mapping settings
	Mappings TypeScriptMappings `json:"mappings" yaml:"mappings"`
	// When true, appends .js extension to relative TypeScript import paths.
	// Required for TypeScript projects using moduleResolution 'nodenext' or 'node16'.
	TSImportJsExtension bool `json:"tsImportJsExtension" yaml:"tsImportJsExtension"`
}

// LoadConfigFile reads and parses the YAML config at file. A relative
// module path is resolved against the directory of file, and the module's
// go.mod is loaded into Module.ModFile if it exists.
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
