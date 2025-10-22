package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/foomo/gotsrpc/v2/example/auth/service"
)

func main() {
	// Create authentication service
	authHandler := service.NewAuthHandler()

	// Create hello service with authentication dependency
	helloHandler := service.NewHelloHandler(authHandler)

	// Create gotsrpc proxies using generated code
	authProxy := service.NewDefaultAuthServiceGoTSRPCProxy(authHandler)
	helloProxy := service.NewDefaultHelloServiceGoTSRPCProxy(helloHandler)

	// Create a custom mux to handle routing properly
	mux := http.NewServeMux()

	// Register gotsrpc handlers first (more specific paths)
	mux.Handle("/auth/", authProxy)
	mux.Handle("/hello/", helloProxy)

	// Serve static files for the client (catch-all for non-API paths)
	// This will only match if the above handlers don't match
	mux.Handle("/", http.FileServer(http.Dir("./client/")))

	fmt.Println("Auth example server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /auth/Login - Login with username/password")
	fmt.Println("  POST /auth/Logout - Logout with token")
	fmt.Println("  POST /auth/ValidateToken - Validate token")
	fmt.Println("  POST /hello/Context - Authentication context (special gotsrpc function)")
	fmt.Println("  POST /hello/SayHello - Say hello (requires authentication)")
	fmt.Println("  POST /hello/GetUserInfo - Get user info (requires authentication)")
	fmt.Println("  GET / - Serve client application")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
