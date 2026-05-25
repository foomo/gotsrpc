package service

import (
	"fmt"
	"net/http"
	"strings"
)

// HelloRequest represents a hello request
type HelloRequest struct {
	Message string `json:"message"`
}

// HelloResponse represents a hello response
type HelloResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

// HelloService handles hello operations
type HelloService interface {
	Context(w http.ResponseWriter, r *http.Request)
	SayHello(req HelloRequest) (HelloResponse, error)
	GetUserInfo() (User, error)
}

// HelloHandler implements the HelloService
type HelloHandler struct {
	authService AuthService
	currentUser *User // Store current user for this request
}

func NewHelloHandler(authService AuthService) *HelloHandler {
	return &HelloHandler{
		authService: authService,
	}
}

// Context handles authentication for all hello service methods
func (h *HelloHandler) Context(w http.ResponseWriter, r *http.Request) {
	// Extract authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	// Validate Bearer token format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Invalid authorization format. Expected 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	// Extract token
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate token
	user, err := h.authService.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Store user for this request
	h.currentUser = &user

	// Set user info in response headers (optional)
	w.Header().Set("X-User-ID", user.ID)
	w.Header().Set("X-Username", user.Username)

	// Don't write response body - let other functions handle it
	// The generated code will check rw.Status() and only send a reply if needed
}

// SayHello returns a personalized greeting
func (h *HelloHandler) SayHello(req HelloRequest) (HelloResponse, error) {
	// Check if user is authenticated (set by Context method)
	if h.currentUser == nil {
		return HelloResponse{}, fmt.Errorf("user not authenticated")
	}

	// Create personalized message
	message := fmt.Sprintf("Hello %s! You said: %s", h.currentUser.Username, req.Message)

	return HelloResponse{
		Message: message,
		User:    *h.currentUser,
	}, nil
}

// GetUserInfo returns information about the current authenticated user
func (h *HelloHandler) GetUserInfo() (User, error) {
	// Check if user is authenticated (set by Context method)
	if h.currentUser == nil {
		return User{}, fmt.Errorf("user not authenticated")
	}

	return *h.currentUser, nil
}
