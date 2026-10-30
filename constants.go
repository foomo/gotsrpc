package gotsrpc

const (
	// HeaderServiceToService is the request header that a Go [Client] sets to
	// mark a call as service-to-service.
	HeaderServiceToService = "X-Foomo-S2s"
)
