# gotsrpc Context Function Pattern

The Context function is a special pattern in gotsrpc that enables centralized authentication and request preprocessing. This document explains how to implement and use this pattern effectively.

## Overview

The Context function is automatically detected by gotsrpc when a service method has the signature:
```go
Context(w http.ResponseWriter, r *http.Request)
```

When present, gotsrpc will:
1. Call the Context function **first** for every request to that service
2. Only proceed to the actual method if Context doesn't write an error response
3. Allow the Context function to set up request context (authentication, user info, etc.)

## Basic Pattern

### Service Interface
```go
type Service interface {
    // Special Context function - called first for all requests
    Context(w http.ResponseWriter, r *http.Request)
    
    // Regular service methods
    GetData() (Data, error)
    UpdateData(data Data) error
}
```

### Implementation
```go
type ServiceHandler struct {
    currentUser *User
    // other fields
}

func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // 1. Extract authentication information
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Authorization required", http.StatusUnauthorized)
        return
    }
    
    // 2. Validate authentication
    token := strings.TrimPrefix(authHeader, "Bearer ")
    user, err := h.validateToken(token)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    // 3. Store user context for other methods
    h.currentUser = &user
    
    // 4. Set response headers (optional)
    w.Header().Set("X-User-ID", user.ID)
    
    // 5. Don't write response body - let other methods handle it
}

func (h *ServiceHandler) GetData() (Data, error) {
    // User is guaranteed to be authenticated here
    if h.currentUser == nil {
        return Data{}, errors.New("user not authenticated")
    }
    
    // Use h.currentUser for personalized data
    return h.getUserData(h.currentUser.ID), nil
}
```

## Authentication Patterns

### JWT Token Authentication
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
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

### API Key Authentication
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
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

### Session-based Authentication
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
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

## Advanced Patterns

### Role-based Access Control
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Authenticate user
    user, err := h.authenticateRequest(r)
    if err != nil {
        http.Error(w, "Authentication failed", http.StatusUnauthorized)
        return
    }
    
    // Check if user has required role for this endpoint
    requiredRole := h.getRequiredRole(r.URL.Path)
    if !h.hasRole(user, requiredRole) {
        http.Error(w, "Insufficient permissions", http.StatusForbidden)
        return
    }
    
    h.currentUser = &user
}

func (h *ServiceHandler) getRequiredRole(path string) string {
    switch {
    case strings.HasPrefix(path, "/admin"):
        return "admin"
    case strings.HasPrefix(path, "/user"):
        return "user"
    default:
        return "guest"
    }
}
```

### Rate Limiting
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    clientIP := h.getClientIP(r)
    
    // Check rate limit
    if !h.rateLimiter.Allow(clientIP) {
        http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
        return
    }
    
    // Continue with authentication
    user, err := h.authenticateRequest(r)
    if err != nil {
        http.Error(w, "Authentication failed", http.StatusUnauthorized)
        return
    }
    
    h.currentUser = &user
}
```

### Request Logging and Metrics
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Log request
    h.logger.Info("Request received",
        "method", r.Method,
        "path", r.URL.Path,
        "user_agent", r.UserAgent(),
        "ip", h.getClientIP(r),
    )
    
    // Record metrics
    h.metrics.IncrementRequestCounter(r.URL.Path)
    
    // Authenticate
    user, err := h.authenticateRequest(r)
    if err != nil {
        h.metrics.IncrementAuthFailureCounter()
        http.Error(w, "Authentication failed", http.StatusUnauthorized)
        return
    }
    
    h.currentUser = &user
    h.metrics.IncrementAuthSuccessCounter()
}
```

## Error Handling

### Custom Error Responses
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    user, err := h.authenticateRequest(r)
    if err != nil {
        h.writeErrorResponse(w, err)
        return
    }
    
    h.currentUser = &user
}

func (h *ServiceHandler) writeErrorResponse(w http.ResponseWriter, err error) {
    w.Header().Set("Content-Type", "application/json")
    
    switch {
    case errors.Is(err, ErrInvalidToken):
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(ErrorResponse{
            Code:    "INVALID_TOKEN",
            Message: "The provided token is invalid",
        })
    case errors.Is(err, ErrExpiredToken):
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(ErrorResponse{
            Code:    "EXPIRED_TOKEN",
            Message: "The provided token has expired",
        })
    default:
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(ErrorResponse{
            Code:    "INTERNAL_ERROR",
            Message: "An internal error occurred",
        })
    }
}
```

## Testing Context Functions

### Unit Testing
```go
func TestContextFunction(t *testing.T) {
    handler := &ServiceHandler{
        // Initialize with test dependencies
    }
    
    t.Run("ValidToken", func(t *testing.T) {
        req := httptest.NewRequest("POST", "/test", nil)
        req.Header.Set("Authorization", "Bearer valid-token")
        w := httptest.NewRecorder()
        
        handler.Context(w, req)
        
        assert.Equal(t, http.StatusOK, w.Code)
        assert.NotNil(t, handler.currentUser)
    })
    
    t.Run("InvalidToken", func(t *testing.T) {
        req := httptest.NewRequest("POST", "/test", nil)
        req.Header.Set("Authorization", "Bearer invalid-token")
        w := httptest.NewRecorder()
        
        handler.Context(w, req)
        
        assert.Equal(t, http.StatusUnauthorized, w.Code)
        assert.Nil(t, handler.currentUser)
    })
}
```

### Integration Testing
```go
func TestContextWithGeneratedProxy(t *testing.T) {
    handler := &ServiceHandler{}
    proxy := service.NewDefaultServiceGoTSRPCProxy(handler)
    
    // Test with generated proxy
    req := httptest.NewRequest("POST", "/Service/GetData", nil)
    req.Header.Set("Authorization", "Bearer test-token")
    w := httptest.NewRecorder()
    
    proxy.ServeHTTP(w, req)
    
    // Context should be called automatically
    assert.Equal(t, http.StatusOK, w.Code)
}
```

## Best Practices

### 1. Keep Context Functions Simple
```go
// ✅ Good - focused on authentication
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    user, err := h.authenticateRequest(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    h.currentUser = &user
}

// ❌ Avoid - too much business logic
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Authentication
    user, err := h.authenticateRequest(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Business logic (should be in service methods)
    if user.Role == "admin" {
        h.processAdminData()
    }
    
    h.currentUser = &user
}
```

### 2. Use Dependency Injection
```go
type ServiceHandler struct {
    authService    AuthService
    rateLimiter    RateLimiter
    logger         Logger
    currentUser    *User
}

func NewServiceHandler(authService AuthService, rateLimiter RateLimiter, logger Logger) *ServiceHandler {
    return &ServiceHandler{
        authService: authService,
        rateLimiter: rateLimiter,
        logger:      logger,
    }
}
```

### 3. Handle Context State Properly
```go
// ✅ Good - thread-safe approach
type ServiceHandler struct {
    mu           sync.RWMutex
    currentUser   *User
}

func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    user, err := h.authenticateRequest(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    h.mu.Lock()
    h.currentUser = &user
    h.mu.Unlock()
}

func (h *ServiceHandler) GetData() (Data, error) {
    h.mu.RLock()
    user := h.currentUser
    h.mu.RUnlock()
    
    if user == nil {
        return Data{}, errors.New("not authenticated")
    }
    
    return h.getUserData(user.ID), nil
}
```

### 4. Provide Clear Error Messages
```go
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Authorization header required", http.StatusUnauthorized)
        return
    }
    
    if !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Invalid authorization format. Expected 'Bearer <token>'", http.StatusUnauthorized)
        return
    }
    
    // Continue with validation...
}
```

## Common Pitfalls

### 1. Writing Response Body in Context
```go
// ❌ Wrong - don't write response body in Context
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Authentication logic...
    
    // Don't do this!
    fmt.Fprintf(w, "Hello %s", user.Username)
}

// ✅ Correct - only write error responses
func (h *ServiceHandler) Context(w http.ResponseWriter, r *http.Request) {
    // Authentication logic...
    
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Don't write success response - let service methods handle it
}
```

### 2. Not Handling Concurrent Requests
```go
// ❌ Wrong - shared state without synchronization
type ServiceHandler struct {
    currentUser *User  // Not thread-safe
}

// ✅ Correct - use request context or thread-safe storage
type ServiceHandler struct {
    // Use request context instead of instance variables
}
```

### 3. Forgetting to Check Authentication in Service Methods
```go
// ❌ Wrong - assuming Context always succeeds
func (h *ServiceHandler) GetData() (Data, error) {
    // This might be nil if Context failed
    return h.getUserData(h.currentUser.ID), nil
}

// ✅ Correct - always check authentication
func (h *ServiceHandler) GetData() (Data, error) {
    if h.currentUser == nil {
        return Data{}, errors.New("not authenticated")
    }
    return h.getUserData(h.currentUser.ID), nil
}
```

## Generated Code Integration

When using the Context function pattern, gotsrpc generates code that:

1. **Calls Context first**: Every request to the service calls Context before the actual method
2. **Handles errors**: If Context writes an error response, the method is not called
3. **Preserves state**: The handler instance maintains state between Context and method calls

### Generated Client Usage
```typescript
// The generated client automatically handles the Context call
const response = await serviceClient.getData();
// This internally calls:
// 1. POST /service/Context (with Authorization header)
// 2. POST /service/GetData (if Context succeeds)
```

This pattern enables clean separation of concerns while maintaining type safety and automatic code generation benefits.
