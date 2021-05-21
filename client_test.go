package gotsrpc

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newRequest(t *testing.T) {
	t.Run("custom headers", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("test", "test")

		request, err := newRequest(context.Background(), "/test", "text/html", nil, headers)
		assert.NoError(t, err)
		assert.Equal(t, "test", request.Header.Get("test"))
	})
	t.Run("default", func(t *testing.T) {
		request, err := newRequest(context.Background(), "/test", "text/html", nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, "/test", request.URL.Path)
		assert.Equal(t, "text/html", request.Header.Get("Accept"))
		assert.Equal(t, "text/html", request.Header.Get("Content-Type"))
	})
}
