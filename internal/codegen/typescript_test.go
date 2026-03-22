package codegen

import "testing"

func TestSplit(t *testing.T) {
	res := Split("git.bestbytes.net/foo-bar", []string{".", "/", "-"})
	for i, expected := range []string{"git", "bestbytes", "net", "foo", "bar"} {
		actual := res[i]
		if actual != expected {
			t.Fatal("expected", expected, "got", actual)
		}
	}
}
