package gotsrpc_test

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/gotsrpc/v2/config"
	"github.com/invopop/jsonschema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	cwd, err := os.Getwd()
	require.NoError(t, err)

	reflector := new(jsonschema.Reflector)
	reflector.RequiredFromJSONSchemaTags = true
	require.NoError(t, reflector.AddGoComments("github.com/foomo/gotsrpc/v2/config", "./"))
	schema := reflector.Reflect(&config.Config{})
	actual, err := json.MarshalIndent(schema, "", "  ")
	require.NoError(t, err)

	filename := path.Join(cwd, "gotsrpc.schema.json")
	expected, err := os.ReadFile(filename)
	if !errors.Is(err, os.ErrNotExist) {
		require.NoError(t, err)
	}

	if !assert.Equal(t, string(expected), string(actual)) {
		require.NoError(t, os.WriteFile(filename, actual, 0600))
	}
}
