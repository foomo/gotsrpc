package gotsrpc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	ReaderTrace = true
	serviceMap := map[string]string{
		"/demo": "Demo",
	}

	packageName := "github.com/foomo/gotsrpc/demo"

	pkg, parseErr := parsePackage([]string{os.Getenv("GOPATH")}, packageName)
	assert.NoError(t, parseErr)

	services, err := readServicesInPackage(pkg, packageName, serviceMap)

	t.Log(services, err)

	missingTypes := map[string]bool{}
	for _, s := range services {
		for _, m := range s.Methods {
			collectStructTypes(m.Return, missingTypes)
			collectStructTypes(m.Args, missingTypes)
			collectScalarTypes(m.Return, missingTypes)
			collectScalarTypes(m.Args, missingTypes)
		}
	}
	//spew.Dump(services)

}
