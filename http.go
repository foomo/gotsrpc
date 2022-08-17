package gotsrpc

import (
	"net"
	"net/http"
	"time"
)

// Default Client Factory
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

var defaultHttpFactory HttpClientFactory = func() *http.Client { //nolint:stylecheck
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   45 * time.Second,
			KeepAlive: 45 * time.Second,
		}).DialContext,
		DisableKeepAlives: true,

		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   2 * time.Minute, // Terminate Connection after 3 Minutes
	}
}

type HttpClientFactory func() *http.Client //nolint:stylecheck

func SetDefaultHttpClientFactory(factory HttpClientFactory) { //nolint:stylecheck
	defaultHttpFactory = factory
}
