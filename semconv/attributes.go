package semconv

import (
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
	otelsemconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

// gotsrpc-specific attribute keys (no OpenTelemetry equivalent). Standard RPC
// and error attributes are reused from otelsemconv (see RPCSystemGoTSRPC,
// RPCMethod, ErrorType below) rather than redefined here.
const (
	PackageKey   = attribute.Key("gotsrpc.package")
	EncodingKey  = attribute.Key("gotsrpc.encoding")
	ErrorCodeKey = attribute.Key("gotsrpc.error.code")

	// Span-only serialization timing/size attributes (microseconds / bytes).
	ExecutionMicrosKey = attribute.Key("gotsrpc.execution_us")
	MarshalMicrosKey   = attribute.Key("gotsrpc.marshal_us")
	UnmarshalMicrosKey = attribute.Key("gotsrpc.unmarshal_us")
	RequestSizeKey     = attribute.Key("gotsrpc.request_size")
	ResponseSizeKey    = attribute.Key("gotsrpc.response_size")
)

// RPCSystemGoTSRPC is the OpenTelemetry rpc.system.name attribute identifying
// gotsrpc, set on every span/metric.
var RPCSystemGoTSRPC = otelsemconv.RPCSystemNameKey.String("gotsrpc")

// RPCMethod returns the OpenTelemetry rpc.method attribute. Following the RPC
// semantic conventions it is the fully-qualified logical method name
// "<service>/<method>" (the service is not a separate attribute).
func RPCMethod(service, method string) attribute.KeyValue {
	return otelsemconv.RPCMethod(service + "/" + method)
}

// ErrorType returns the OpenTelemetry error.type attribute derived from err
// (handles wrapping, an optional ErrorType() string method, and reflection).
func ErrorType(err error) attribute.KeyValue {
	return otelsemconv.ErrorType(err)
}

// ErrorTypeString returns the OpenTelemetry error.type attribute for a
// low-cardinality string identifier, for failures that have no error value
// (e.g. a non-2xx transport status or an unknown method).
func ErrorTypeString(t string) attribute.KeyValue {
	return otelsemconv.ErrorTypeKey.String(t)
}

// Package returns the gotsrpc.package attribute (the Go package of the service).
func Package(pkg string) attribute.KeyValue { return PackageKey.String(pkg) }

// Encoding returns the gotsrpc.encoding attribute ("msgpack", "json" or "gob").
func Encoding(encoding string) attribute.KeyValue { return EncodingKey.String(encoding) }

// GoTSRPCErrorCode returns the gotsrpc.error.code attribute.
func GoTSRPCErrorCode(code int) attribute.KeyValue { return ErrorCodeKey.Int(code) }

// ExecutionMicros returns the gotsrpc.execution_us span attribute.
func ExecutionMicros(d time.Duration) attribute.KeyValue {
	return ExecutionMicrosKey.Int64(d.Microseconds())
}

// MarshalMicros returns the gotsrpc.marshal_us span attribute.
func MarshalMicros(d time.Duration) attribute.KeyValue {
	return MarshalMicrosKey.Int64(d.Microseconds())
}

// UnmarshalMicros returns the gotsrpc.unmarshal_us span attribute.
func UnmarshalMicros(d time.Duration) attribute.KeyValue {
	return UnmarshalMicrosKey.Int64(d.Microseconds())
}

// RequestSizeAttr returns the gotsrpc.request_size span attribute (bytes).
func RequestSizeAttr(size int) attribute.KeyValue { return RequestSizeKey.Int(size) }

// ResponseSizeAttr returns the gotsrpc.response_size span attribute (bytes).
func ResponseSizeAttr(size int) attribute.KeyValue { return ResponseSizeKey.Int(size) }

// EncodingFromContentType maps a transport content-type to a low-cardinality
// gotsrpc.encoding value.
func EncodingFromContentType(contentType string) string {
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		return "json"
	case strings.HasPrefix(contentType, "application/msgpack"):
		return "msgpack"
	default:
		return "unknown"
	}
}
