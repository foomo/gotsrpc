package gotsrpc

import (
	"net/http"
)

type MonitorFn func(w http.ResponseWriter, r *http.Request, args, rets []any, stats *CallStats)

var Monitor MonitorFn = NoopMonitor

func NoopMonitor(w http.ResponseWriter, r *http.Request, args, rets []any, stats *CallStats) {
	// doing nothing
}
