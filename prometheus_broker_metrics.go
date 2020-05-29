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

	reqSendDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "request_send_duration_seconds",
		Help:      "Duration of write request into connection with send() method in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"request_type", "broker"})
	reqReceiveDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "request_receive_duration_seconds",
		Help:      "Elapsed time of reading header (first 8b) and body from connection by send() in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"request_type", "broker"})
	reqReceiveBytes := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "request_receive_bytes",
		Help:      "Amount of bytes that was read from connection",
		Buckets:   []float64{10, 100, 1000, 5000, 10000, 1e6, 1e7},
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

	fetchThrottleTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "fetch_throttle_seconds",
		Help:      "Broker Fetch() throttle time in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"topic", "partition", "broker"})

	leaveGroupDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: bNamespace,
		Subsystem: bSubsystem,
		Name:      "leave_group_duration_seconds",
		Help:      "LeaveGroup request duration in seconds",
		Buckets:   []float64{1, 2, 5, 10, 60, 120, 300, 600, 900, 1800},
	}, []string{"member", "group", "broker"})

	prometheus.MustRegister(
		errorsTotal,
		reqSendDuration,
		reqReceiveDuration,
		reqReceiveBytes,
		joinGroupDuration,
		syncGroupDuration,
		heartBeatDuration,
		fetchThrottleTime,
		leaveGroupDuration,
	)

	bMetrics = &prometheusBrokerMetrics{
		errorsTotal: errorsTotal,

		reqSendDuration:    reqSendDuration,
		reqReceiveDuration: reqReceiveDuration,
		reqReceiveBytes:    reqReceiveBytes,

		joinGroupDuration:  joinGroupDuration,
		syncGroupDuration:  syncGroupDuration,
		heartBeatDuration:  heartBeatDuration,
		fetchThrottleTime:  fetchThrottleTime,
		leaveGroupDuration: leaveGroupDuration,
	}
}
