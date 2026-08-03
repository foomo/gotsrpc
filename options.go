package gotsrpc

import (
	"log/slog"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/foomo/gotsrpc/v3/semconv"
)

// Option configures the OpenTelemetry instrumentation of generated proxies and
// clients. It is an alias of semconv.Option so it can be passed straight
// through to the httpconv/gorpcconv constructors.
type Option = semconv.Option

// WithTracerProvider sets the trace.TracerProvider (defaults to the global).
func WithTracerProvider(tp trace.TracerProvider) Option {
	return semconv.WithTracerProvider(tp)
}

// WithMeterProvider sets the metric.MeterProvider (defaults to the global).
func WithMeterProvider(mp metric.MeterProvider) Option {
	return semconv.WithMeterProvider(mp)
}

// WithPropagators sets the propagation.TextMapPropagator (defaults to the global).
func WithPropagators(p propagation.TextMapPropagator) Option {
	return semconv.WithPropagators(p)
}

// WithLogger sets the slog.Logger used for minimal, error-only log records
// (defaults to slog.Default()).
func WithLogger(l *slog.Logger) Option {
	return semconv.WithLogger(l)
}

// WithClientService seeds the rpc.service / gotsrpc.package identity on a
// generated client. Generated client constructors pass this automatically.
func WithClientService(pkg, service string) Option {
	return semconv.WithClientService(pkg, service)
}
