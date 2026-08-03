// Package httpconv provides OpenTelemetry instrumentation for the gotsrpc
// HTTP transport. It emits only rpc.* spans and metrics (never http.*), so it
// composes cleanly with otelhttp: when a consumer wraps the handler with
// otelhttp.NewHandler and/or the client transport with otelhttp.NewTransport,
// the gotsrpc RPC spans nest under the otelhttp HTTP spans and no metrics are
// duplicated.
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

// Server instruments the incoming (server-side) HTTP RPC calls of one service.
type Server struct {
	tracer   trace.Tracer
	prop     propagation.TextMapPropagator
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
		prop:     cfg.Propagators,
		logger:   cfg.Logger,
		duration: duration,
		pkg:      pkg,
		service:  service,
	}
}

// Start begins an RPC server span for the call and returns the request carrying
// the span context together with the ServerCall used to record the outcome.
//
// If no active span is present in the request context (otelhttp is not in front
// of us) the remote parent is extracted from the request headers; if otelhttp
// already started an http.server span, the RPC span simply nests under it.
func (s *Server) Start(_ http.ResponseWriter, r *http.Request, method string) (*http.Request, *ServerCall) {
	ctx := r.Context()
	if !trace.SpanContextFromContext(ctx).IsValid() {
		ctx = s.prop.Extract(ctx, propagation.HeaderCarrier(r.Header))
	}

	ctx, span := s.tracer.Start(ctx, s.service+"/"+method,
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(
			semconv.RPCSystemGoTSRPC,
			semconv.RPCMethod(s.service, method),
			semconv.Package(s.pkg),
			semconv.Encoding(semconv.EncodingFromContentType(r.Header.Get("Content-Type"))),
		),
	)

	return r.WithContext(ctx), &ServerCall{
		s:      s,
		ctx:    ctx,
		span:   span,
		method: method,
		start:  time.Now(),
	}
}

// ServerCall records the lifecycle of a single server-side RPC call. Its
// RecordUnmarshal/RecordMarshal/RecordError methods are called by
// gotsrpc.LoadArgs and gotsrpc.Reply.
type ServerCall struct {
	s         *Server
	ctx       context.Context
	span      trace.Span
	method    string
	start     time.Time
	execStart time.Time
	errType   string
	logged    bool
}

// ExecutionStart marks the beginning of the service method execution.
func (c *ServerCall) ExecutionStart() { c.execStart = time.Now() }

// ExecutionEnd records the service method execution duration on the span.
func (c *ServerCall) ExecutionEnd() {
	if c.execStart.IsZero() {
		return
	}

	c.span.SetAttributes(semconv.ExecutionMicros(time.Since(c.execStart)))
}

// RecordUnmarshal records the request decode duration and size (called by gotsrpc.LoadArgs).
func (c *ServerCall) RecordUnmarshal(d time.Duration, requestSize int) {
	c.span.SetAttributes(semconv.UnmarshalMicros(d), semconv.RequestSizeAttr(requestSize))
}

// RecordMarshal records the response encode duration and size (called by gotsrpc.Reply).
func (c *ServerCall) RecordMarshal(d time.Duration, responseSize int) {
	c.span.SetAttributes(semconv.MarshalMicros(d), semconv.ResponseSizeAttr(responseSize))
}

// RecordError records a service error and its gotsrpc error code (called by gotsrpc.Reply).
func (c *ServerCall) RecordError(err error, code int) {
	if err == nil {
		return
	}

	et := semconv.ErrorType(err)
	c.errType = et.Value.AsString()
	c.span.SetStatus(codes.Error, err.Error())
	c.span.SetAttributes(et, semconv.GoTSRPCErrorCode(code))
	c.log(err)
}

// SetError records a transport-level failure (failed to load args / reply).
func (c *ServerCall) SetError(err error) {
	if err == nil {
		return
	}

	et := semconv.ErrorType(err)
	c.errType = et.Value.AsString()
	c.span.SetStatus(codes.Error, err.Error())
	c.span.SetAttributes(et)
	c.log(err)
}

// SetStatus marks the span errored for a non-2xx session-style handler status.
// The http.* status attribute is deliberately left to otelhttp.
func (c *ServerCall) SetStatus(status int) {
	if status >= http.StatusBadRequest && c.errType == "" {
		c.errType = http.StatusText(status)
		c.span.SetStatus(codes.Error, http.StatusText(status))
		c.span.SetAttributes(semconv.ErrorTypeString(c.errType))
	}
}

// NotFound marks the call as targeting an unknown method.
func (c *ServerCall) NotFound() {
	c.errType = "method_not_found"
	c.span.SetStatus(codes.Error, "method not found")
	c.span.SetAttributes(semconv.ErrorTypeString(c.errType))
}

// End records the RPC duration metric and finishes the span.
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

func (c *ServerCall) log(err error) {
	if c.logged {
		return
	}

	c.logged = true
	sc := c.span.SpanContext()
	c.s.logger.ErrorContext(c.ctx, "gotsrpc: rpc call failed",
		slog.String("rpc.method", c.s.service+"/"+c.method),
		slog.String("error.type", c.errType),
		slog.String("error", err.Error()),
		slog.String("trace_id", sc.TraceID().String()),
		slog.String("span_id", sc.SpanID().String()),
	)
}
