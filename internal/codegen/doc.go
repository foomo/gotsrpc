// Package codegen renders Go and TypeScript source code from the
// [model] representation of services and types.
//
// The Go renderers produce gotsrpc HTTP proxies and clients and gorpc binary
// proxies and clients. The TypeScript renderers produce service client
// classes and type definitions grouped by Go package.
package codegen
