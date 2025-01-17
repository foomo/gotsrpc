package gotsrpc

import (
	"context"
	"encoding/json"
	"fmt"
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
	contentTypeHeaderMap := map[ClientEncoding]string{
		EncodingMsgpack: "application/msgpack; charset=utf-8",
		EncodingJson:    "application/json; charset=utf-8",
	}

	contentEncodingHeaderMap := map[Compressor]string{
		CompressorGZIP:   "gzip",
		CompressorSnappy: "snappy",
	}

	var testRequestData []interface{}
	data, err := os.ReadFile("testdata/request.json")
	require.NoError(t, err)

	err = json.Unmarshal(data, &testRequestData)
	require.NoError(t, err)

	testClient := func(
		t *testing.T,
		encoding ClientEncoding,
		compressor Compressor,
	) {
		requiredResponseMessage := "Fake Response Message"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var args []map[string]interface{}
			err := LoadArgs(&args, nil, r)
			require.NoError(t, err)

			require.Equal(t, contentTypeHeaderMap[encoding], r.Header.Get("Content-Type"))
			require.Equal(t, contentEncodingHeaderMap[compressor], r.Header.Get("Content-Encoding"))

			_ = Reply([]interface{}{requiredResponseMessage}, nil, r, w)
		}))
		defer server.Close()

		client := NewBufferedClient(
			WithCompressor(compressor),
			WithHTTPClient(server.Client()),
			WithClientEncoding(encoding),
		)

		require.NotNil(t, client)

		var actualResponseMessage string
		response := []interface{}{&actualResponseMessage}

		err := client.Call(context.Background(), server.URL, "/Example", "Example", testRequestData, response)
		require.NoError(t, err)
		require.Equal(t, requiredResponseMessage, actualResponseMessage)
	}

	for _, encoding := range []ClientEncoding{EncodingMsgpack, EncodingJson} {
		for _, compressor := range []Compressor{CompressorNone, CompressorGZIP, CompressorSnappy} {
			t.Run(fmt.Sprintf("%s/%s", encoding, compressor), func(t *testing.T) {
				testClient(t, encoding, compressor)
			})
		}
	}
}

func BenchmarkBufferedClient(b *testing.B) {
	var testRequestData []interface{}
	data, err := os.ReadFile("testdata/request.json")
	require.NoError(b, err)

	err = json.Unmarshal(data, &testRequestData)
	require.NoError(b, err)

	benchClient := func(b *testing.B, client Client) {
		b.Helper()
		b.ReportAllocs()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var args []map[string]interface{}
			err := LoadArgs(&args, nil, r)
			require.NoError(b, err)

			_ = Reply([]interface{}{"HI"}, nil, r, w)
		}))
		defer server.Close()

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
	runs := 1

	for name, compressor := range benchmarks {
		b.Run(name, func(b *testing.B) {
			for index := 0; index < runs; index++ {
				b.Run(fmt.Sprintf("%d", index), func(b *testing.B) { benchClient(b, NewBufferedClient(WithCompressor(compressor))) })
			}
		})
	}
}
