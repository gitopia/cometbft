// Code generated by metricsgen. DO NOT EDIT.

package state

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		BlockProcessingTime: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_processing_time",
			Help:      "Time spent processing FinalizeBlock",

			Buckets: stdprometheus.LinearBuckets(1, 10, 10),
		}, labels).With(labelsAndValues...),
		ConsensusParamUpdates: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "consensus_param_updates",
			Help:      "Number of consensus parameter updates returned by the application since process start.",
		}, labels).With(labelsAndValues...),
		ValidatorSetUpdates: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_set_updates",
			Help:      "Number of validator set updates returned by the application since process start.",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		BlockProcessingTime:   discard.NewHistogram(),
		ConsensusParamUpdates: discard.NewCounter(),
		ValidatorSetUpdates:   discard.NewCounter(),
	}
}
