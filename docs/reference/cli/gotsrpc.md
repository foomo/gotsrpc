---
title: gotsrpc
---

## gotsrpc

Generate type-safe RPC bindings between Go services and TypeScript clients

### Synopsis

gotsrpc generates type-safe RPC bindings between Go services and TypeScript
clients, with optional Go-to-Go RPC support. It parses your Go source via the
AST, reads a YAML config (gotsrpc.yml) and writes Go proxies/clients and
TypeScript clients and type definitions.

Running gotsrpc with a config file and no subcommand is shorthand for
`gotsrpc generate <config-file>`.

Flags can also be set via environment variables prefixed with GOTSRPC_, e.g.
GOTSRPC_LOG_LEVEL=debug or GOTSRPC_DEBUG=true.

### Examples

```
  # Generate from gotsrpc.yml in the current directory
  gotsrpc gotsrpc.yml

  # Equivalent, using the explicit subcommand
  gotsrpc generate path/to/gotsrpc.yml

  # Print the version
  gotsrpc version
```

### Options

```
      --debug              enable debug output
  -h, --help               help for gotsrpc
      --log-level string   log level (debug, info, warn, error) (default "info")
```

### SEE ALSO

* [gotsrpc generate](./gotsrpc_generate)	 - Generate RPC bindings from a gotsrpc config file
* [gotsrpc version](./gotsrpc_version)	 - Display version information

