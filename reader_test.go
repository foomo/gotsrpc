package gotsrpc

import (
	"github.com/foomo/gotsrpc/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestServiceList(t *testing.T) ServiceList {
	// ReaderTrace = true
	serviceMap := map[string]string{
		"/demo": "Demo",
	}

	packageName := "github.com/foomo/gotsrpc/demo"

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

func TestInterfaceReading(t *testing.T) {
	services := getTestServiceList(t)
	//spew.Dump(services)
	demoService := services.getServiceByName("Demo")
	assert.NotNil(t, demoService)
	helloInterfaceMethod := demoService.Methods.getMethodByName("HelloInterface")
	assert.NotNil(t, helloInterfaceMethod)
	assert.True(t, helloInterfaceMethod.Args[0].Value.IsInterface)
	assert.True(t, helloInterfaceMethod.Args[1].Value.Map.Value.IsInterface)
	assert.True(t, helloInterfaceMethod.Args[2].Value.Array.Value.IsInterface)
}
