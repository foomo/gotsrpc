package gotsrpc

import (
	"net"
	"net/http"
	"time"
)

// Default Client Factory
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

var defaultHttpFactory HttpClientFactory = func() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   45 * time.Second,
			KeepAlive: 45 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConnsPerHost: 20,
		MaxIdleConns:        20,
		IdleConnTimeout:     5 * time.Minute,

		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   2 * time.Minute, // Terminate Connection after 3 Minutes
	}
}

type HttpClientFactory func() *http.Client

func SetDefaultHttpClientFactory(factory HttpClientFactory) {
	defaultHttpFactory = factory
}
