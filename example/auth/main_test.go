package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/example/auth/service"
)

// TestGotsrpcContextFunction demonstrates the special Context function
// for authentication in gotsrpc services
func TestGotsrpcContextFunction(t *testing.T) {
	// Create services and generated gotsrpc proxies (same as main.go)
	authHandler := service.NewAuthHandler()
	helloHandler := service.NewHelloHandler(authHandler)
	helloProxy := service.NewDefaultHelloServiceGoTSRPCProxy(helloHandler)

	t.Run("ContextFunctionWithValidToken", func(t *testing.T) {
		// Get a valid token first
		loginReq := service.LoginRequest{
			Username: "alice",
			Password: "password123",
		}
		loginResp, err := authHandler.Login(loginReq)
		if err != nil {
			t.Fatalf("Login failed: %v", err)
		}

		// Test the special gotsrpc Context function for authentication
		req := httptest.NewRequest("POST", "/hello/Context", bytes.NewReader([]byte("[]")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+loginResp.Token)
		w := httptest.NewRecorder()

		// Call the generated gotsrpc Context function
		helloProxy.ServeHTTP(w, req)

		// The Context function should succeed with valid token
		if w.Code != http.StatusOK {
			t.Fatalf("gotsrpc Context function failed with status %d: %s", w.Code, w.Body.String())
		}
	})

	t.Run("ContextFunctionWithoutToken", func(t *testing.T) {
		// Test Context function without authentication (should fail)
		req := httptest.NewRequest("POST", "/hello/Context", bytes.NewReader([]byte("[]")))
		req.Header.Set("Content-Type", "application/json")
		// No Authorization header
		w := httptest.NewRecorder()

		helloProxy.ServeHTTP(w, req)

		// Context function should return unauthorized
		if w.Code == http.StatusOK {
			t.Fatal("gotsrpc Context function should fail without authentication")
		}
	})

	t.Run("ContextFunctionWithInvalidToken", func(t *testing.T) {
		// Test Context function with invalid token (should fail)
		req := httptest.NewRequest("POST", "/hello/Context", bytes.NewReader([]byte("[]")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer invalid-token")
		w := httptest.NewRecorder()

		helloProxy.ServeHTTP(w, req)

		// Context function should return unauthorized
		if w.Code == http.StatusOK {
			t.Fatal("gotsrpc Context function should fail with invalid token")
		}
	})
}

// TestGotsrpcGeneratedProxies demonstrates how to use generated gotsrpc proxy functions
func TestGotsrpcGeneratedProxies(t *testing.T) {
	// This test shows how to use the generated gotsrpc proxy functions
	authHandler := service.NewAuthHandler()
	helloHandler := service.NewHelloHandler(authHandler)

	// Create generated gotsrpc proxies
	authProxy := service.NewDefaultAuthServiceGoTSRPCProxy(authHandler)
	helloProxy := service.NewDefaultHelloServiceGoTSRPCProxy(helloHandler)

	t.Run("ProxyCreation", func(t *testing.T) {
		// Test that we can create the generated proxies
		if authProxy == nil {
			t.Fatal("Auth proxy should not be nil")
		}
		if helloProxy == nil {
			t.Fatal("Hello proxy should not be nil")
		}
	})

	t.Run("ProxyHTTPHandler", func(t *testing.T) {
		// Test that proxies implement http.Handler interface
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		// Should not panic when called
		authProxy.ServeHTTP(w, req)
		helloProxy.ServeHTTP(w, req)
	})
}
