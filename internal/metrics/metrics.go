package metrics

import "github.com/prometheus/client_golang/prometheus"

var AnalyzeRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "analyze_requests_total",
		Help: "Total analyze requests",
	},
)

func Init() {
	prometheus.MustRegister(
		AnalyzeRequests,
	)
}