package gotsrpc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaultHttpClientFactory(t *testing.T) {
	newFactory := func() *http.Client {
		return nil
	}

	SetDefaultHttpClientFactory(newFactory)
	assert.Nil(t, defaultHttpFactory())
}
