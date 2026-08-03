// Package gorpcconv provides OpenTelemetry instrumentation for the gotsrpc
// binary (valyala/gorpc) transport. The gorpc protocol has no header carrier,
// so there is no cross-process trace-context propagation: spans are local
// (parentless) and carry the rpc.* attributes plus a duration metric.
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

// Server instruments the incoming gorpc calls of one service.
type Server struct {
	tracer   trace.Tracer
	logger   *slog.Logger
	duration rpcconv.ServerCallDuration
	pkg      string
	service  string
}

// NewServer builds the server instrumentation for the given service identity.
func NewServer(pkg, service string, opts ...semconv.Option) *Server {
	cfg := semconv.NewConfig(opts...)

	duration, err := rpcconv.NewServerCallDuration(semconv.Meter(cfg))
	if err != nil {
		cfg.Logger.Error("gotsrpc: failed to create rpc.server.call.duration metric", slog.Any("error", err))
	}

	return &Server{
		tracer:   semconv.Tracer(cfg),
		logger:   cfg.Logger,
		duration: duration,
		pkg:      pkg,
		service:  service,
	}
}

// Handle begins a local RPC server span for the given method.
func (s *Server) Handle(method string) *ServerCall {
	ctx, span := s.tracer.Start(context.Background(), s.service+"/"+method,
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(
			semconv.RPCSystemGoTSRPC,
			semconv.RPCMethod(s.service, method),
			semconv.Package(s.pkg),
			semconv.Encoding("gob"),
		),
	)

	return &ServerCall{s: s, ctx: ctx, span: span, method: method, start: time.Now()}
}

// ServerCall records the lifecycle of a single gorpc server call.
type ServerCall struct {
	s       *Server
	ctx     context.Context
	span    trace.Span
	method  string
	start   time.Time
	errType string
}

// RecordError records a handler error.
func (c *ServerCall) RecordError(err error, code int) {
	if err == nil {
		return
	}

	et := semconv.ErrorType(err)
	c.errType = et.Value.AsString()
	c.span.SetStatus(codes.Error, err.Error())
	c.span.SetAttributes(et, semconv.GoTSRPCErrorCode(code))

	sc := c.span.SpanContext()
	c.s.logger.ErrorContext(c.ctx, "gotsrpc: gorpc call failed",
		slog.String("rpc.method", c.s.service+"/"+c.method),
		slog.String("error.type", c.errType),
		slog.String("error", err.Error()),
		slog.String("trace_id", sc.TraceID().String()),
		slog.String("span_id", sc.SpanID().String()),
	)
}

// End records the duration metric and finishes the span.
func (c *ServerCall) End() {
	attrs := []attribute.KeyValue{
		semconv.RPCSystemGoTSRPC,
		semconv.RPCMethod(c.s.service, c.method),
	}
	if c.errType != "" {
		attrs = append(attrs, semconv.ErrorTypeString(c.errType))
	}

	c.s.duration.RecordSet(c.ctx, time.Since(c.start).Seconds(), attribute.NewSet(attrs...))
	c.span.End()
}
