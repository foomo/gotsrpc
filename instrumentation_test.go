package gotsrpc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstrumentedService(t *testing.T) {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		if s, ok := GetStatsForRequest(r); ok && s != nil {
			s.Func = "func"
			s.Package = "package"
			s.Service = "service"
		}
	}

	t.Run("stats", func(t *testing.T) {
		count := 0
		handler := InstrumentedService(middleware, func(s *CallStats) {
			assert.Equal(t, "func", s.Func)
			assert.Equal(t, "package", s.Package)
			assert.Equal(t, "service", s.Service)
			assert.NotNil(t, s)
			count++
		})

		rsp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		handler(rsp, req)

		assert.Equal(t, 1, count)
	})
}
