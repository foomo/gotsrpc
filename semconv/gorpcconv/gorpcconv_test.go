package gorpcconv_test

import (
	"errors"
	"testing"

	"github.com/foomo/gotsrpc/v3/semconv"
	"github.com/foomo/gotsrpc/v3/semconv/gorpcconv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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

	srv := gorpcconv.NewServer("github.com/foo/bar", "UserService", h.opts...)
	call := srv.Handle("GetUser")
	call.End()

	spans := h.recorder.Ended()
	require.Len(t, spans, 1)
	span := spans[0]
	assert.Equal(t, "UserService/GetUser", span.Name())
	assert.Equal(t, codes.Unset, span.Status().Code)

	enc, ok := attrValue(span.Attributes(), "gotsrpc.encoding")
	require.True(t, ok)
	assert.Equal(t, "gob", enc.AsString())

	assert.Contains(t, h.metricNames(t), "rpc.server.call.duration")
}

func TestServer_ErrorLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	srv := gorpcconv.NewServer("pkg", "Svc", h.opts...)
	call := srv.Handle("M")
	call.RecordError(errors.New("boom"), 500)
	call.End()

	span := h.recorder.Ended()[0]
	assert.Equal(t, codes.Error, span.Status().Code)
	code, ok := attrValue(span.Attributes(), "gotsrpc.error.code")
	require.True(t, ok)
	assert.Equal(t, int64(500), code.AsInt64())
}

func TestClient_SuccessLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	client := gorpcconv.NewClient(append(h.opts, semconv.WithClientService("github.com/foo/bar", "UserService"))...)
	call := client.Start("GetUser")
	call.End()

	spans := h.recorder.Ended()
	require.Len(t, spans, 1)
	assert.Equal(t, "UserService/GetUser", spans[0].Name())
	assert.Contains(t, h.metricNames(t), "rpc.client.call.duration")
}

func TestClient_ErrorLifecycle(t *testing.T) {
	t.Parallel()

	h := newHarness()

	client := gorpcconv.NewClient(append(h.opts, semconv.WithClientService("pkg", "Svc"))...)
	call := client.Start("M")
	call.RecordError(errors.New("boom"), 502)
	call.End()

	span := h.recorder.Ended()[0]
	assert.Equal(t, codes.Error, span.Status().Code)
	code, ok := attrValue(span.Attributes(), "gotsrpc.error.code")
	require.True(t, ok)
	assert.Equal(t, int64(502), code.AsInt64())
}
