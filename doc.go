// Package gotsrpc is the runtime library used by code that the gotsrpc
// generator produces.
//
// Generated proxies decode requests with [LoadArgs], call the service and
// encode results with [Reply]. Generated clients send calls through a
// [Client]. Both sides support msgpack and JSON encodings (see
// [ClientEncoding]) and emit OpenTelemetry rpc.* spans and metrics configured
// through [Option].
//
// # Errors
//
// Errors returned by service methods travel over the wire as [Error], which
// keeps the message, the Go type, the package path, the error data and the
// cause chain. Use [errors.Is] and [errors.As] on the client side as with a
// local error. Transport failures on the client side are reported as
// [ClientError], and non-200 responses as [HTTPError].
//
// # Union types
//
// Structs that model a union (one non-nil pointer field out of several) must
// be registered with [RegisterUnionExt] so that JSON encoding emits only the
// set member.
package gotsrpc
