package sarama

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	cNamespace = "sarama"
	cSubsystem = "consumer"
)

type prometheusConsumerMetrics struct {
	fetchRequestDuration *prometheus.HistogramVec
}

var cMetrics *prometheusConsumerMetrics

func init() {
	fetchRequestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "fetch_request_duration_seconds",
		Help:      "Duration of send() requests in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"partition", "broker"})

	prometheus.MustRegister(fetchRequestDuration)

	cMetrics = &prometheusConsumerMetrics{
		fetchRequestDuration: fetchRequestDuration,
	}
}
