// Package config loads and describes the gotsrpc.yml file that drives code
// generation.
//
// A [Config] names the Go module, a set of [Target] entries that select the
// services to generate, and [TypeScriptMappings] that map Go packages to
// TypeScript output files. Load it with [LoadConfigFile].
package config
