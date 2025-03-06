package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sampleConf = `---
targets:
  demo:
    services:
      /service/demo: Service
    package: github.com/foomo/gotsrpc/v2/demo
    module: My.Service
    modulekind: commonjs
    out: /tmp/my-service.ts 
mappings:
  foo/bar:
    module: Sample.Module
    out: path/to/ts
  github.com/foomo/gotsrpc/v2:
    module: Sample.Module.RPC
    out: path/to/other/folder

`

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig([]byte(sampleConf))
	if err != nil {
		t.Fatal(err)
	}
	goPackage := "foo/bar"
	foo, ok := c.Mappings[goPackage]
	if !ok {
		t.Fatal("foo/bar not found")
	}

	if foo.GoPackage != goPackage {
		t.Fatal("wrong go package value")
	}
	if foo.Out != "path/to/ts" || foo.TypeScriptModule != "Sample.Module" {
		t.Fatal("unexpected data", foo)
	}

	// looking at the targets

	demoTarget, ok := c.Targets["demo"]
	if !ok {
		t.Fatal("demo target not found")
	}
	if demoTarget.Out != "/tmp/my-service.ts" {
		t.Fatal("demo target out is wrong")
	}
	if demoTarget.Package != "github.com/foomo/gotsrpc/v2/demo" {
		t.Fatal("wrong target package")
	}
	if demoTarget.TypeScriptModule != "My.Service" {
		t.Fatal("wromg ts module")
	}
	if len(demoTarget.Services) != 1 {
		t.Fatal("wrong number of services")
	}
	if demoTarget.Services["/service/demo"] != "Service" {
		t.Fatal("first service is wrong")
	}
}

func TestLoadConfigFile_GomodAbsolute(t *testing.T) {
	config, err := LoadConfigFile("testdata/gomod.absolute.yml")
	require.NoError(t, err)
	assert.Equal(t, "/go/src/github.com/foomo/gotsrpc", config.Module.Path)
}

func TestLoadConfigFile_GomodRelative(t *testing.T) {
	config, err := LoadConfigFile("testdata/gomod.relative.yml")
	require.NoError(t, err)
	fmt.Println(config.Module.Path)
	assert.True(t, strings.HasSuffix(config.Module.Path, "gotsrpc/config/demo"))
}
