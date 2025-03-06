package gotsrpc

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newRequest(t *testing.T) {
	t.Run("custom headers", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Test", "test")

		request, err := newRequest(context.Background(), "/test", "text/html", nil, headers)
		require.NoError(t, err)
		assert.Equal(t, "test", request.Header.Get("Test"))
	})
	t.Run("default", func(t *testing.T) {
		request, err := newRequest(context.Background(), "/test", "text/html", nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "/test", request.URL.Path)
		assert.Equal(t, "text/html", request.Header.Get("Accept"))
		assert.Equal(t, "text/html", request.Header.Get("Content-Type"))
	})
}
