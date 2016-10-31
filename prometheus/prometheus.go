package prometheus

import (
	"net/http"

	"github.com/foomo/gotsrpc"
	p "github.com/prometheus/client_golang/prometheus"
)

func InstrumentService(namespace string, s http.HandlerFunc) (handler http.HandlerFunc) {
	counterVec := p.NewSummaryVec(p.SummaryOpts{
		Namespace: namespace,
		Subsystem: "gotsrpc_service",
		Name:      "time_nanoseconds",
		Help:      "milliseconds to unmarshal, execute and marshal requests and reponses",
	}, []string{"package", "service", "func", "type"})
	p.MustRegister(counterVec)
	return func(w http.ResponseWriter, r *http.Request) {
		r = gotsrpc.RequestWithStatsContext(r)
		s.ServeHTTP(w, r)
		stats := gotsrpc.GetStatsForRequest(r)
		if stats != nil {
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "unmarshalling").Observe(float64(stats.Unmarshalling))
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "execution").Observe(float64(stats.Execution))
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "marshalling").Observe(float64(stats.Marshalling))
		}
	}
}
