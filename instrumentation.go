package gotsrpc

import "net/http"

// Deprecated: will be removed
func InstrumentedService(middleware http.HandlerFunc, handleStats GoRPCCallStatsHandlerFun) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*r = *RequestWithStatsContext(r)
		middleware(w, r)
		if stats, ok := GetStatsForRequest(r); ok {
			handleStats(stats)
		}
	}
}
