package semconv_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v3/semconv"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Parallel()
	// In the test binary the module version is typically unreadable, so the
	// fallback "unknown" is expected; either way it must be non-empty.
	assert.NotEmpty(t, semconv.Version())
}

func TestTracerAndMeter(t *testing.T) {
	t.Parallel()

	cfg := semconv.NewConfig()
	assert.NotNil(t, semconv.Tracer(cfg))
	assert.NotNil(t, semconv.Meter(cfg))
}
