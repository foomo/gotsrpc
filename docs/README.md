# gotsrpc Documentation

This directory contains comprehensive documentation for gotsrpc, designed to help developers and AI assistants implement client/server code effectively.

## Documentation Overview

### For Developers
- **[Configuration Guide](CONFIGURATION_GUIDE.md)** - Complete guide to configuring gotsrpc.yml files
- **[Context Function Pattern](CONTEXT_FUNCTION_PATTERN.md)** - Special authentication pattern documentation
- **[Troubleshooting Guide](TROUBLESHOOTING.md)** - Common issues and solutions

### For AI Assistants
- **[AI Assistant Guide](AI_ASSISTANT_GUIDE.md)** - Patterns, decision trees, and implementation workflows
- **[Configuration Guide](CONFIGURATION_GUIDE.md)** - Detailed configuration examples
- **[Troubleshooting Guide](TROUBLESHOOTING.md)** - Error diagnosis and solutions

## Quick Start

### 1. Basic Configuration
```yaml
# gotsrpc.yml
module:
  name: github.com/your-org/project
  path: ./

targets:
  api:
    services:
      /api: APIService
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - APIService

mappings:
  github.com/your-org/project/service:
    out: ./client/src/api-vo.ts
```

### 2. Generate Code
```bash
gotsrpc gotsrpc.yml
```

### 3. Implement Service
```go
type APIService interface {
    GetData(id string) (Data, error)
    CreateData(data Data) (Data, error)
}
```

## Key Concepts

### Context Function Pattern
The Context function is a special gotsrpc pattern for authentication:

```go
type Service interface {
    Context(w http.ResponseWriter, r *http.Request)  // Called first
    GetData() (Data, error)                          // Protected method
}
```

### Configuration Structure
- **module**: Go module configuration
- **targets**: Client/server generation targets
- **mappings**: TypeScript value object mappings

### Service Types
- **tsrpc**: Go-to-TypeScript RPC (web clients)
- **gorpc**: Go-to-Go RPC (microservices)

## Common Patterns

### Authentication Service
```yaml
targets:
  auth:
    services:
      /auth: AuthService
    package: github.com/your-org/project/service
    out: ./client/src/auth-client.ts
    tsrpc:
      - AuthService
```

### Multi-Service API
```yaml
targets:
  api:
    services:
      /auth: AuthService
      /api: APIService
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - AuthService
      - APIService
```

### Microservices
```yaml
targets:
  auth:
    services:
      /auth: AuthService
    package: github.com/your-org/auth/service
    out: ./clients/auth/src/auth-client.ts
    tsrpc:
      - AuthService

  api:
    services:
      /api: APIService
    package: github.com/your-org/api/service
    out: ./clients/api/src/api-client.ts
    tsrpc:
      - APIService
```

## Troubleshooting

### Common Issues
1. **Configuration errors** - Check YAML syntax and paths
2. **Package not found** - Verify package paths match Go module
3. **Generated code errors** - Check service interfaces exist
4. **HTTP routing issues** - Match service paths with handlers

### Debug Commands
```bash
# Validate configuration
gotsrpc -debug gotsrpc.yml

# Check generated code
cat service/*_gen.go
cat client/src/*_gen.ts
```

## Examples

See the `example/` directory for complete working examples:
- **basic** - Simple API service
- **auth** - Authentication with Context function
- **errors** - Error handling patterns
- **monitor** - Monitoring and metrics
- **time** - Time handling
- **types** - Complex type mappings

## Documentation

- **[Configuration Guide](./CONFIGURATION_GUIDE.md)** - Complete gotsrpc.yml configuration reference
- **[AI Assistant Guide](./AI_ASSISTANT_GUIDE.md)** - Guidelines for AI assistants implementing gotsrpc
- **[Context Function Pattern](./CONTEXT_FUNCTION_PATTERN.md)** - Authentication patterns
- **[Troubleshooting](./TROUBLESHOOTING.md)** - Common issues and solutions

## Schema Reference

The JSON schema is available at `gotsrpc.schema.json` with enhanced documentation for all configuration options.

## Getting Help

1. Check the [Troubleshooting Guide](TROUBLESHOOTING.md) for common issues
2. Review [Configuration Guide](CONFIGURATION_GUIDE.md) for detailed examples
3. Examine the `example/` directory for working implementations
4. Use `gotsrpc -debug gotsrpc.yml` for detailed error messages

## Contributing

When adding new examples or documentation:
1. Follow the patterns in existing examples
2. Include both server and client code
3. Add comprehensive README files
4. Test with `gotsrpc -debug` before submitting
