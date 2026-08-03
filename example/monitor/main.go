package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/foomo/gotsrpc/v3/example/monitor/service"
)

// setupOTel wires a minimal OpenTelemetry SDK writing to stdout so the example
// shows gotsrpc's rpc.* traces and metrics. Returns a shutdown func that
// flushes the exporters.
func setupOTel(_ context.Context) (func(context.Context) error, error) {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExporter))
	otel.SetTracerProvider(tracerProvider)

	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(
		sdkmetric.NewPeriodicReader(metricExporter, sdkmetric.WithInterval(5*time.Second)),
	))
	otel.SetMeterProvider(meterProvider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return func(ctx context.Context) error {
		return errors.Join(tracerProvider.Shutdown(ctx), meterProvider.Shutdown(ctx))
	}, nil
}

func main() {
	ctx := context.Background()

	shutdown, err := setupOTel(ctx)
	if err != nil {
		panic(err)
	}
	defer func() { _ = shutdown(ctx) }()

	fs := http.FileServer(http.Dir("./monitor/client"))
	ws := service.NewDefaultServiceGoTSRPCProxy(&service.Handler{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/service/"):
			ws.ServeHTTP(w, r)
		default:
			fs.ServeHTTP(w, r)
		}
	})

	// Wrap the handler with otelhttp: the http.server span/metrics come from
	// otelhttp, the rpc.server span/metrics from gotsrpc — the RPC span nests
	// under the HTTP span and no metrics are duplicated.
	handler := otelhttp.NewHandler(mux, "gotsrpc")

	go func() {
		time.Sleep(time.Second)

		_ = exec.CommandContext(ctx, "open", "http://127.0.0.1:3000").Run()
		call(ctx)
	}()

	panic(http.ListenAndServe("localhost:3000", handler)) //nolint:gosec
}

func call(ctx context.Context) {
	// An otelhttp-wrapped transport contributes the http.client span/metrics;
	// gotsrpc contributes the rpc.client span/metrics on top.
	httpClient := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	c := service.NewServiceGoTSRPCClientWithClient("http://127.0.0.1:3000", "/service", httpClient)

	res, _ := c.Hello(ctx, "Hello World")
	fmt.Println(res)
}
