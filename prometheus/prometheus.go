package prometheus

import (
	"fmt"
	"net/http"

	"github.com/foomo/gotsrpc"
	p "github.com/prometheus/client_golang/prometheus"
)

func InstrumentService(s http.HandlerFunc) (handler http.HandlerFunc) {
	counterVec := p.NewSummaryVec(p.SummaryOpts{
		Namespace: "gotsrpc",
		Subsystem: "service",
		Name:      "time_nanoseconds",
		Help:      "nanoseconds to unmarshal requests, execute a service function and marshal its reponses",
	}, []string{"package", "service", "func", "type"})
	p.MustRegister(counterVec)

	return func(w http.ResponseWriter, r *http.Request) {
		*r = *gotsrpc.RequestWithStatsContext(r)
		s.ServeHTTP(w, r)
		stats := gotsrpc.GetStatsForRequest(r)
		if stats != nil {
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "unmarshalling").Observe(float64(stats.Unmarshalling))
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "execution").Observe(float64(stats.Execution))
			counterVec.WithLabelValues(stats.Package, stats.Service, stats.Func, "marshalling").Observe(float64(stats.Marshalling))
			fmt.Println("Observed execution time", stats.Package+"."+stats.Service+"."+stats.Func, stats.Execution, float64(stats.Execution))
		}
	}
}
