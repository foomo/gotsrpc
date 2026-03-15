package gotsrpc

import (
	"context"
	"net/http"
)

type contextKey string

const contextStatsKey contextKey = "gotsrpcStats"

func RequestWithStatsContext(r *http.Request) *http.Request {
	stats := &CallStats{}
	return r.WithContext(context.WithValue(r.Context(), contextStatsKey, stats))
}

func GetStatsForRequest(r *http.Request) (*CallStats, bool) {
	if value, ok := r.Context().Value(contextStatsKey).(*CallStats); ok && value != nil {
		return value, true
	}
	return nil, false
}

func ClearStats(r *http.Request) {
	*r = *r.WithContext(context.WithValue(r.Context(), contextStatsKey, nil))
}
