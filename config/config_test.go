package config

import "testing"

const sampleConf = `---
targets:
  demo:
    services:
      /service/demo: Service
    package: github.com/foomo/gotsrpc/demo
    module: My.Service
    modulekind: commonjs
    out: /tmp/my-service.ts 
mappings:
  foo/bar:
    module: Sample.Module
    out: path/to/ts
  github.com/foomo/gotsrpc:
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
	if demoTarget.Package != "github.com/foomo/gotsrpc/demo" {
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
