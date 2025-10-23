[![Build Status](https://github.com/foomo/gotsrpc/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/gotsrpc/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/gotsrpc)](https://goreportcard.com/report/github.com/foomo/gotsrpc)
[![GoDoc](https://godoc.org/github.com/foomo/gotsrpc?status.svg)](https://godoc.org/github.com/foomo/gotsrpc)

<p align="center">
  <img alt="gotsrpc" src=".github/assets/gotsrpc.png"/>
</p>

# Go / TypeScript and Go / Go RPC

## Documentation

Please refer to the [documentation](https://www.foomo.org/docs/projects/gotsrpc).

### Enhanced Documentation

This repository now includes comprehensive documentation in the `docs/` directory:

- **[📖 Documentation Overview](docs/README.md)** - Complete guide to all documentation
- **[⚙️ Configuration Guide](docs/CONFIGURATION_GUIDE.md)** - Detailed gotsrpc.yml configuration reference
- **[🤖 AI Assistant Guide](docs/AI_ASSISTANT_GUIDE.md)** - Guidelines for AI assistants implementing gotsrpc
- **[🔐 Context Function Pattern](docs/CONTEXT_FUNCTION_PATTERN.md)** - Authentication patterns and best practices
- **[🔧 Troubleshooting Guide](docs/TROUBLESHOOTING.md)** - Common issues and solutions

**Quick Start**: Start with the [Documentation Overview](docs/README.md) for a complete introduction to all available guides.

```shell
$ gotsrpc
gotsrpc

Usage:
  gotsrpc [options] <config-file>

Options:
  -version   Display version information
  -debug     Print debug information

Examples:
  $ gotsrpc path/to/gotsrpc.yaml
```

## Installation

Please follow the [documentation](https://www.foomo.org/docs/projects/gotsrpc/cli#installation).

### Download binary

Download a [binary release](https://github.com/foomo/gotsrpc/releases)

### Build from source

```
go install github.com/foomo/gotsrpc@latest
```

### Homebrew (Linux/macOS)

If you use [Homebrew](https://brew.sh), you can install like this:
```
brew install foomo/tap/gotsrpc
```

### Mise

If you use [mise](https://https://mise.jdx.dev), you can install like this:
```
mise use github.com:foomo/gotsrpc
```

Release downloads:

[https://github.com/foomo/gotsrpc/releases](https://github.com/foomo/gotsrpc/releases)

## How to Contribute

Please refer to the [CONTRIBUTING](.gihub/CONTRIBUTING.md) details and follow the [CODE_OF_CONDUCT](.gihub/CODE_OF_CONDUCT.md) and [SECURITY](.github/SECURITY.md) guidelines.

## License

Distributed under MIT License, please see license file within the code for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
