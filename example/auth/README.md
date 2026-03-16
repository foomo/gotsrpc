# gotsrpc Authentication Example

This example demonstrates how to implement client authentication using gotsrpc's special `Context` function handling. It shows a complete authentication flow with two services:

1. **Authentication Service** - Handles login/logout operations
2. **Hello Service** - Requires authentication and provides personalized responses

## Architecture

### Authentication Flow
1. Client logs in with username/password
2. Server validates credentials and returns a JWT-like token
3. Client stores token and includes it in Authorization header for subsequent requests
4. Hello service's `Context` method validates the token before processing requests
5. Authenticated user information is available to other service methods

### Key Features
- **Centralized Authentication**: The `Context` function handles auth for all Hello service methods
- **Token-based Security**: Uses secure random tokens for session management
- **User Context**: Authenticated user information is available throughout the request
- **Clean Separation**: Authentication logic is separate from business logic

## Services

### AuthService
- `Login(username, password)` - Authenticates user and returns token
- `Logout(token)` - Invalidates a token
- `ValidateToken(token)` - Validates token and returns user info

### HelloService
- `Context(w, r)` - **Special function** that handles authentication via Authorization header
- `SayHello(message)` - Returns personalized greeting (requires authentication)
- `GetUserInfo()` - Returns current user information (requires authentication)

## How the Context Function Works

The `Context` function is specially handled by gotsrpc:

1. **Detection**: gotsrpc automatically detects functions with `http.ResponseWriter` and `*http.Request` parameters
2. **Special Generation**: The generated code calls `Context` first, then other methods
3. **Authentication**: `Context` validates the Authorization header and sets up user context
4. **Error Handling**: If authentication fails, `Context` writes an error response and returns early
5. **Success**: If authentication succeeds, other methods can access the authenticated user

## Running the Example

### Prerequisites
- Go 1.19+
- Node.js (for TypeScript compilation)
- `gotsrpc` tool installed

### Quick Start
```bash
cd example/auth
make dev
```

This will:
1. Clean generated files
2. Generate gotsrpc server and client code
3. Compile TypeScript to JavaScript
4. Start the server on `http://localhost:8080`

### Manual Steps
If you prefer to run steps manually:

```bash
# Generate gotsrpc code
make generate

# Compile TypeScript client
make build-client

# Start the server
make run-server
```

### Available Makefile Commands
- `make generate` - Generate gotsrpc server and client code
- `make build-client` - Compile TypeScript to JavaScript
- `make run-server` - Start the Go server
- `make run` - Build client and run server
- `make dev` - Full workflow: generate + build + run
- `make clean` - Remove compiled JavaScript files

### Access the Client
Open `http://localhost:8080` in your browser to access the web interface.

### Test Accounts
- **Username**: `alice`, **Password**: `password123`
- **Username**: `bob`, **Password**: `secret456`
- **Username**: `admin`, **Password**: `admin789`

## API Endpoints

### Authentication
- `POST /auth/Login` - Login with username/password
- `POST /auth/Logout` - Logout with token
- `POST /auth/ValidateToken` - Validate token and get user info

### Hello Service (Requires Authentication)
- `POST /hello/Context` - Authentication context (special gotsrpc function)
- `POST /hello/SayHello` - Say hello with personalized message
- `POST /hello/GetUserInfo` - Get current user information

## Code Structure

```
example/auth/
├── service/
│   ├── auth.go              # Authentication service implementation
│   ├── hello.go             # Hello service with Context authentication
│   ├── gotsrpc_gen.go       # Generated gotsrpc server code
│   └── gotsrpcclient_gen.go # Generated gotsrpc client code
├── client/
│   ├── index.html           # Web interface
│   ├── src/
│   │   ├── app.ts           # TypeScript client code
│   │   ├── client_gen.ts     # Generated gotsrpc client
│   │   ├── vo_gen.ts         # Generated value objects
│   │   ├── vo-time_gen.ts    # Generated time value objects
│   │   └── types.d.ts        # TypeScript type definitions
│   ├── dist/
│   │   ├── app.js            # Compiled JavaScript
│   │   ├── client_gen.js     # Compiled generated client
│   │   ├── vo_gen.js         # Compiled value objects
│   │   └── vo-time_gen.js    # Compiled time objects
│   └── tsconfig.json         # TypeScript configuration
├── main.go                   # Server setup and HTTP handlers
├── main_test.go              # Tests for gotsrpc Context function
├── gotsrpc.yml               # gotsrpc configuration
├── Makefile                  # Build automation
└── README.md                 # This file
```

## Key Implementation Details

### Authentication in Context Function
```go
func (h *HelloHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Extract and validate Authorization header
    authHeader := r.Header.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
        return
    }
    
    // Validate token and store user context
    token := strings.TrimPrefix(authHeader, "Bearer ")
    user, err := h.authService.ValidateToken(token)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    // Store user for other methods to use
    h.currentUser = &user
}
```

### Client-Side Token Management
```typescript
// Include token in requests
if (requiresAuth && authToken) {
    options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${authToken}`,
    };
}
```

### Generated Code Usage
The example uses generated gotsrpc code for both server and client:

**Server-side (Go):**
```go
// Generated gotsrpc proxies handle HTTP routing and method calls
authProxy := service.NewDefaultAuthServiceGoTSRPCProxy(authHandler)
helloProxy := service.NewDefaultHelloServiceGoTSRPCProxy(helloHandler)

// Register with HTTP mux
mux.Handle("/auth/", authProxy)
mux.Handle("/hello/", helloProxy)
```

**Client-side (TypeScript):**
```typescript
// Import generated clients and types
import { AuthServiceClient, HelloServiceClient } from './client_gen.js';
import * as types from './vo_gen.js';

// Use generated clients
const authClient = new AuthServiceClient(transport);
const helloClient = new HelloServiceClient(transport);
```

## Benefits of This Pattern

1. **Single Point of Authentication**: All auth logic is centralized in the `Context` function
2. **No Code Duplication**: Other service methods don't need to implement authentication
3. **Flexible**: Can implement any authentication scheme (JWT, API keys, OAuth, etc.)
4. **Clean Separation**: Business logic is separate from authentication concerns
5. **Type Safety**: Full TypeScript support for client-side code
6. **Code Generation**: gotsrpc generates both server and client code automatically
7. **Build Automation**: Makefile provides clean build and development workflow

## Production Considerations

For production use, consider:

- **Token Storage**: Use Redis or database instead of in-memory storage
- **Password Hashing**: Use bcrypt or similar for password storage
- **Token Expiration**: Implement token expiration and refresh mechanisms
- **HTTPS**: Always use HTTPS in production
- **Rate Limiting**: Implement rate limiting for authentication endpoints
- **Logging**: Add comprehensive logging for security events
