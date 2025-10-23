# gotsrpc AI Assistant Implementation Guide

This guide is specifically designed to help AI assistants implement client/server code using gotsrpc effectively. It provides clear patterns, examples, and decision trees for common scenarios.

## Table of Contents
1. [Quick Start Patterns](#quick-start-patterns)
2. [Configuration Decision Tree](#configuration-decision-tree)
3. [Service Implementation Patterns](#service-implementation-patterns)
4. [Authentication Implementation](#authentication-implementation)
5. [Common Implementation Scenarios](#common-implementation-scenarios)
6. [Code Generation Workflow](#code-generation-workflow)
7. [Testing and Validation](#testing-and-validation)

## Quick Start Patterns

### Pattern 1: Basic API Service
**Use when**: Creating a simple API with one service

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

```go
// service/service.go
package service

type APIService interface {
    GetData(id string) (Data, error)
    CreateData(data Data) (Data, error)
    UpdateData(id string, data Data) error
    DeleteData(id string) error
}
```

### Pattern 2: Multi-Service API
**Use when**: Multiple related services (auth, users, posts, etc.)

```yaml
# gotsrpc.yml
module:
  name: github.com/your-org/project
  path: ./

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

mappings:
  github.com/your-org/project/service:
    out: ./client/src/api-vo.ts
```

### Pattern 3: Authentication-Required API
**Use when**: API requires authentication with Context function

```yaml
# gotsrpc.yml
module:
  name: github.com/your-org/project
  path: ./

targets:
  api:
    services:
      /auth: AuthService
      /api: APIService  # Requires authentication
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - AuthService
      - APIService
```

```go
// service/auth.go
type AuthService interface {
    Login(req LoginRequest) (LoginResponse, error)
    Logout(token string) error
    ValidateToken(token string) (User, error)
}

// service/api.go
type APIService interface {
    Context(w http.ResponseWriter, r *http.Request)  // Authentication
    GetUserData() (UserData, error)
    UpdateProfile(profile Profile) error
}
```

## Configuration Decision Tree

### Step 1: Determine Your Architecture

**Question**: How many services do you need?

- **Single Service** → Use Pattern 1 (Basic API Service)
- **Multiple Services** → Use Pattern 2 (Multi-Service API)
- **Services with Authentication** → Use Pattern 3 (Authentication-Required API)

### Step 2: Choose Service Paths

**Question**: What are your service endpoints?

```yaml
# Option A: Single service
services:
  /api: APIService

# Option B: Multiple services
services:
  /auth: AuthService
  /api: APIService
  /admin: AdminService

# Option C: Versioned API
services:
  /api/v1: V1Service
  /api/v2: V2Service
```

### Step 3: Choose RPC Types

**Question**: What clients will consume your services?

- **TypeScript/JavaScript clients only** → Use `tsrpc` only
- **Go services only** → Use `gorpc` only
- **Both** → Use separate targets for each type

```yaml
# TypeScript clients only
targets:
  web:
    services:
      /api: APIService
      /auth: AuthService
    package: github.com/your-org/project/services
    out: ./client/src/web-client.ts
    tsrpc:
      - APIService
      - AuthService

# Go services only
targets:
  internal:
    services:
      /internal: InternalService
      /admin: AdminService
    package: github.com/your-org/project/services
    out: ./internal/internal-client.go
    gorpc:
      - InternalService
      - AdminService

# Both (Mixed RPC)
targets:
  web:
    services:
      /api: APIService
      /auth: AuthService
    package: github.com/your-org/project/services
    out: ./client/src/web-client.ts
    tsrpc:
      - APIService
      - AuthService

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

### Step 4: Configure Value Objects

**Question**: Do you need separate value object files?

```yaml
# Single VO file
mappings:
  github.com/your-org/project/services:
    out: ./client/src/vo.ts
  time:
    out: ./client/src/vo-time.ts

# Multiple VO files
mappings:
  github.com/your-org/project/services:
    out: ./client/src/service-vo.ts
  github.com/your-org/project/types:
    out: ./client/src/types-vo.ts
  time:
    out: ./client/src/vo-time.ts
```

### Step 5: RPC Type Selection

**Question**: What type of clients do you need?

- **TypeScript clients only** → Use `tsrpc` array
- **Go clients only** → Use `gorpc` array  
- **Both TypeScript and Go clients** → Use both `tsrpc` and `gorpc` arrays

**Important for AI Assistants**: Always consider both TSRPC and GORPC possibilities:
- **TSRPC**: For web frontends, mobile apps, external APIs
- **GORPC**: For internal microservices, backend-to-backend communication
- **Mixed**: When the same service needs to serve both web and internal clients

**Remember**: gotsrpc always generates Go server code. The `tsrpc`/`gorpc` arrays only determine which client libraries are generated.

**Configuration Examples:**

**TypeScript-only:**
```yaml
targets:
  web:
    services:
      /api: Service
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - Service
```

**Go-only:**
```yaml
targets:
  internal:
    services:
      /internal: Service
    package: github.com/your-org/project/service
    out: ./internal/internal-client.go
    gorpc:
      - Service
```

**Mixed (both TypeScript and Go):**
```yaml
targets:
  admin:
    services:
      /admin: Service
    package: github.com/your-org/project/service
    out: ./admin/admin-client.ts
    tsrpc:
      - Service    # TypeScript client
    gorpc:
      - Service    # Go client
```

### Step 6: Service Organization

**Question**: How many services do you have?

- **Single service** → Use simple naming
- **Multiple services** → Choose between two approaches

## Service Organization Approaches

### Approach A: Single Package with Unique Names

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

### Approach B: Separate Packages with Simple Names

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

### Recommendation

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



## Understanding TSRPC vs GORPC for AI Assistants

### ✅ **Critical Knowledge for AI Assistants**

**Always consider both RPC types when implementing gotsrpc:**

1. **TSRPC (TypeScript RPC)**:
   - **Purpose**: Go backend ↔ TypeScript frontend communication
   - **Use cases**: Web applications, mobile apps, external APIs
   - **Generated**: TypeScript client libraries
   - **Configuration**: `tsrpc` array in targets

2. **GORPC (Go RPC)**:
   - **Purpose**: Go backend ↔ Go backend communication
   - **Use cases**: Internal microservices, backend-to-backend
   - **Generated**: Go client libraries
   - **Configuration**: `gorpc` array in targets

3. **Mixed RPC**:
   - **Purpose**: Same service serves both TypeScript and Go clients
   - **Use cases**: Admin services, APIs used by both web and internal services
   - **Generated**: Both TypeScript and Go client libraries
   - **Configuration**: Both `tsrpc` and `gorpc` arrays

**Key Point**: gotsrpc always generates Go server code. The `tsrpc`/`gorpc` arrays only determine which client libraries are generated.

**AI Assistant Decision Tree**:
- Does the service need web frontend access? → Include `tsrpc`
- Does the service need internal Go service access? → Include `gorpc`
- Does the service need both? → Include both `tsrpc` and `gorpc`

## Service Implementation Patterns

### Pattern A: Simple Service (No Authentication)
```go
type Service interface {
    GetData(id string) (Data, error)
    CreateData(data Data) (Data, error)
    UpdateData(id string, data Data) error
    DeleteData(id string) error
}

type ServiceHandler struct {
    // Dependencies
}

func (h *ServiceHandler) GetData(id string) (Data, error) {
    // Implementation
}
```

### Pattern B: Service with HTTP Context
```go
type Service interface {
    Context(w http.ResponseWriter, r *http.Request)  // Special function
    GetData() (Data, error)
    CreateData(data Data) (Data, error)
}

type ServiceHandler struct {
    currentUser *User
    // Other dependencies
}

func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Authentication logic
    user, err := h.authenticate(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    h.currentUser = &user
}

func (h *ServiceHandler) GetData() (Data, error) {
    if h.currentUser == nil {
        return Data{}, errors.New("not authenticated")
    }
    return h.getUserData(h.currentUser.ID), nil
}
```

### Pattern C: Mixed Service (Some methods need auth, others don't)
```go
type Service interface {
    // Public methods (no authentication required)
    GetPublicData() (PublicData, error)
    GetStatus() (Status, error)
    
    // Protected methods (authentication required)
    Context(w http.ResponseWriter, r *http.Request)
    GetUserData() (UserData, error)
    UpdateProfile(profile Profile) error
}
```

## Authentication Implementation

### Step 1: Choose Authentication Method

**JWT Tokens** (Recommended for APIs):
```go
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
        return
    }
    
    token := strings.TrimPrefix(authHeader, "Bearer ")
    claims, err := h.validateJWT(token)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    h.currentUser = &User{
        ID:       claims.UserID,
        Username: claims.Username,
        Role:     claims.Role,
    }
}
```

**API Keys**:
```go
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    apiKey := r.Header.Get("X-API-Key")
    if apiKey == "" {
        http.Error(w, "API key required", http.StatusUnauthorized)
        return
    }
    
    client, err := h.validateAPIKey(apiKey)
    if err != nil {
        http.Error(w, "Invalid API key", http.StatusUnauthorized)
        return
    }
    
    h.currentClient = &client
}
```

**Session-based**:
```go
func (h *Handler) Context(w http.ResponseWriter, r *http.Request) {
    sessionID, err := r.Cookie("session_id")
    if err != nil {
        http.Error(w, "Session required", http.StatusUnauthorized)
        return
    }
    
    session, err := h.validateSession(sessionID.Value)
    if err != nil {
        http.Error(w, "Invalid session", http.StatusUnauthorized)
        return
    }
    
    h.currentUser = &session.User
}
```

### Step 2: Implement Authentication Service

```go
type AuthService interface {
    Login(req LoginRequest) (LoginResponse, error)
    Logout(token string) error
    ValidateToken(token string) (User, error)
}

type AuthHandler struct {
    // Dependencies
}

func (h *AuthHandler) Login(req LoginRequest) (LoginResponse, error) {
    // Validate credentials
    user, err := h.validateCredentials(req.Username, req.Password)
    if err != nil {
        return LoginResponse{}, err
    }
    
    // Generate token
    token, err := h.generateToken(user)
    if err != nil {
        return LoginResponse{}, err
    }
    
    return LoginResponse{
        Token: token,
        User:  user,
    }, nil
}
```

## Common Implementation Scenarios

### Scenario 1: E-commerce API
```yaml
# gotsrpc.yml
module:
  name: github.com/your-org/ecommerce
  path: ./

targets:
  api:
    services:
      /auth: AuthService
      /users: UserService
      /products: ProductService
      /orders: OrderService
      /admin: AdminService
    package: github.com/your-org/ecommerce/service
    out: ./client/src/api-client.ts
    tsrpc:
      - AuthService
      - UserService
      - ProductService
      - OrderService
      - AdminService

mappings:
  github.com/your-org/ecommerce/service:
    out: ./client/src/api-vo.ts
```

### Scenario 2: Microservices Architecture
```yaml
# Auth service
targets:
  auth:
    services:
      /auth: AuthService
    package: github.com/your-org/auth/service
    out: ./clients/auth/src/auth-client.ts
    tsrpc:
      - AuthService

# User service
targets:
  users:
    services:
      /users: UserService
    package: github.com/your-org/users/service
    out: ./clients/users/src/user-client.ts
    tsrpc:
      - UserService
```

### Scenario 3: Internal + External APIs
```yaml
targets:
  external:
    services:
      /api: APIService
    package: github.com/your-org/project/service
    out: ./client/src/api-client.ts
    tsrpc:
      - APIService

  internal:
    services:
      /internal: InternalService
    package: github.com/your-org/project/service
    out: ./internal/client.go
    gorpc:
      - InternalService
```

## Code Generation Workflow

### Step 1: Create Configuration
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

### Step 2: Implement Service Interfaces
```go
// service/service.go
package service

type APIService interface {
    GetData(id string) (Data, error)
    CreateData(data Data) (Data, error)
}

type Data struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

### Step 3: Generate Code
```bash
# Generate gotsrpc code
gotsrpc gotsrpc.yml

# Or with debug output
gotsrpc -debug gotsrpc.yml
```

### Step 4: Implement Handlers
```go
// service/handler.go
type APIHandler struct {
    // Dependencies
}

func (h *APIHandler) GetData(id string) (Data, error) {
    // Implementation
}

func (h *APIHandler) CreateData(data Data) (Data, error) {
    // Implementation
}
```

### Step 5: Set Up HTTP Server
```go
// main.go
func main() {
    handler := service.NewAPIHandler()
    proxy := service.NewDefaultAPIServiceGoTSRPCProxy(handler)
    
    mux := http.NewServeMux()
    mux.Handle("/api/", proxy)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

### Step 6: Create TypeScript Client
```typescript
// client/src/app.ts
import { APIServiceClient } from './api-client.js';
import * as types from './api-vo.js';

const transport = <T>(endpoint: string) => async (method: string, data: any[] = []): Promise<T> => {
    const url = `http://localhost:8080${endpoint}/${encodeURIComponent(method)}`;
    
    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: data.length > 0 ? JSON.stringify(data) : undefined,
    });
    
    if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${await response.text()}`);
    }
    
    return response.text() ? JSON.parse(await response.text()) : {};
};

const client = new APIServiceClient(transport('/api'));

// Use the client
const data = await client.getData('123');
```

## Testing and Validation

### Step 1: Validate Configuration
```bash
# Check configuration syntax
gotsrpc -debug gotsrpc.yml

# Look for errors like:
# - "config load error" → Check YAML syntax and paths
# - "package not found" → Check package paths
# - "cannot create output file" → Check output directories
```

### Step 2: Test Generated Code
```bash
# Test Go compilation
go build ./service

# Test TypeScript compilation
cd client && npx tsc
```

### Step 3: Test HTTP Endpoints
```go
// Test with httptest
func TestAPIService(t *testing.T) {
    handler := service.NewAPIHandler()
    proxy := service.NewDefaultAPIServiceGoTSRPCProxy(handler)
    
    req := httptest.NewRequest("POST", "/api/GetData", strings.NewReader(`["123"]`))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    
    proxy.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusOK, w.Code)
}
```

### Step 4: Test Client Integration
```typescript
// Test TypeScript client
async function testClient() {
    try {
        const client = new APIServiceClient(transport('/api'));
        const data = await client.getData('123');
        console.log('Success:', data);
    } catch (error) {
        console.error('Error:', error);
    }
}
```

## Best Practices for AI Assistants

### 1. Always Start with Configuration
- Create `gotsrpc.yml` first
- Validate with `gotsrpc -debug gotsrpc.yml`
- Fix configuration errors before implementing code

### 2. Use Clear Naming Conventions
```yaml
# ✅ Good - descriptive names
services:
  /auth: AuthService
  /api: APIService
  /admin: AdminService

# ❌ Avoid - generic names
services:
  /service1: Service1
  /service2: Service2
```

### 3. Implement Authentication Early
- Add Context function to services that need authentication
- Test authentication flow before implementing business logic
- Use consistent authentication patterns across services

### 4. Test Incrementally
- Test configuration after each change
- Test generated code compilation
- Test HTTP endpoints individually
- Test client integration

### 5. Handle Errors Gracefully
```go
// Always check for authentication in service methods
func (h *Handler) GetData() (Data, error) {
    if h.currentUser == nil {
        return Data{}, errors.New("not authenticated")
    }
    // Implementation
}
```

### 6. Use Type Safety
```typescript
// Import generated types
import { APIServiceClient } from './api-client.js';
import * as types from './api-vo.js';

// Use typed requests
const request: types.CreateDataRequest = {
    name: 'Test Data',
    value: 123,
};

const response = await client.createData(request);
```

This guide provides AI assistants with the patterns and decision trees needed to implement gotsrpc services effectively. The key is to start with proper configuration, implement services incrementally, and test at each step.
