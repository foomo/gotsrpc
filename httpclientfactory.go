package gotsrpc

import (
	"net"
	"net/http"
	"time"
)

// defaultHttpFactory creates the HTTP client of new clients. Unlike
// http.DefaultClient it sets timeouts; see
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779.
var defaultHttpFactory HttpClientFactory = func() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   45 * time.Second,
			KeepAlive: 45 * time.Second,
		}).DialContext,
		// TODO MaxIdleConnsPerHost: 32,
		DisableKeepAlives: true,

		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   2 * time.Minute, // Terminate Connection after 2 Minutes
	}
}

// HttpClientFactory returns the *http.Client that a new [Client] uses.
type HttpClientFactory func() *http.Client

// SetDefaultHttpClientFactory replaces the factory used by [NewClient] and by
// [NewClientWithHttpClient] with a nil client. It affects only clients created
// afterwards and is not safe for concurrent use; call it during program
// initialization.
func SetDefaultHttpClientFactory(factory HttpClientFactory) {
	defaultHttpFactory = factory
}
