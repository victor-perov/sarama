package sarama

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	cgNamespace = "sarama"
	cgSubsystem = "consumer_group"
)

type prometheusConsumerGroupMetrics struct {
	sessionRetriesTotal *prometheus.CounterVec
}

var cgMetrics *prometheusConsumerGroupMetrics

func init() {
	sessionRetriesTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: cgNamespace,
		Subsystem: cgSubsystem,
		Name:      "session_retries_total",
		Help:      "Total number of received retryNewSession()",
	}, []string{"member", "group"})

	prometheus.MustRegister(sessionRetriesTotal)

	cgMetrics = &prometheusConsumerGroupMetrics{
		sessionRetriesTotal: sessionRetriesTotal,
	}
}
