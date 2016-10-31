package gotsrpc

import "time"

type CallStats struct {
	Package, Service, Func string
	Execution              time.Duration
	Marshalling            time.Duration
	Unmarshalling          time.Duration
}
