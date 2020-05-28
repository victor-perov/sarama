package sarama

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	bNamespace = "sarama"
	bSubsystem = "broker"
)

type prometheusBrokerMetrics struct {
	errorsTotal *prometheus.CounterVec

	reqSendDuration    *prometheus.HistogramVec
	reqReceiveDuration *prometheus.HistogramVec
	reqReceiveBytes    *prometheus.HistogramVec

	joinGroupDuration  *prometheus.HistogramVec
	syncGroupDuration  *prometheus.HistogramVec
	heartBeatDuration  *prometheus.HistogramVec
	fetchThrottleTime  *prometheus.HistogramVec
	leaveGroupDuration *prometheus.HistogramVec
}

var bMetrics *prometheusBrokerMetrics

func init() {
	errorsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "errors_total",
		Help:      "Total number of received errors on send()",
	}, []string{"request_type", "broker"})

	fetchThrottleTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "fetch_throttle_seconds",
		Help:      "Broker Fetch() request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"topic", "partition", "broker"})

	prometheus.MustRegister(
		errorsTotal,
		fetchThrottleTime,
	)

	bMetrics = &prometheusBrokerMetrics{
		errorsTotal: errorsTotal,

		fetchThrottleTime: fetchThrottleTime,
	}
}
