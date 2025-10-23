# gotsrpc Troubleshooting Guide

This guide helps diagnose and fix common issues when using gotsrpc, especially for AI assistants implementing client/server code.

## Table of Contents
1. [Configuration Issues](#configuration-issues)
2. [Code Generation Problems](#code-generation-problems)
3. [Runtime Errors](#runtime-errors)
4. [Authentication Issues](#authentication-issues)
5. [Debugging Techniques](#debugging-techniques)

## Configuration Issues

### 1. Module Path Errors

#### Error: "config load error, could not load config from gotsrpc.yml"
**Symptoms:**
- gotsrpc fails to load configuration
- Error mentions path issues

**Common Causes:**
```yaml
# ❌ Wrong - incorrect module path
module:
  name: github.com/your-org/project
  path: ../  # Should be ./

# ❌ Wrong - path doesn't match go.mod location
module:
  name: github.com/your-org/project
  path: ./subdir/  # If go.mod is in parent directory
```

**Solutions:**
```yaml
# ✅ Correct - path matches go.mod location
module:
  name: github.com/your-org/project
  path: ./

# ✅ Correct - if go.mod is in parent directory
module:
  name: github.com/your-org/project
  path: ../
```

**Debug Steps:**
1. Check where your `go.mod` file is located
2. Ensure the `path` in gotsrpc.yml points to that directory
3. Use `gotsrpc -debug gotsrpc.yml` for detailed error messages

### 2. Package Path Mismatches

#### Error: "package not found" or "import path not found"
**Symptoms:**
- Generated code has import errors
- Go build fails with "package not found"

**Common Causes:**
```yaml
# ❌ Wrong - package doesn't exist
targets:
  api:
    package: github.com/your-org/nonexistent/service

# ❌ Wrong - incorrect package path
targets:
  api:
    package: ./service  # Should be full Go module path
```

**Solutions:**
```yaml
# ✅ Correct - use actual package path from go.mod
targets:
  api:
    package: github.com/your-org/project/service
```

**Debug Steps:**
1. Check your `go.mod` file for the module name
2. Verify the package directory exists
3. Ensure the package path matches your Go module structure

### 3. Service Path Conflicts

#### Error: "duplicate service path" or routing conflicts
**Symptoms:**
- HTTP routing conflicts
- Services not accessible
- 404 errors for expected endpoints

**Common Causes:**
```yaml
# ❌ Wrong - conflicting paths
targets:
  api:
    services:
      /api: Service
      /api/auth: AuthService  # Conflicts with /api

# ❌ Wrong - overlapping paths
targets:
  api:
    services:
      /: MainService
      /api: APIService  # /api conflicts with /
```

**Solutions:**
```yaml
# ✅ Correct - non-conflicting paths
targets:
  api:
    services:
      /api: APIService
      /auth: AuthService
      /admin: AdminService

# ✅ Correct - hierarchical paths
targets:
  api:
    services:
      /api/v1: V1Service
      /api/v2: V2Service
      /auth: AuthService
```

### 4. Output Path Issues

#### Error: "cannot create output file" or "permission denied"
**Symptoms:**
- File creation errors
- Permission denied errors

**Common Causes:**
```yaml
# ❌ Wrong - invalid output path
targets:
  api:
    out: /nonexistent/path/client.ts  # Directory doesn't exist

# ❌ Wrong - relative path issues
targets:
  api:
    out: ../client/src/client.ts  # May not exist
```

**Solutions:**
```yaml
# ✅ Correct - ensure directory exists
targets:
  api:
    out: ./client/src/client.ts

# ✅ Correct - use absolute paths if needed
targets:
  api:
    out: /home/user/project/client/src/client.ts
```

**Debug Steps:**
1. Create the output directory: `mkdir -p ./client/src`
2. Check file permissions
3. Use absolute paths if relative paths cause issues

## Code Generation Problems

### 1. Missing Generated Files

#### Error: Generated files not created
**Symptoms:**
- No `*_gen.go` or `*_gen.ts` files created
- Build fails with missing imports

**Common Causes:**
- Configuration errors (see above)
- No services defined in `tsrpc` or `gorpc` lists
- Invalid service interfaces

**Solutions:**
```yaml
# ✅ Ensure services are listed
targets:
  api:
    services:
      /api: Service
    package: github.com/your-org/project/service
    out: ./client/src/client.ts
    tsrpc:
      - Service  # Must list services to generate
```

**Debug Steps:**
1. Run `gotsrpc -debug gotsrpc.yml` to see detailed output
2. Check that service interfaces exist in the specified package
3. Verify the package path is correct

### 2. Import Errors in Generated Code

#### Error: "import not found" in generated files
**Symptoms:**
- Go build fails with import errors
- TypeScript compilation fails

**Common Causes:**
- Incorrect module path in configuration
- Missing dependencies
- Go module issues

**Solutions:**
```bash
# Ensure Go module is properly initialized
go mod tidy
go mod download

# Regenerate with correct paths
gotsrpc -debug gotsrpc.yml
```

### 3. TypeScript Compilation Errors

#### Error: TypeScript compilation fails
**Symptoms:**
- `tsc` fails with type errors
- Missing type definitions

**Common Causes:**
- Missing value object mappings
- Incorrect TypeScript configuration
- Missing dependencies

**Solutions:**
```yaml
# ✅ Include value object mappings
mappings:
  github.com/your-org/project/service:
    out: ./client/src/vo.ts
```

**Debug Steps:**
1. Check `tsconfig.json` configuration
2. Ensure all dependencies are installed: `npm install`
3. Verify generated TypeScript files are valid

## Runtime Errors

### 1. HTTP Handler Registration Issues

#### Error: "404 Not Found" for service endpoints
**Symptoms:**
- Service endpoints return 404
- Routes not registered properly

**Common Causes:**
```go
// ❌ Wrong - incorrect handler registration
mux.Handle("/api", proxy)  // Missing trailing slash

// ❌ Wrong - wrong path pattern
mux.Handle("/api/", proxy)  // But service expects /api
```

**Solutions:**
```go
// ✅ Correct - match service path from config
mux.Handle("/api/", apiProxy)  // Matches /api: Service in config

// ✅ Correct - use generated proxy functions
authProxy := service.NewDefaultAuthServiceGoTSRPCProxy(authHandler)
mux.Handle("/auth/", authProxy)
```

### 2. Context Function Issues

#### Error: Authentication not working with Context function
**Symptoms:**
- Context function not called
- Authentication bypassed
- User context not available

**Common Causes:**
```go
// ❌ Wrong - incorrect Context function signature
func (h *Handler) Context(w http.ResponseWriter, r *http.Request, extra string) {
    // Extra parameter prevents gotsrpc from recognizing this as Context function
}

// ❌ Wrong - not implementing the interface
type Handler struct{}
// Missing Context method
```

**Solutions:**
```go
// ✅ Correct - exact signature required
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    // Authentication logic
}

// ✅ Correct - implement the interface
type Service interface {
    Context(w http.ResponseWriter, r *http.Request)
    GetData() (Data, error)
}
```

### 3. Client-Side Issues

#### Error: TypeScript client not working
**Symptoms:**
- Client compilation errors
- Runtime errors in browser
- Missing type definitions

**Common Causes:**
- Incorrect transport function
- Missing generated files
- TypeScript configuration issues

**Solutions:**
```typescript
// ✅ Correct transport function
const transport = <T>(endpoint: string) => async (method: string, data: any[] = []): Promise<T> => {
    const url = `http://localhost:8080${endpoint}/${encodeURIComponent(method)}`;
    
    const options: RequestInit = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    };
    
    if (data.length > 0) {
        options.body = JSON.stringify(data);
    }
    
    const response = await fetch(url, options);
    if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${await response.text()}`);
    }
    
    return response.text() ? JSON.parse(await response.text()) : {};
};

// ✅ Correct client usage
const client = new ServiceClient(transport('/api'));
```

## Authentication Issues

### 1. Context Function Not Called

#### Error: Authentication bypassed
**Symptoms:**
- Context function not executed
- Unauthenticated requests succeed

**Debug Steps:**
1. Check function signature: `Context(w http.ResponseWriter, r *http.Request)`
2. Ensure it's part of the service interface
3. Verify the service is listed in `tsrpc` configuration

### 2. Authorization Header Issues

#### Error: "Authorization header required" even with header
**Symptoms:**
- Context function receives empty Authorization header
- Token validation fails

**Common Causes:**
- CORS issues
- Header not sent by client
- Case sensitivity issues

**Solutions:**
```typescript
// ✅ Correct - include Authorization header
const options: RequestInit = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,  // Include token
    },
};
```

### 3. Token Validation Issues

#### Error: "Invalid token" for valid tokens
**Symptoms:**
- Valid tokens rejected
- Authentication always fails

**Debug Steps:**
1. Check token format in Context function
2. Verify token validation logic
3. Check for timing issues (token expiration)

## Debugging Techniques

### 1. Use Debug Mode
```bash
# Enable detailed logging
gotsrpc -debug gotsrpc.yml
```

### 2. Check Generated Code
```bash
# Examine generated files
cat service/*_gen.go
cat client/src/*_gen.ts
```

### 3. Test Configuration
```bash
# Validate configuration
gotsrpc -debug gotsrpc.yml > debug.log 2>&1
cat debug.log
```

### 4. HTTP Request Debugging
```go
// Add logging to Context function
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    log.Printf("Context called: %s %s", r.Method, r.URL.Path)
    log.Printf("Headers: %v", r.Header)
    
    // Authentication logic...
}
```

### 5. Client-Side Debugging
```typescript
// Add logging to transport function
const transport = <T>(endpoint: string) => async (method: string, data: any[] = []): Promise<T> => {
    console.log('Making request to:', endpoint, method, data);
    
    const response = await fetch(url, options);
    console.log('Response:', response.status, await response.text());
    
    return result;
};
```

## Common Error Messages and Solutions

### "config load error, could not load config from gotsrpc.yml"
- **Cause**: Invalid YAML syntax or file not found
- **Solution**: Check YAML syntax, ensure file exists

### "package not found"
- **Cause**: Incorrect package path in configuration
- **Solution**: Verify package path matches Go module structure

### "duplicate service path"
- **Cause**: Conflicting service paths in configuration
- **Solution**: Use unique paths for each service

### "cannot create output file"
- **Cause**: Output directory doesn't exist or permission issues
- **Solution**: Create output directory, check permissions

### "import not found" in generated code
- **Cause**: Module path issues or missing dependencies
- **Solution**: Run `go mod tidy`, check module paths

### "404 Not Found" for service endpoints
- **Cause**: Incorrect HTTP handler registration
- **Solution**: Match service paths in configuration with handler registration

## Best Practices for AI Assistants

### 1. Always Validate Configuration
```yaml
# Use schema validation
# yaml-language-server: $schema=gotsrpc.schema.json
```

### 2. Test Configuration Early
```bash
# Validate before implementing
gotsrpc -debug gotsrpc.yml
```

### 3. Use Consistent Naming
```yaml
# Consistent service and path naming
targets:
  api:
    services:
      /api: APIService      # Clear, descriptive names
      /auth: AuthService    # Consistent patterns
```

### 4. Include Error Handling
```go
// Always handle errors in Context function
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    if err := h.authenticate(r); err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
}
```

### 5. Document Your Configuration
```yaml
# Add comments to explain complex configurations
targets:
  api:
    services:
      /api: APIService      # Main API service
      /auth: AuthService    # Authentication service
    package: github.com/your-org/project/service
    out: ./client/src/client.ts
    tsrpc:
      - APIService
      - AuthService
```

This troubleshooting guide should help resolve most common issues when implementing gotsrpc services, especially for AI assistants working with the framework.
