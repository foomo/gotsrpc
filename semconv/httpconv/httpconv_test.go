package httpconv_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v3/semconv"
	"github.com/foomo/gotsrpc/v3/semconv/httpconv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

type harness struct {
	recorder *tracetest.SpanRecorder
	reader   *sdkmetric.ManualReader
	opts     []semconv.Option
}

func newHarness() *harness {
	recorder := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))

	return &harness{
		recorder: recorder,
		reader:   reader,
		opts: []semconv.Option{
			semconv.WithTracerProvider(tp),
			semconv.WithMeterProvider(mp),
			semconv.WithPropagators(propagation.TraceContext{}),
		},
	}
}

func (h *harness) metricNames(t *testing.T) []string {
	t.Helper()

	var rm metricdata.ResourceMetrics
	require.NoError(t, h.reader.Collect(t.Context(), &rm))

	var names []string

	for _, sm := range rm.ScopeMetrics {
		for _, m := range sm.Metrics {
			names = append(names, m.Name)
		}
	}

	return names
}

func attrValue(attrs []attribute.KeyValue, key string) (attribute.Value, bool) {
	for _, a := range attrs {
		if string(a.Key) == key {
			return a.Value, true
		}
	}

	return attribute.Value{}, false
}

func TestServer_SuccessLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	srv := httpconv.NewServer("github.com/foo/bar", "UserService", h.opts...)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/user", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	_, call := srv.Start(nil, req, "GetUser")
	call.ExecutionStart()
	call.ExecutionEnd()
	call.RecordUnmarshal(2*time.Millisecond, 128)
	call.RecordMarshal(3*time.Millisecond, 256)
	call.SetStatus(http.StatusOK)
	call.End()

	spans := h.recorder.Ended()
	require.Len(t, spans, 1)
	span := spans[0]
	assert.Equal(t, "UserService/GetUser", span.Name())
	assert.Equal(t, codes.Unset, span.Status().Code)

	method, ok := attrValue(span.Attributes(), "rpc.method")
	require.True(t, ok)
	assert.Equal(t, "UserService/GetUser", method.AsString())

	enc, ok := attrValue(span.Attributes(), "gotsrpc.encoding")
	require.True(t, ok)
	assert.Equal(t, "json", enc.AsString())

	assert.Contains(t, h.metricNames(t), "rpc.server.call.duration")
}

func TestServer_ErrorLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	srv := httpconv.NewServer("github.com/foo/bar", "UserService", h.opts...)
	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/user", nil)

	_, call := srv.Start(nil, req, "GetUser")
	call.RecordError(errors.New("boom"), 500)
	call.End()

	spans := h.recorder.Ended()
	require.Len(t, spans, 1)
	span := spans[0]
	assert.Equal(t, codes.Error, span.Status().Code)

	_, ok := attrValue(span.Attributes(), "error.type")
	assert.True(t, ok)
	code, ok := attrValue(span.Attributes(), "gotsrpc.error.code")
	require.True(t, ok)
	assert.Equal(t, int64(500), code.AsInt64())
}

func TestServer_NotFoundAndSetError(t *testing.T) {
	t.Parallel()

	t.Run("not found", func(t *testing.T) {
		t.Parallel()

		h := newHarness()
		srv := httpconv.NewServer("pkg", "Svc", h.opts...)
		req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/x", nil)
		_, call := srv.Start(nil, req, "Missing")
		call.NotFound()
		call.End()

		span := h.recorder.Ended()[0]
		assert.Equal(t, codes.Error, span.Status().Code)
		et, ok := attrValue(span.Attributes(), "error.type")
		require.True(t, ok)
		assert.Equal(t, "method_not_found", et.AsString())
	})

	t.Run("transport error via SetError and SetStatus", func(t *testing.T) {
		t.Parallel()

		h := newHarness()
		srv := httpconv.NewServer("pkg", "Svc", h.opts...)
		req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/x", nil)
		_, call := srv.Start(nil, req, "M")
		call.SetError(errors.New("decode failed"))
		call.SetStatus(http.StatusInternalServerError) // no-op: errType already set
		call.End()

		span := h.recorder.Ended()[0]
		assert.Equal(t, codes.Error, span.Status().Code)
	})
}

func TestClient_SuccessLifecycleAndInject(t *testing.T) {
	t.Parallel()

	h := newHarness()

	client := httpconv.NewClient(append(h.opts, semconv.WithClientService("github.com/foo/bar", "UserService"))...)

	_, call := client.Start(context.Background(), "GetUser")
	call.RecordRequestSize(128)
	call.RecordResponseSize(256)

	header := http.Header{}
	call.Inject(header)
	call.SetStatus(http.StatusOK)
	call.End()

	// TraceContext propagator injects a valid traceparent for the active span.
	assert.NotEmpty(t, header.Get("Traceparent"))

	spans := h.recorder.Ended()
	require.Len(t, spans, 1)
	assert.Equal(t, "UserService/GetUser", spans[0].Name())
	assert.Contains(t, h.metricNames(t), "rpc.client.call.duration")
}

func TestClient_ErrorLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	client := httpconv.NewClient(append(h.opts, semconv.WithClientService("pkg", "Svc"))...)
	_, call := client.Start(context.Background(), "M")
	call.RecordError(errors.New("boom"), 502)
	call.End()

	span := h.recorder.Ended()[0]
	assert.Equal(t, codes.Error, span.Status().Code)
	code, ok := attrValue(span.Attributes(), "gotsrpc.error.code")
	require.True(t, ok)
	assert.Equal(t, int64(502), code.AsInt64())
}

func TestClient_SetStatusMarksError(t *testing.T) {
	t.Parallel()

	h := newHarness()

	client := httpconv.NewClient(append(h.opts, semconv.WithClientService("pkg", "Svc"))...)
	_, call := client.Start(context.Background(), "M")
	call.SetStatus(http.StatusBadGateway)
	call.End()

	span := h.recorder.Ended()[0]
	assert.Equal(t, codes.Error, span.Status().Code)
	et, ok := attrValue(span.Attributes(), "error.type")
	require.True(t, ok)
	assert.Equal(t, http.StatusText(http.StatusBadGateway), et.AsString())
}
