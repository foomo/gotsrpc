# gotsrpc Configuration Guide

This guide provides comprehensive documentation for configuring gotsrpc to help AI assistants and developers implement client/server code effectively.

## Table of Contents
1. [Basic Configuration](#basic-configuration)
2. [Service Configuration](#service-configuration)
3. [Authentication Patterns](#authentication-patterns)
4. [Common Patterns](#common-patterns)
5. [Troubleshooting](#troubleshooting)
6. [AI Assistant Guidelines](#ai-assistant-guidelines)

## Basic Configuration

### RPC Types: TSRPC vs GORPC

gotsrpc supports two types of RPC communication:

**TSRPC (TypeScript RPC):**
- **Purpose**: Go-to-TypeScript communication
- **Use case**: Web frontends, mobile apps, external APIs
- **Generated code**: TypeScript client libraries
- **Configuration**: `tsrpc` array in targets

**GORPC (Go RPC):**
- **Purpose**: Go-to-Go communication
- **Use case**: Internal microservices, backend-to-backend communication
- **Generated code**: Go client libraries
- **Configuration**: `gorpc` array in targets

**Mixed RPC:**
- Same service can serve both TypeScript and Go clients
- Configure both `tsrpc` and `gorpc` arrays
- Useful for services that need both web and internal access

**Note**: Examples in this guide may show only `tsrpc` or only `gorpc` for clarity, but you can configure both as needed for your use case.

**Important**: In all cases, gotsrpc generates **Go server code** for the backend. The `tsrpc` and `gorpc` arrays only determine which **client libraries** are generated (TypeScript clients, Go clients, or both).

### Configuration Examples

**TypeScript-only service:**
```yaml
# yaml-language-server: $schema=gotsrpc.schema.json
module:
  name: github.com/your-org/your-project
  path: ./

targets:
  web:
    services:
      /api: Service
    package: github.com/your-org/your-project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - Service

mappings:
  github.com/your-org/your-project/service:
    out: ./client/src/service-vo.ts
```

**Go-only service:**
```yaml
# yaml-language-server: $schema=gotsrpc.schema.json
module:
  name: github.com/your-org/your-project
  path: ./

targets:
  internal:
    services:
      /internal: Service
    package: github.com/your-org/your-project/service
    out: ./internal/internal-client.go
    gorpc:
      - Service
```

**Mixed service (both TypeScript and Go clients):**
```yaml
# yaml-language-server: $schema=gotsrpc.schema.json
module:
  name: github.com/your-org/your-project
  path: ./

targets:
  admin:
    services:
      /admin: Service
    package: github.com/your-org/your-project/service
    out: ./admin/admin-client.ts
    tsrpc:
      - Service    # TypeScript client
    gorpc:
      - Service    # Go client

mappings:
  github.com/your-org/your-project/service:
    out: ./client/src/service-vo.ts
```

### Configuration Fields Explained

#### Module Section
- `name`: Your Go module name (must match go.mod)
- `path`: Relative path to your Go module root (usually `./`)

#### Targets Section
Each target represents a client/server pair:

- `services`: Maps URL paths to service names
  - Key: URL path (e.g., `/api`, `/auth`)
  - Value: Service interface name (e.g., `Service`, `AuthService`)
- `package`: Full Go package path containing your services
- `out`: TypeScript client output file
- `tsrpc`: List of services to generate TypeScript RPC for
- `gorpc`: List of services to generate Go RPC for (optional)

#### Mappings Section
Maps Go packages to TypeScript output files for value objects (VOs).

## Service Configuration

### Single Service Example
```yaml
targets:
  basic:
    services:
      /service: Service
    package: github.com/your-org/project/service
    out: ./client/src/service-client.ts
    tsrpc:
      - Service
```

### Multiple Services Example
```yaml
targets:
  api:
    services:
      /auth: AuthService
      /users: UserService
      /posts: PostService
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - AuthService
      - UserService
      - PostService
```

### Mixed Go/TypeScript RPC
```yaml
targets:
  # TypeScript RPC services (for web clients)
  web:
    services:
      /api: APIService
      /auth: AuthService
    package: github.com/your-org/project/services
    out: ./client/src/web-client.ts
    tsrpc:
      - APIService
      - AuthService

  # Go RPC services (for internal microservice communication)
  internal:
    services:
      /internal: InternalService
      /admin: AdminService
    package: github.com/your-org/project/services
    out: ./internal/internal-client.go
    gorpc:
      - InternalService
      - AdminService
```

## Authentication Patterns

### Context Function Pattern
The `Context` function is a special gotsrpc pattern for authentication:

```go
// Service interface with Context function
type AuthService interface {
    // Special function: gotsrpc calls this first for all requests
    Context(w http.ResponseWriter, r *http.Request)
    
    // Regular service methods
    Login(req LoginRequest) (LoginResponse, error)
    Logout(token string) error
}

// Implementation
func (h *AuthHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Extract and validate Authorization header
    authHeader := r.Header.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    token := strings.TrimPrefix(authHeader, "Bearer ")
    user, err := h.validateToken(token)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    // Store user context for other methods
    h.currentUser = &user
}
```

### Configuration for Authentication
```yaml
targets:
  auth:
    services:
      /auth: AuthService
      /api: APIService  # Requires authentication
    package: github.com/your-org/project/service
    out: ./client/src/client.ts
    tsrpc:
      - AuthService
      - APIService
```

## Common Patterns

### 1. Microservices Architecture
```yaml
targets:
  auth:
    services:
      /auth: AuthService
    package: github.com/your-org/auth/service
    out: ./clients/auth/src/auth-client.ts
    tsrpc:
      - AuthService

  users:
    services:
      /users: UserService
    package: github.com/your-org/users/service
    out: ./clients/users/src/user-client.ts
    tsrpc:
      - UserService
```

### 2. Monolithic API
```yaml
targets:
  api:
    services:
      /auth: AuthService
      /users: UserService
      /posts: PostService
    package: github.com/your-org/api/service
    out: ./client/src/api-client.ts
    tsrpc:
      - AuthService
      - UserService
      - PostService
```

### 3. Separate Value Objects
```yaml
targets:
  api:
    services:
      /api: Service
    package: github.com/your-org/project/service
    out: ./client/src/service-client.ts
    tsrpc:
      - Service

mappings:
  github.com/your-org/project/service:
    out: ./client/src/service-vo.ts
  github.com/your-org/project/types:
    out: ./client/src/types-vo.ts
  time:
    out: ./client/src/vo-time.ts
```

### 4. Service Package Organization

There are two main approaches for organizing multiple services:

#### Approach A: Single Package with Unique Names

**Structure:**
```
service/
├── auth.go      # type AuthService interface { ... }
├── api.go       # type APIService interface { ... }
├── admin.go     # type AdminService interface { ... }
└── services.go  # Re-exports all services
```

**Configuration:**
```yaml
targets:
  web:
    services:
      /auth: AuthService
      /api: APIService
      /admin: AdminService
    package: github.com/your-org/project/service
    out: ./client/src/web-client.ts
    tsrpc:
      - AuthService
      - APIService
      - AdminService
```

**Pros:**
- All services in one place
- Simple configuration
- Easy to share types between services

**Cons:**
- Naming conflicts if not careful
- Can become large and unwieldy
- Harder to separate concerns

#### Approach B: Separate Packages with Simple Names

**Structure:**
```
services/
├── auth/          # Package: github.com/your-org/project/services/auth
│   └── service.go # type Service interface { ... }
├── api/           # Package: github.com/your-org/project/services/api
│   └── service.go # type Service interface { ... }
└── admin/         # Package: github.com/your-org/project/services/admin
    └── service.go # type Service interface { ... }
```

**File Naming Recommendation:**
- Use `service.go` instead of `auth.go`, `api.go`, etc.
- This emphasizes that the interface name is defined by the package
- Makes it clear that each package contains a `Service` interface
- Avoids confusion about which file contains which service

**Configuration:**
```yaml
targets:
  auth:
    services:
      /auth: Service
    package: github.com/your-org/project/services/auth
    out: ./client/src/auth-client.ts
    tsrpc:
      - Service

  api:
    services:
      /api: Service
    package: github.com/your-org/project/services/api
    out: ./client/src/api-client.ts
    tsrpc:
      - Service

  admin:
    services:
      /admin: Service
    package: github.com/your-org/project/services/admin
    out: ./client/src/admin-client.ts
    tsrpc:
      - Service
```

**Pros:**
- Clean separation of concerns
- No naming conflicts
- Easy to add new services
- Better for microservices architecture
- Each service can have its own dependencies

**Cons:**
- More complex configuration
- More files to manage
- Potential code duplication

#### Recommendation

**Use Approach A when:**
- Small to medium projects
- Services are closely related
- You want simple configuration
- Team prefers monolithic structure

**Use Approach B when:**
- Large projects with many services
- Services have different concerns
- Planning microservices architecture
- Team prefers modular structure
- Services have different dependencies

### 5. URL Paths and Endpoints

The URL paths in the `services` section define the HTTP endpoints:

```yaml
services:
  /api: Service      # Endpoint: /api
  /auth: Service     # Endpoint: /auth
  /admin: Service    # Endpoint: /admin
  /internal: Service # Endpoint: /internal
```

**Generated Server Code:**
```go
// Default endpoints (uses the path from services section)
proxy := NewDefaultServiceGoTSRPCProxy(service)  // Uses "/api"

// Custom endpoints
proxy := NewServiceGoTSRPCProxy(service, "/custom-api")  // Uses "/custom-api"
```

**Server Setup:**
```go
// Register handlers with custom endpoints
http.Handle("/api", NewServiceGoTSRPCProxy(apiService, "/api"))
http.Handle("/auth", NewServiceGoTSRPCProxy(authService, "/auth"))
http.Handle("/admin", NewServiceGoTSRPCProxy(adminService, "/admin"))
http.Handle("/internal", NewServiceGoTSRPCProxy(internalService, "/internal"))
```

### 6. Package Names and Generated Files

**Go Package Structure:**
```
services/
├── api/           # Package: github.com/your-org/project/services/api
│   └── api.go     # Contains: type Service interface { ... }
├── auth/          # Package: github.com/your-org/project/services/auth
│   └── auth.go    # Contains: type Service interface { ... }
└── admin/         # Package: github.com/your-org/project/services/admin
    └── admin.go   # Contains: type Service interface { ... }
```

**Configuration Mapping:**
```yaml
targets:
  api:
    services:
      /api: Service                    # Maps to api.Service interface
    package: github.com/your-org/project/services/api  # Go package path
    out: ./client/src/api-client.ts    # Generated TypeScript client
    tsrpc:
      - Service                        # Interface name in package

  auth:
    services:
      /auth: Service                   # Maps to auth.Service interface
    package: github.com/your-org/project/services/auth # Go package path
    out: ./client/src/auth-client.ts  # Generated TypeScript client
    tsrpc:
      - Service                        # Interface name in package
```

**Generated Files:**
- **TypeScript clients**: Generated to paths specified in `out` field
- **Go RPC proxies**: Generated in the service package directories
- **Value objects**: Generated to paths specified in `mappings` section

## Time Handling

### Time.Time Support
gotsrpc handles `time.Time` correctly with no unmarshaling issues:

**Go Side:**
- Marshals to RFC3339 format: `"2025-10-22T20:23:19.939773+02:00"`
- Unmarshaling works perfectly with timezone preservation
- Zero time values handled correctly: `"0001-01-01T00:00:00Z"`

**TypeScript Side:**
- `time.Time` maps to `number` (Unix timestamps in milliseconds)
- Easy conversion: `new Date(timestamp)`
- Proper handling of zero time values

**Configuration:**
```yaml
mappings:
  time:
    out: ./client/src/vo-time.ts
```

**Usage Example:**
```go
// Go struct
type User struct {
    CreatedAt time.Time `json:"created_at"`
    LastLogin time.Time `json:"last_login"`
}
```

```typescript
// Generated TypeScript interface
interface User {
    created_at: number;  // Unix timestamp
    last_login: number; // Unix timestamp
}

// Usage
const user = await client.getUser("user-123");
const createdAt = new Date(user.created_at);
const lastLogin = new Date(user.last_login);
```

## Troubleshooting

### Common Configuration Errors

#### 1. Module Path Issues
```yaml
# ❌ Wrong - path doesn't match go.mod location
module:
  name: github.com/your-org/project
  path: ../  # This should be ./

# ✅ Correct
module:
  name: github.com/your-org/project
  path: ./
```

#### 2. Service Path Conflicts
```yaml
# ❌ Wrong - conflicting paths
targets:
  api:
    services:
      /api: Service
      /api/auth: AuthService  # This conflicts with /api

# ✅ Correct - use different paths
targets:
  api:
    services:
      /api: Service
      /auth: AuthService
```

#### 3. Package Path Mismatch
```yaml
# ❌ Wrong - package doesn't exist
targets:
  api:
    package: github.com/your-org/nonexistent/service

# ✅ Correct - use actual package path
targets:
  api:
    package: github.com/your-org/project/service
```

### Debug Configuration
```bash
# Use debug flag to see detailed output
gotsrpc -debug gotsrpc.yml
```

## AI Assistant Guidelines

### When Implementing gotsrpc Services

1. **Always include the module section**:
   ```yaml
   module:
     name: github.com/your-org/your-project
     path: ./
   ```

2. **Use descriptive service paths**:
   ```yaml
   services:
     /auth: AuthService      # Authentication
     /api: APIService       # Main API
     /admin: AdminService   # Admin functions
   ```

3. **Include both tsrpc and gorpc when needed**:
   ```yaml
   tsrpc:
     - PublicService    # For TypeScript clients
   gorpc:
     - InternalService  # For Go-to-Go communication
   ```

4. **Map value objects properly**:
   ```yaml
   mappings:
     github.com/your-org/project/service:
       out: ./client/src/service-vo.ts
   ```

### Service Implementation Patterns

#### Basic Service
```go
type Service interface {
    // Regular methods
    GetData(id string) (Data, error)
    CreateData(data Data) (Data, error)
}
```

#### Service with Authentication
```go
type Service interface {
    // Special Context function for auth
    Context(w http.ResponseWriter, r *http.Request)
    
    // Protected methods
    GetUserData() (UserData, error)
    UpdateProfile(profile Profile) error
}
```

#### Service with HTTP Context
```go
type Service interface {
    // For HTTP-specific operations
    Context(w http.ResponseWriter, r *http.Request)
    
    // Regular methods
    HandleRequest(req Request) (Response, error)
}
```

### Generated Code Usage

#### Server-side (Go)
```go
// Create handlers
authHandler := service.NewAuthHandler()
apiHandler := service.NewAPIHandler()

// Create generated proxies
authProxy := service.NewDefaultAuthServiceGoTSRPCProxy(authHandler)
apiProxy := service.NewDefaultAPIServiceGoTSRPCProxy(apiHandler)

// Register with HTTP mux
mux.Handle("/auth/", authProxy)
mux.Handle("/api/", apiProxy)
```

#### Client-side (TypeScript)
```typescript
// Import generated clients
import { AuthServiceClient, APIServiceClient } from './client.ts';
import * as types from './vo.ts';

// Create transport function
const transport = <T>(endpoint: string) => async (method: string, data: any[] = []): Promise<T> => {
    // Implementation
};

// Create clients
const authClient = new AuthServiceClient(transport('/auth'));
const apiClient = new APIServiceClient(transport('/api'));
```

## Best Practices

1. **Use meaningful service names**: `AuthService`, `UserService`, not `Service1`, `Service2`
2. **Organize by domain**: Group related services together
3. **Use consistent naming**: Keep service names consistent across Go and TypeScript
4. **Document your services**: Add comments to service interfaces
5. **Test your configuration**: Use `gotsrpc -debug` to validate
6. **Version your generated code**: Don't commit generated files to version control

## Example: Complete Authentication Setup

### gotsrpc.yml
```yaml
module:
  name: github.com/your-org/auth-example
  path: ./

targets:
  auth:
    services:
      /auth: AuthService
      /api: APIService
    package: github.com/your-org/auth-example/service
    out: ./client/src/client.ts
    tsrpc:
      - AuthService
      - APIService

mappings:
  github.com/your-org/auth-example/service:
    out: ./client/src/vo.ts
```

### Service Implementation
```go
// auth.go
type AuthService interface {
    Login(req LoginRequest) (LoginResponse, error)
    Logout(token string) error
    ValidateToken(token string) (User, error)
}

// api.go
type APIService interface {
    Context(w http.ResponseWriter, r *http.Request)  // Authentication
    GetUserData() (UserData, error)
    UpdateProfile(profile Profile) error
}
```

This configuration will generate:
- TypeScript client with authentication support
- Go server proxies with Context function handling
- Proper value object mappings
- Full type safety across the stack
