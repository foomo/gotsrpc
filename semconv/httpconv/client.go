package httpconv

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/semconv/v1.41.0/rpcconv"
	"go.opentelemetry.io/otel/trace"

	"github.com/foomo/gotsrpc/v3/semconv"
)

// Client instruments the outgoing (client-side) HTTP RPC calls of one service.
type Client struct {
	tracer   trace.Tracer
	prop     propagation.TextMapPropagator
	duration rpcconv.ClientCallDuration
	pkg      string
	service  string
}

// NewClient builds the client instrumentation. The service identity is read
// from the config (seed it with semconv.WithClientService).
func NewClient(opts ...semconv.Option) *Client {
	cfg := semconv.NewConfig(opts...)

	duration, err := rpcconv.NewClientCallDuration(semconv.Meter(cfg))
	if err != nil {
		cfg.Logger.Error("gotsrpc: failed to create rpc.client.call.duration metric", slog.Any("error", err))
	}

	return &Client{
		tracer:   semconv.Tracer(cfg),
		prop:     cfg.Propagators,
		duration: duration,
		pkg:      cfg.Package,
		service:  cfg.Service,
	}
}

// Start begins an RPC client span for the call and returns the derived context
// (carrying the span) plus the ClientCall used to record the outcome.
func (c *Client) Start(ctx context.Context, method string) (context.Context, *ClientCall) {
	ctx, span := c.tracer.Start(ctx, c.service+"/"+method,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			semconv.RPCSystemGoTSRPC,
			semconv.RPCMethod(c.service, method),
			semconv.Package(c.pkg),
		),
	)

	return ctx, &ClientCall{
		c:      c,
		ctx:    ctx,
		span:   span,
		method: method,
		start:  time.Now(),
	}
}

// ClientCall records the lifecycle of a single client-side RPC call.
type ClientCall struct {
	c       *Client
	ctx     context.Context
	span    trace.Span
	method  string
	start   time.Time
	errType string
}

// Inject writes the current trace context into the outgoing request headers.
// This is a no-op with the default (no-op) global propagator; if otelhttp's
// transport wraps the client it re-injects afterwards (idempotent).
func (cc *ClientCall) Inject(h http.Header) {
	cc.c.prop.Inject(cc.ctx, propagation.HeaderCarrier(h))
}

// RecordRequestSize records the request payload size (bytes) on the span.
func (cc *ClientCall) RecordRequestSize(size int) {
	cc.span.SetAttributes(semconv.RequestSizeAttr(size))
}

// RecordResponseSize records the response payload size (bytes) on the span.
func (cc *ClientCall) RecordResponseSize(size int) {
	cc.span.SetAttributes(semconv.ResponseSizeAttr(size))
}

// SetStatus marks the span errored for a non-2xx transport status. The http.*
// status attribute is deliberately left to otelhttp.
func (cc *ClientCall) SetStatus(status int) {
	if status >= http.StatusBadRequest && cc.errType == "" {
		cc.errType = http.StatusText(status)
		cc.span.SetStatus(codes.Error, http.StatusText(status))
		cc.span.SetAttributes(semconv.ErrorTypeString(cc.errType))
	}
}

// RecordError records a client-side error and its gotsrpc error code.
func (cc *ClientCall) RecordError(err error, code int) {
	if err == nil {
		return
	}

	et := semconv.ErrorType(err)
	cc.errType = et.Value.AsString()
	cc.span.SetStatus(codes.Error, err.Error())
	cc.span.SetAttributes(et, semconv.GoTSRPCErrorCode(code))
}

// End records the RPC duration metric and finishes the span.
func (cc *ClientCall) End() {
	attrs := []attribute.KeyValue{
		semconv.RPCSystemGoTSRPC,
		semconv.RPCMethod(cc.c.service, cc.method),
	}
	if cc.errType != "" {
		attrs = append(attrs, semconv.ErrorTypeString(cc.errType))
	}

	cc.c.duration.RecordSet(cc.ctx, time.Since(cc.start).Seconds(), attribute.NewSet(attrs...))
	cc.span.End()
}
