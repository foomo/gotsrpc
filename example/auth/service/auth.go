package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// User represents an authenticated user
type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	LoginAt  time.Time `json:"loginAt"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// AuthService handles authentication operations
type AuthService interface {
	Login(req LoginRequest) (LoginResponse, error)
	Logout(token string) error
	ValidateToken(token string) (User, error)
}

// In-memory token storage (in production, use Redis or database)
type tokenStore struct {
	mu     sync.RWMutex
	tokens map[string]User
}

func newTokenStore() *tokenStore {
	return &tokenStore{
		tokens: make(map[string]User),
	}
}

func (ts *tokenStore) store(token string, user User) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.tokens[token] = user
}

func (ts *tokenStore) get(token string) (User, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	user, exists := ts.tokens[token]
	return user, exists
}

func (ts *tokenStore) delete(token string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.tokens, token)
}

// AuthHandler implements the AuthService
type AuthHandler struct {
	tokens *tokenStore
	// In production, you'd have a proper user database
	users map[string]string // username -> password
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		tokens: newTokenStore(),
		users: map[string]string{
			"alice": "password123",
			"bob":   "secret456",
			"admin": "admin789",
		},
	}
}

// generateToken creates a secure random token
func (h *AuthHandler) generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Login authenticates a user and returns a token
func (h *AuthHandler) Login(req LoginRequest) (LoginResponse, error) {
	// Validate credentials (in production, use proper password hashing)
	password, exists := h.users[req.Username]
	if !exists || password != req.Password {
		return LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	// Generate token
	token, err := h.generateToken()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create user
	user := User{
		ID:       fmt.Sprintf("user_%s", req.Username),
		Username: req.Username,
		Email:    fmt.Sprintf("%s@example.com", req.Username),
		LoginAt:  time.Now(),
	}

	// Store token
	h.tokens.store(token, user)

	return LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

// Logout invalidates a token
func (h *AuthHandler) Logout(token string) error {
	h.tokens.delete(token)
	return nil
}

// ValidateToken checks if a token is valid and returns the user
func (h *AuthHandler) ValidateToken(token string) (User, error) {
	user, exists := h.tokens.get(token)
	if !exists {
		return User{}, fmt.Errorf("invalid token")
	}
	return user, nil
}

// ContextKey type for context values
type ContextKey string

const UserContextKey ContextKey = "user"

// GetUserFromContext extracts user from request context
func GetUserFromContext(r *http.Request) (User, bool) {
	user, ok := r.Context().Value(UserContextKey).(User)
	return user, ok
}

// SetUserInContext adds user to request context
func SetUserInContext(r *http.Request, user User) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), UserContextKey, user))
}
