package gotsrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewBufferedClient(t *testing.T) {
	var testRequestData []interface{}
	data, err := os.ReadFile("testdata/request.json")
	require.NoError(t, err)

	err = json.Unmarshal(data, &testRequestData)
	require.NoError(t, err)
	t.Run("gzip", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//require.Equal(t, "application/msgpack; charset=utf-8", request.Header.Get("Content-Type"))
			require.Equal(t, "gzip", request.Header.Get("Content-Encoding"))
			data, _ := io.ReadAll(request.Body)
			fmt.Println(string(data))

			_, _ = writer.Write([]byte("[]"))
		}))
		defer server.Close()

		client := NewBufferedClient(
			WithCompressor(CompressorGZIP),
		)

		assert.NotNil(t, client)
		err := client.Call(context.Background(), server.URL, "/test", "test", testRequestData, nil)
		assert.NoError(t, err)
	})
	t.Run("snappy", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//require.Equal(t, "application/msgpack; charset=utf-8", request.Header.Get("Content-Type"))
			require.Equal(t, "snappy", request.Header.Get("Content-Encoding"))
			data, _ := io.ReadAll(request.Body)
			fmt.Println(string(data))

			_, _ = writer.Write([]byte("[]"))
		}))
		defer server.Close()

		client := NewBufferedClient(
			WithCompressor(CompressorSnappy),
		)

		assert.NotNil(t, client)
		err := client.Call(context.Background(), server.URL, "/test", "test", testRequestData, nil)
		assert.NoError(t, err)
	})
	t.Run("plain", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//require.Equal(t, "application/msgpack; charset=utf-8", request.Header.Get("Content-Type"))
			require.Empty(t, request.Header.Get("Content-Encoding"))
			data, _ := io.ReadAll(request.Body)
			fmt.Println(string(data))

			_, _ = writer.Write([]byte("[]"))
		}))
		defer server.Close()

		client := NewBufferedClient()

		assert.NotNil(t, client)
		err := client.Call(context.Background(), server.URL, "/test", "test", testRequestData, nil)
		assert.NoError(t, err)
	})
}

func BenchmarkBufferedClient(b *testing.B) {
	var testRequestData []interface{}
	data, err := os.ReadFile("testdata/request.json")
	require.NoError(b, err)

	err = json.Unmarshal(data, &testRequestData)
	require.NoError(b, err)

	benchClient := func(b *testing.B, client Client) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("[]"))
		}))
		defer server.Close()
		b.ReportAllocs()
		b.ResetTimer()
		if bc, ok := client.(*BufferedClient); ok {
			bc.client = server.Client()
		}

		for i := 0; i < b.N; i++ {
			err := client.Call(context.Background(), server.URL, "/test", "test", testRequestData, nil)
			require.NoError(b, err)
		}
	}
	benchmarks := map[string]Compressor{
		"none":   CompressorNone,
		"gzip":   CompressorGZIP,
		"snappy": CompressorSnappy,
	}
	runs := 5

	for name, compressor := range benchmarks {
		b.Run(name, func(b *testing.B) {
			for index := 0; index < runs; index++ {
				b.Run(fmt.Sprintf("%d", index), func(b *testing.B) { benchClient(b, NewBufferedClient(WithCompressor(compressor))) })
			}
		})
	}
}
