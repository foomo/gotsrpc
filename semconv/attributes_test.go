package semconv_test

import (
	"errors"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v3/semconv"
	"github.com/stretchr/testify/assert"
)

func TestRPCMethod(t *testing.T) {
	t.Parallel()

	kv := semconv.RPCMethod("UserService", "GetUser")
	assert.Equal(t, "rpc.method", string(kv.Key))
	assert.Equal(t, "UserService/GetUser", kv.Value.AsString())
}

func TestErrorType(t *testing.T) {
	t.Parallel()

	kv := semconv.ErrorType(errors.New("boom"))
	assert.Equal(t, "error.type", string(kv.Key))
	assert.NotEmpty(t, kv.Value.AsString())
}

func TestErrorTypeString(t *testing.T) {
	t.Parallel()

	kv := semconv.ErrorTypeString("Not Found")
	assert.Equal(t, "error.type", string(kv.Key))
	assert.Equal(t, "Not Found", kv.Value.AsString())
}

func TestPackageAndEncoding(t *testing.T) {
	t.Parallel()

	pkg := semconv.Package("github.com/foo/bar")
	assert.Equal(t, "gotsrpc.package", string(pkg.Key))
	assert.Equal(t, "github.com/foo/bar", pkg.Value.AsString())

	enc := semconv.Encoding("json")
	assert.Equal(t, "gotsrpc.encoding", string(enc.Key))
	assert.Equal(t, "json", enc.Value.AsString())
}

func TestGoTSRPCErrorCode(t *testing.T) {
	t.Parallel()

	kv := semconv.GoTSRPCErrorCode(500)
	assert.Equal(t, "gotsrpc.error.code", string(kv.Key))
	assert.Equal(t, int64(500), kv.Value.AsInt64())
}

func TestDurationAttrs(t *testing.T) {
	t.Parallel()

	d := 1500 * time.Microsecond

	exec := semconv.ExecutionMicros(d)
	assert.Equal(t, "gotsrpc.execution_us", string(exec.Key))
	assert.Equal(t, int64(1500), exec.Value.AsInt64())

	marshal := semconv.MarshalMicros(d)
	assert.Equal(t, "gotsrpc.marshal_us", string(marshal.Key))
	assert.Equal(t, int64(1500), marshal.Value.AsInt64())

	unmarshal := semconv.UnmarshalMicros(d)
	assert.Equal(t, "gotsrpc.unmarshal_us", string(unmarshal.Key))
	assert.Equal(t, int64(1500), unmarshal.Value.AsInt64())
}

func TestSizeAttrs(t *testing.T) {
	t.Parallel()

	req := semconv.RequestSizeAttr(128)
	assert.Equal(t, "gotsrpc.request_size", string(req.Key))
	assert.Equal(t, int64(128), req.Value.AsInt64())

	resp := semconv.ResponseSizeAttr(256)
	assert.Equal(t, "gotsrpc.response_size", string(resp.Key))
	assert.Equal(t, int64(256), resp.Value.AsInt64())
}

func TestEncodingFromContentType(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"application/json; charset=utf-8":    "json",
		"application/msgpack; charset=utf-8": "msgpack",
		"text/plain":                         "unknown",
		"":                                   "unknown",
	}
	for contentType, want := range tests {
		assert.Equal(t, want, semconv.EncodingFromContentType(contentType), contentType)
	}
}
