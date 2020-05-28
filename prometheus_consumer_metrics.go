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
	fetchSizeBytes       *prometheus.HistogramVec
	blockSizeRecords     *prometheus.HistogramVec
}

var cMetrics *prometheusConsumerMetrics

func init() {
	fetchRequestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cNamespace,
		Subsystem: cSubsystem,
		Name:      "fetch_request_duration_seconds",
		Help:      "Duration of send() requests in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"topic", "partition", "broker"})

	blockSizeRecords := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cNamespace,
		Subsystem: cSubsystem,
		Name:      "response_block_size_records",
		Help:      "The size of each block in response in records",
		Buckets:   []float64{100, 1e3, 1e5, 1e6, 1e8},
	}, []string{"topic", "partition"})

	fetchSizeBytes := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cNamespace,
		Subsystem: cSubsystem,
		Name:      "fetch_size_bytes",
		Help:      "The partitoinConsumer.fetchSize might be changed based on some logic of sarama and/or Kafka responses, the size in the bytes",
		Buckets:   []float64{100, 1e3, 1e5, 1e6, 1e8},
	}, []string{"topic", "partition"})

	prometheus.MustRegister(
		fetchRequestDuration,
		fetchSizeBytes,
		blockSizeRecords,
	)

	cMetrics = &prometheusConsumerMetrics{
		fetchRequestDuration: fetchRequestDuration,
		fetchSizeBytes:       fetchSizeBytes,
		blockSizeRecords:     blockSizeRecords,
	}
}
