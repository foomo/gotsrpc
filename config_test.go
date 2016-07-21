package gotsrpc

import "testing"

const sampleConf = `---
mappings:
  foo/bar:
    module: Sample.Module
    dir: path/to/ts
  github.com/foomo/gotsrpc:
    module: Sample.Module.RPC
    dir: path/to/other/folder

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
	if foo.TypeScriptDir != "path/to/ts" || foo.TypeScriptModule != "Sample.Module" {
		t.Fatal("unexpected data", foo)
	}
}
