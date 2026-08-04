package config //nolint:testpackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTarget_IsGoRPC(t *testing.T) {
	t.Parallel()

	target := &Target{GoRPC: []string{"UserService"}}
	assert.True(t, target.IsGoRPC("UserService"))
	assert.False(t, target.IsGoRPC("OtherService"))
	assert.False(t, (&Target{}).IsGoRPC("UserService"))
}

func TestTarget_IsTSRPC(t *testing.T) {
	t.Parallel()

	// empty list means all services are TSRPC
	assert.True(t, (&Target{}).IsTSRPC("Any"))

	target := &Target{TSRPC: []string{"UserService"}}
	assert.True(t, target.IsTSRPC("UserService"))
	assert.False(t, target.IsTSRPC("OtherService"))
}

func TestTarget_ServiceName(t *testing.T) {
	t.Parallel()

	target := &Target{ServiceNames: map[string]string{
		"Service": "Monitor",
		"Empty":   "",
	}}

	// override present -> display name
	assert.Equal(t, "Monitor", target.ServiceName("Service"))
	// override absent -> falls back to the service name
	assert.Equal(t, "OtherService", target.ServiceName("OtherService"))
	// empty override -> falls back to the service name
	assert.Equal(t, "Empty", target.ServiceName("Empty"))
	// no map at all -> falls back to the service name
	assert.Equal(t, "Service", (&Target{}).ServiceName("Service"))
}

func TestLoadConfig_MalformedYAML(t *testing.T) {
	t.Parallel()

	_, err := loadConfig([]byte("targets: [this is not: valid yaml"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not parse yaml")
}

func TestLoadConfigFile_MissingFile(t *testing.T) {
	t.Parallel()

	_, err := LoadConfigFile("testdata/does-not-exist.yml")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not read config file")
}
