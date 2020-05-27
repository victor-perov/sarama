package sarama

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	bNamespace = "sarama"
	bSubsystem = "broker"
)

type prometheusBrokerMetrics struct {
	errorsTotal        *prometheus.CounterVec
	requestDuration    *prometheus.HistogramVec
	joinGroupDuration  *prometheus.HistogramVec
	syncGroupDuration  *prometheus.HistogramVec
	leaveGroupDuration *prometheus.HistogramVec
	heartBeatDuration  *prometheus.HistogramVec
}

var bMetrics *prometheusBrokerMetrics

func init() {
	errorsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "errors_total",
		Help:      "Total number of received errors on send()",
	}, []string{"request_type", "broker"})
	requestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "request_duration_seconds",
		Help:      "Duration of send() requests in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"request_type", "broker"})

	joinGroupDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "join_group_duration_seconds",
		Help:      "JoinGroup request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"member", "group", "broker"})

	syncGroupDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "sync_group_duration_seconds",
		Help:      "SyncGroup request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"member", "group", "broker"})

	heartBeatDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "heart_beat_duration_seconds",
		Help:      "HeartBeat request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"member", "group", "broker"})

	leaveGroupDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "leave_group_duration_seconds",
		Help:      "LeaveGroup request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"member", "group", "broker"})

	prometheus.MustRegister(errorsTotal, requestDuration, joinGroupDuration, syncGroupDuration, heartBeatDuration, leaveGroupDuration)

	bMetrics = &prometheusBrokerMetrics{
		errorsTotal:        errorsTotal,
		requestDuration:    requestDuration,
		joinGroupDuration:  joinGroupDuration,
		syncGroupDuration:  syncGroupDuration,
		heartBeatDuration:  heartBeatDuration,
		leaveGroupDuration: leaveGroupDuration,
	}
}
