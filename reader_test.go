package gotsrpc

import (
	"fmt"
	"go/build"
	"os"
	"testing"

	"github.com/foomo/gotsrpc/config"

	"github.com/stretchr/testify/assert"
)

func TestServiceList(t *testing.T) {
	c, err := config.LoadConfigFile("demo/config-substruct.yml")
	if err != nil {
		t.Fatal(err)
	}

	target := c.Targets["demo"]

	pkg, parseErr := parsePackage([]string{os.Getenv("GOPATH"), build.Default.GOPATH}, c.Module, target.Package)
	assert.NoError(t, parseErr)

	services, err := readServicesInPackage(pkg, target.Package, target.Services)
	assert.NoError(t, err)

	missingTypes := map[string]bool{}
	for _, s := range services {
		for _, m := range s.Methods {
			collectStructTypes(m.Return, missingTypes)
			collectStructTypes(m.Args, missingTypes)
			collectScalarTypes(m.Return, missingTypes)
			collectScalarTypes(m.Args, missingTypes)
		}
	}
	fmt.Println(missingTypes)
}
