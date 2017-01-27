package prometheus

import (
	"net/http"

	"github.com/foomo/gotsrpc"
	p "github.com/prometheus/client_golang/prometheus"
)

func InstrumentService(s http.HandlerFunc) (handler http.HandlerFunc) {
	requestDuration := p.NewSummaryVec(p.SummaryOpts{
		Namespace: "gotsrpc",
		Subsystem: "service",
		Name:      "time_nanoseconds",
		Help:      "nanoseconds to unmarshal requests, execute a service function and marshal its reponses",
	}, []string{"package", "service", "func", "type"})
	requestSize := p.NewSummaryVec(p.SummaryOpts{
		Namespace: "gotsrpc",
		Subsystem: "service",
		Name:      "size_bytes",
		Help:      "request and response sizes in bytes",
	}, []string{"package", "service", "func", "type"})

	p.MustRegister(requestSize)
	p.MustRegister(requestDuration)

	return func(w http.ResponseWriter, r *http.Request) {
		*r = *gotsrpc.RequestWithStatsContext(r)
		s.ServeHTTP(w, r)
		stats := gotsrpc.GetStatsForRequest(r)
		if stats != nil {
			requestSize.WithLabelValues(stats.Package, stats.Service, stats.Func, "request").Observe(float64(stats.RequestSize))
			requestSize.WithLabelValues(stats.Package, stats.Service, stats.Func, "response").Observe(float64(stats.ResponseSize))
			requestDuration.WithLabelValues(stats.Package, stats.Service, stats.Func, "unmarshalling").Observe(float64(stats.Unmarshalling))
			requestDuration.WithLabelValues(stats.Package, stats.Service, stats.Func, "execution").Observe(float64(stats.Execution))
			requestDuration.WithLabelValues(stats.Package, stats.Service, stats.Func, "marshalling").Observe(float64(stats.Marshalling))
		}
	}
}

func InstrumentGoRPCService() gotsrpc.GoRPCCallStatsHandlerFun {
	callsCounter := p.NewSummaryVec(p.SummaryOpts{
		Namespace: "gorpc",
		Subsystem: "service",
		Name:      "time_seconds",
		Help:      "seconds to execute a service method",
	}, []string{"package", "service", "func", "type"})
	p.MustRegister(callsCounter)

	return func(stats *gotsrpc.CallStats) {
		callsCounter.WithLabelValues(stats.Package, stats.Service, stats.Func, "execution").Observe(float64(stats.Execution))
	}
}
