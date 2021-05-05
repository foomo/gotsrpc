package gotsrpc

import (
	"github.com/foomo/gotsrpc/v2/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestServiceList(t *testing.T) ServiceList {
	// ReaderTrace = true
	serviceMap := map[string]string{
		"/demo": "Demo",
	}

	packageName := "github.com/foomo/gotsrpc/v2/demo"

	pkg, parseErr := parsePackage([]string{os.Getenv("GOPATH")}, config.Namespace{}, packageName)
	assert.NoError(t, parseErr)

	services, err := readServicesInPackage(pkg, packageName, serviceMap)
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
	return services
}
