// Package semconv defines the OpenTelemetry semantic conventions used by
// gotsrpc: the RPC attribute keys/values, the metric instruments, and the
// functional options + provider resolution shared by the transport-specific
// instrumentation packages (httpconv, gorpcconv).
//
// It depends only on the OpenTelemetry API (never on an SDK or on otelhttp),
// so that generated proxies and clients emit rpc.* signals exclusively and
// never duplicate the http.* signals contributed by otelhttp.
package semconv
