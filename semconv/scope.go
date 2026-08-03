// Package semconv defines the OpenTelemetry semantic conventions used by
// gotsrpc: the RPC attribute keys/values, the metric instruments, and the
// functional options + provider resolution shared by the transport-specific
// instrumentation packages (httpconv, gorpcconv).
//
// It depends only on the OpenTelemetry API (never on an SDK or on otelhttp),
// so that generated proxies and clients emit rpc.* signals exclusively and
// never duplicate the http.* signals contributed by otelhttp.
package semconv

import (
	"runtime/debug"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// ScopeName is the instrumentation scope reported on all emitted telemetry.
const ScopeName = "github.com/foomo/gotsrpc/v3"

// Version returns the module version of gotsrpc if it can be read from the
// build info, otherwise "unknown". It is used as the instrumentation version.
func Version() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range bi.Deps {
			if dep.Path == ScopeName && dep.Version != "" {
				return dep.Version
			}
		}
	}

	return "unknown"
}

// Tracer returns the tracer for the gotsrpc instrumentation scope.
func Tracer(cfg Config) trace.Tracer {
	return cfg.TracerProvider.Tracer(ScopeName, trace.WithInstrumentationVersion(Version()))
}

// Meter returns the meter for the gotsrpc instrumentation scope.
func Meter(cfg Config) metric.Meter {
	return cfg.MeterProvider.Meter(ScopeName, metric.WithInstrumentationVersion(Version()))
}
