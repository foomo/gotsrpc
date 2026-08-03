package semconv_test

import (
	"log/slog"
	"testing"

	"github.com/foomo/gotsrpc/v3/semconv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestNewConfig_Defaults(t *testing.T) {
	t.Parallel()

	cfg := semconv.NewConfig()
	assert.Equal(t, otel.GetTracerProvider(), cfg.TracerProvider)
	assert.Equal(t, otel.GetMeterProvider(), cfg.MeterProvider)
	assert.Equal(t, otel.GetTextMapPropagator(), cfg.Propagators)
	assert.NotNil(t, cfg.Logger)
}

func TestNewConfig_NilOptionsFallBackToGlobals(t *testing.T) {
	t.Parallel()

	cfg := semconv.NewConfig(
		semconv.WithTracerProvider(nil),
		semconv.WithMeterProvider(nil),
		semconv.WithPropagators(nil),
		semconv.WithLogger(nil),
	)
	assert.NotNil(t, cfg.TracerProvider)
	assert.NotNil(t, cfg.MeterProvider)
	assert.NotNil(t, cfg.Propagators)
	assert.NotNil(t, cfg.Logger)
}

func TestNewConfig_Options(t *testing.T) {
	t.Parallel()

	tp := sdktrace.NewTracerProvider()
	mp := sdkmetric.NewMeterProvider()
	prop := propagation.TraceContext{}
	logger := slog.Default()

	cfg := semconv.NewConfig(
		semconv.WithTracerProvider(tp),
		semconv.WithMeterProvider(mp),
		semconv.WithPropagators(prop),
		semconv.WithLogger(logger),
		semconv.WithClientService("github.com/foo/bar", "UserService"),
	)

	assert.Same(t, tp, cfg.TracerProvider)
	assert.Same(t, mp, cfg.MeterProvider)
	assert.Equal(t, prop, cfg.Propagators)
	assert.Same(t, logger, cfg.Logger)
	assert.Equal(t, "github.com/foo/bar", cfg.Package)
	assert.Equal(t, "UserService", cfg.Service)

	require.NoError(t, tp.Shutdown(t.Context()))
	require.NoError(t, mp.Shutdown(t.Context()))
}
