// Package httpconv provides OpenTelemetry instrumentation for the gotsrpc
// HTTP transport. It emits only rpc.* spans and metrics (never http.*), so it
// composes cleanly with otelhttp: when a consumer wraps the handler with
// otelhttp.NewHandler and/or the client transport with otelhttp.NewTransport,
// the gotsrpc RPC spans nest under the otelhttp HTTP spans and no metrics are
// duplicated.
package httpconv
