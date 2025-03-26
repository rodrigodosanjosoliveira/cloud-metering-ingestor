package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	PulsesReceived = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pulses_received_total",
		Help: "Total number of pulses received",
	})

	FlushTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "flush_total",
		Help: "Total number of flush operations performed",
	})

	AggregationCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_aggregation",
		Help: "Number of active aggregation groups",
	})
)

func Init() {
	prometheus.MustRegister(PulsesReceived)
	prometheus.MustRegister(FlushTotal)
	prometheus.MustRegister(AggregationCount)
}
