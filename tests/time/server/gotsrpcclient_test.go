package server_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v2/tests/time/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	now := time.Now()

	t.Run("Time", func(t *testing.T) {
		ret, clientErr := c.Time(t.Context(), now)
		require.NoError(t, clientErr)
		assert.Equal(t, now.UnixNano(), ret.UnixNano())
	})

	t.Run("TimeStruct", func(t *testing.T) {
		ret, clientErr := c.TimeStruct(t.Context(), server.TimeStruct{
			Time:        now,
			TimePtr:     &now,
			TimePtrOmit: &now,
		})
		require.NoError(t, clientErr)
		assert.Equal(t, now.UnixNano(), ret.Time.UnixNano())
		assert.Equal(t, now.UnixNano(), ret.TimePtr.UnixNano())
		assert.Equal(t, now.UnixNano(), ret.TimePtrOmit.UnixNano())
	})
}
