package gotsrpc

import (
	"net/http"
)

type MonitorFn func(w http.ResponseWriter, r *http.Request, args, rets []interface{}, stats *CallStats)

var Monitor MonitorFn = NoopMonitor

func NoopMonitor(w http.ResponseWriter, r *http.Request, args, rets []interface{}, stats *CallStats) {
	// doing nothing
}
