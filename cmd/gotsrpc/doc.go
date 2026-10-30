// Gotsrpc generates type-safe RPC bindings between Go services and TypeScript
// clients, with optional Go-to-Go RPC support.
//
// Usage:
//
//	gotsrpc [flags] [config-file]
//	gotsrpc generate [flags] [config-file]
//	gotsrpc version [flags]
//
// A bare invocation with a config file is shorthand for the generate
// subcommand. Without a config file, gotsrpc.yml in the current directory is
// used.
//
// The flags are:
//
//	--debug
//		Enable debug output. Also prints commit and build time on version.
//	--log-level level
//		Set the log level: debug, info, warn or error. Defaults to info.
//
// Every flag can also be set through an environment variable with the
// GOTSRPC_ prefix, for example GOTSRPC_LOG_LEVEL=debug.
package main
