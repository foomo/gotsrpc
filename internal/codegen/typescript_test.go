package codegen_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v3/internal/codegen"
)

func TestSplit(t *testing.T) {
	res := codegen.Split("git.bestbytes.net/foo-bar", []string{".", "/", "-"})
	for i, expected := range []string{"git", "bestbytes", "net", "foo", "bar"} {
		actual := res[i]
		if actual != expected {
			t.Fatal("expected", expected, "got", actual)
		}
	}
}
