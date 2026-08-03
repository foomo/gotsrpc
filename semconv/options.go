package semconv

import (
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Config holds the resolved instrumentation configuration.
type Config struct {
	TracerProvider trace.TracerProvider
	MeterProvider  metric.MeterProvider
	Propagators    propagation.TextMapPropagator
	Logger         *slog.Logger
	// Package and Service identify the instrumented service. Servers set them
	// via NewServer arguments; clients seed them through WithClientService.
	Package string
	Service string
}

// Option configures instrumentation. It mirrors the standard OpenTelemetry Go
// functional-option style.
type Option interface {
	apply(c *Config)
}

type optionFunc func(*Config)

func (f optionFunc) apply(c *Config) { f(c) }

// NewConfig resolves a Config from the given options, defaulting every
// provider to its OpenTelemetry global (a no-op until an SDK is installed) and
// the logger to slog.Default().
func NewConfig(opts ...Option) Config {
	c := Config{
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  otel.GetMeterProvider(),
		Propagators:    otel.GetTextMapPropagator(),
		Logger:         slog.Default(),
	}

	for _, o := range opts {
		o.apply(&c)
	}

	// Guard against options that were passed an explicit nil.
	if c.TracerProvider == nil {
		c.TracerProvider = otel.GetTracerProvider()
	}

	if c.MeterProvider == nil {
		c.MeterProvider = otel.GetMeterProvider()
	}

	if c.Propagators == nil {
		c.Propagators = otel.GetTextMapPropagator()
	}

	if c.Logger == nil {
		c.Logger = slog.Default()
	}

	return c
}

// WithTracerProvider sets the trace.TracerProvider. Defaults to the global.
func WithTracerProvider(tp trace.TracerProvider) Option {
	return optionFunc(func(c *Config) { c.TracerProvider = tp })
}

// WithMeterProvider sets the metric.MeterProvider. Defaults to the global.
func WithMeterProvider(mp metric.MeterProvider) Option {
	return optionFunc(func(c *Config) { c.MeterProvider = mp })
}

// WithPropagators sets the propagation.TextMapPropagator used for trace-context
// propagation. Defaults to the global.
func WithPropagators(p propagation.TextMapPropagator) Option {
	return optionFunc(func(c *Config) { c.Propagators = p })
}

// WithLogger sets the slog.Logger used for the (minimal, error-only) log
// records. Defaults to slog.Default().
func WithLogger(l *slog.Logger) Option {
	return optionFunc(func(c *Config) { c.Logger = l })
}

// WithClientService seeds the rpc.service identity on a client. Generated
// clients pass this so the RPC client span/metrics carry the correct
// rpc.service and gotsrpc.package attributes.
func WithClientService(pkg, service string) Option {
	return optionFunc(func(c *Config) {
		c.Package = pkg
		c.Service = service
	})
}
