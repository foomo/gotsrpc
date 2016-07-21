package config

import "testing"

const sampleConf = `---
targets:
  demo:
    package: github.com/foomo/gotsrpc/demo
    out: /tmp/test.ts 
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
}
