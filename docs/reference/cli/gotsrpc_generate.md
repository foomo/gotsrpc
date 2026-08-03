---
title: gotsrpc generate
---

## gotsrpc generate

Generate RPC bindings from a gotsrpc config file

### Synopsis

Generate RPC bindings from a gotsrpc config file.

If no config file is given, gotsrpc.yml in the current directory is used.

The referenced Go code has to compile and mappings for all used packages must
be configured. Previously generated files are overwritten; obsolete files are
not removed, so add a clean step to your build if needed.

```
gotsrpc generate [config-file] [flags]
```

### Examples

```
  # Use gotsrpc.yml in the current directory
  gotsrpc generate

  # Use an explicit config file
  gotsrpc generate path/to/gotsrpc.yml

  # Increase log verbosity
  gotsrpc generate --log-level debug
```

### Options

```
  -h, --help   help for generate
```

### Options inherited from parent commands

```
      --debug              enable debug output
      --log-level string   log level (debug, info, warn, error) (default "info")
```

### SEE ALSO

* [gotsrpc](./gotsrpc)	 - Generate type-safe RPC bindings between Go services and TypeScript clients

