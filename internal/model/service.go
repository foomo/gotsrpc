package model

// Service is a Go type whose exported methods are exposed over RPC.
type Service struct {
	Name    string
	Methods ServiceMethods
	// Endpoint is the HTTP path the service is served on.
	Endpoint string
	// IsInterface reports whether the service is declared as an interface.
	IsInterface bool
}
