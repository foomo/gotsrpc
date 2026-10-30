// Package gorpcconv provides OpenTelemetry instrumentation for the gotsrpc
// binary (valyala/gorpc) transport. The gorpc protocol has no header carrier,
// so there is no cross-process trace-context propagation: spans are local
// (parentless) and carry the rpc.* attributes plus a duration metric.
package gorpcconv
