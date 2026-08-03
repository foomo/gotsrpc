package gorpcconv

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	rpcconv "go.opentelemetry.io/otel/semconv/v1.41.0/rpcconv"
	"go.opentelemetry.io/otel/trace"

	"github.com/foomo/gotsrpc/v3/semconv"
)

// Client instruments the outgoing gorpc calls of one service.
type Client struct {
	tracer   trace.Tracer
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
		duration: duration,
		pkg:      cfg.Package,
		service:  cfg.Service,
	}
}

// Start begins a local RPC client span for the given method.
func (c *Client) Start(method string) *ClientCall {
	ctx, span := c.tracer.Start(context.Background(), c.service+"/"+method,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			semconv.RPCSystemGoTSRPC,
			semconv.RPCMethod(c.service, method),
			semconv.Package(c.pkg),
			semconv.Encoding("gob"),
		),
	)

	return &ClientCall{c: c, ctx: ctx, span: span, method: method, start: time.Now()}
}

// ClientCall records the lifecycle of a single gorpc client call.
type ClientCall struct {
	c       *Client
	ctx     context.Context
	span    trace.Span
	method  string
	start   time.Time
	errType string
}

// RecordError records a client-side error.
func (cc *ClientCall) RecordError(err error, code int) {
	if err == nil {
		return
	}

	et := semconv.ErrorType(err)
	cc.errType = et.Value.AsString()
	cc.span.SetStatus(codes.Error, err.Error())
	cc.span.SetAttributes(et, semconv.GoTSRPCErrorCode(code))
}

// End records the duration metric and finishes the span.
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
