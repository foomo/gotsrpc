package gotsrpc

import "testing"

func TestSplit(t *testing.T) {
	res := split("git.bestbytes.net/foo-bar", []string{".", "/", "-"})
	for i, expected := range []string{"git", "bestbytes", "net", "foo", "bar"} {
		actual := res[i]
		if actual != expected {
			t.Fatal("expected", expected, "got", actual)
		}
	}

}
