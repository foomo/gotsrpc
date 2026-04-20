package gotsrpc_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2"
	"github.com/stretchr/testify/assert"
)

func TestInstrumentedService(t *testing.T) {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		if s, ok := gotsrpc.GetStatsForRequest(r); ok && s != nil {
			s.Func = "func"
			s.Package = "package"
			s.Service = "service"
		}
	}

	t.Run("stats", func(t *testing.T) {
		count := 0
		handler := gotsrpc.InstrumentedService(middleware, func(s *gotsrpc.CallStats) {
			assert.Equal(t, "func", s.Func)
			assert.Equal(t, "package", s.Package)
			assert.Equal(t, "service", s.Service)
			assert.NotNil(t, s)

			count++
		})

		rsp := httptest.NewRecorder()
		req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/test", nil)

		handler(rsp, req)

		assert.Equal(t, 1, count)
	})
}
