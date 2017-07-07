package gotsrpc

import (
	"testing"
)

func TestLoadArgs(t *testing.T) {
	jsonBytes := []byte(`["a", ["a", "b", "c"]]`)
	foo := ""
	bar := []string{}
	args := []interface{}{&foo, &bar}
	errLoad := loadArgs(&args, jsonBytes)
	if errLoad != nil {
		t.Fatal(errLoad)
	}
	if foo != "a" {
		t.Fatal("foo should have been a")
	}
	if len(bar) != 3 {
		t.Fatal("bar len wrong", len(bar), "!=", len(bar))
	}
	if bar[1] != "b" {
		t.Fatal("bar[1] (", bar[1], ") != b")
	}
}
