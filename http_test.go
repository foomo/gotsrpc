package gotsrpc

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSetDefaultHttpClientFactory(t *testing.T) {
	newFactory := func() *http.Client {
		return nil
	}

	SetDefaultHttpClientFactory(newFactory)
	assert.Nil(t, defaultHttpFactory())
}
