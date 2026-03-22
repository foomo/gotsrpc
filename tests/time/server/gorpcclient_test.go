package server_test

import (
	"net"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v2/tests/time/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServiceGoRPCClient(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0") //nolint:noctx
	require.NoError(t, err)
	require.NoError(t, l.Close())

	s := server.NewServiceGoRPCProxy(l.Addr().String(), &server.Handler{}, nil)
	require.NoError(t, s.Start())
	defer s.Stop()

	c := server.NewServiceGoRPCClient(l.Addr().String(), nil)
	c.Start()
	defer c.Stop()

	t.Run("Time", func(t *testing.T) {
		now := time.Now()
		ret, clientErr := c.Time(now)
		require.NoError(t, clientErr)
		assert.Equal(t, now.UnixNano(), ret.UnixNano())
	})

	t.Run("TimeStruct", func(t *testing.T) {
		now := time.Now()
		v := server.TimeStruct{
			Time:        now,
			TimePtr:     &now,
			TimePtrOmit: &now,
		}
		ret, clientErr := c.TimeStruct(v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.Time.UnixNano(), ret.Time.UnixNano())
		assert.Equal(t, v.TimePtr.UnixNano(), ret.TimePtr.UnixNano())
		assert.Equal(t, v.TimePtrOmit.UnixNano(), ret.TimePtrOmit.UnixNano())
	})
}
