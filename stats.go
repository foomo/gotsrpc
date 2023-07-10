package gotsrpc

import "time"

type CallStats struct {
	Package       string
	Service       string
	Func          string
	Execution     time.Duration
	Marshalling   time.Duration
	Unmarshalling time.Duration
	RequestSize   int
	ResponseSize  int
	ErrorCode     int
	ErrorMessage  string
}
