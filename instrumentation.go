package gotsrpc

import "net/http"

func InstrumentedService(middleware http.HandlerFunc, handleStats GoRPCCallStatsHandlerFun) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*r = *RequestWithStatsContext(r)
		middleware(w, r)
		stats := GetStatsForRequest(r)
		if stats != nil {
			handleStats(stats)
		}
	}
}
