package aggregator

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	AggregatorNamespace = "sffl_aggregator"
)

type AggregatorEventListener interface {
	ObserveLastStateRootUpdateAggregated(rollupId uint32, blockNumber uint64)
	ObserveLastStateRootUpdateReceived(rollupId uint32, blockNumber uint64)
	ObserveLastOperatorSetUpdateAggregated(operatorSetUpdateId uint64)
	ObserveLastOperatorSetUpdateReceived(operatorSetUpdateId uint64)
	IncExpiredMessages()
	IncExpiredTasks()
	IncErroredSubmissions()
	IncAggregatorInitializations()
	ObserveLastCheckpointReferenceSent(referenceId uint32)
	ObserveLastCheckpointTaskReferenceReceived(referenceId uint32)
	ObserveLastCheckpointTaskReferenceAggregated(referenceId uint32)
}

type SelectiveAggregatorListener struct {
	ObserveLastStateRootUpdateAggregatedCb         func(rollupId uint32, blockNumber uint64)
	ObserveLastStateRootUpdateReceivedCb           func(rollupId uint32, blockNumber uint64)
	ObserveLastOperatorSetUpdateAggregatedCb       func(operatorSetUpdateId uint64)
	ObserveLastOperatorSetUpdateReceivedCb         func(operatorSetUpdateId uint64)
	IncExpiredMessagesCb                           func()
	IncExpiredTasksCb                              func()
	IncErroredSubmissionsCb                        func()
	IncAggregatorInitializationsCb                 func()
	ObserveLastCheckpointReferenceSentCb           func(referenceId uint32)
	ObserveLastCheckpointTaskReferenceReceivedCb   func(referenceId uint32)
	ObserveLastCheckpointTaskReferenceAggregatedCb func(referenceId uint32)
}

func (l *SelectiveAggregatorListener) ObserveLastOperatorSetUpdateAggregated(operatorSetUpdateId uint64) {
	if l.ObserveLastOperatorSetUpdateAggregatedCb != nil {
		l.ObserveLastOperatorSetUpdateAggregatedCb(operatorSetUpdateId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastOperatorSetUpdateReceived(operatorSetUpdateId uint64) {
	if l.ObserveLastOperatorSetUpdateReceivedCb != nil {
		l.ObserveLastOperatorSetUpdateReceivedCb(operatorSetUpdateId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastStateRootUpdateAggregated(rollupId uint32, blockNumber uint64) {
	if l.ObserveLastStateRootUpdateAggregatedCb != nil {
		l.ObserveLastStateRootUpdateAggregatedCb(rollupId, blockNumber)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastStateRootUpdateReceived(rollupId uint32, blockNumber uint64) {
	if l.ObserveLastStateRootUpdateReceivedCb != nil {
		l.ObserveLastStateRootUpdateReceivedCb(rollupId, blockNumber)
	}
}

func (l *SelectiveAggregatorListener) IncExpiredMessages() {
	if l.IncExpiredMessagesCb != nil {
		l.IncExpiredMessagesCb()
	}
}

func (l *SelectiveAggregatorListener) IncExpiredTasks() {
	if l.IncExpiredTasksCb != nil {
		l.IncExpiredTasksCb()
	}
}

func (l *SelectiveAggregatorListener) IncErroredSubmissions() {
	if l.IncErroredSubmissionsCb != nil {
		l.IncErroredSubmissionsCb()
	}
}

func (l *SelectiveAggregatorListener) IncAggregatorInitializations() {
	if l.IncAggregatorInitializationsCb != nil {
		l.IncAggregatorInitializationsCb()
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointReferenceSent(referenceId uint32) {
	if l.ObserveLastCheckpointReferenceSentCb != nil {
		l.ObserveLastCheckpointReferenceSentCb(referenceId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointTaskReferenceReceived(referenceId uint32) {
	if l.ObserveLastCheckpointTaskReferenceReceivedCb != nil {
		l.ObserveLastCheckpointTaskReferenceReceivedCb(referenceId)
	}
}

func (l *SelectiveAggregatorListener) ObserveLastCheckpointTaskReferenceAggregated(referenceId uint32) {
	if l.ObserveLastCheckpointTaskReferenceAggregatedCb != nil {
		l.ObserveLastCheckpointTaskReferenceAggregatedCb(referenceId)
	}
}

func MakeAggregatorMetrics(registry *prometheus.Registry) (AggregatorEventListener, error) {
	lastStateRootUpdateAggregated := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_state_root_update_aggregated",
			Help:      "Last state root update aggregated per rollup ID",
		},
		[]string{"rollup_id"},
	)
	if err := registry.Register(lastStateRootUpdateAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastStateRootUpdateAggregated gauge: %w", err)
	}

	lastStateRootUpdateReceived := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_state_root_update_received",
			Help:      "Last state root update received per rollup ID",
		},
		[]string{"rollup_id"},
	)
	if err := registry.Register(lastStateRootUpdateReceived); err != nil {
		return nil, fmt.Errorf("error registering lastStateRootUpdateReceived gauge: %w", err)
	}

	lastOperatorSetUpdateAggregated := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_operator_set_update_aggregated",
			Help:      "Last operator set update aggregated",
		},
	)
	if err := registry.Register(lastOperatorSetUpdateAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastOperatorSetUpdateAggregated gauge: %w", err)
	}

	lastOperatorSetUpdateReceived := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_operator_set_update_received",
			Help:      "Last operator set update received",
		},
	)
	if err := registry.Register(lastOperatorSetUpdateReceived); err != nil {
		return nil, fmt.Errorf("error registering lastBlockReceived gauge: %w", err)
	}

	expiredMessages := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "expired_messages_total",
			Help:      "Total number of expired messages",
		},
	)
	if err := registry.Register(expiredMessages); err != nil {
		return nil, fmt.Errorf("error registering expiredMessages counter: %w", err)
	}

	expiredTasks := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "expired_tasks_total",
			Help:      "Total number of expired tasks",
		},
	)
	if err := registry.Register(expiredTasks); err != nil {
		return nil, fmt.Errorf("error registering expiredTasks counter: %w", err)
	}

	erroredSubmissions := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "errored_submissions_total",
			Help:      "Total number of errored submissions",
		},
	)
	if err := registry.Register(erroredSubmissions); err != nil {
		return nil, fmt.Errorf("error registering erroredSubmissions counter: %w", err)
	}

	aggregatorInitializations := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: AggregatorNamespace,
			Name:      "reinitializations_total",
			Help:      "Total number of aggregator reinitializations",
		},
	)
	if err := registry.Register(aggregatorInitializations); err != nil {
		return nil, fmt.Errorf("error registering aggregatorInitializations counter: %w", err)
	}

	lastCheckpointReferenceSent := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_reference_sent",
			Help:      "Last checkpoint reference sent",
		},
	)
	if err := registry.Register(lastCheckpointReferenceSent); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointReferenceSent gauge: %w", err)
	}

	lastCheckpointTaskReferenceReceived := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_task_reference_received",
			Help:      "Last checkpoint task reference received",
		},
	)
	if err := registry.Register(lastCheckpointTaskReferenceReceived); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointTaskReferenceReceived gauge: %w", err)
	}

	lastCheckpointTaskReferenceAggregated := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: AggregatorNamespace,
			Name:      "last_checkpoint_task_reference_aggregated",
			Help:      "Last checkpoint task reference aggregated",
		},
	)
	if err := registry.Register(lastCheckpointTaskReferenceAggregated); err != nil {
		return nil, fmt.Errorf("error registering lastCheckpointTaskReferenceAggregated gauge: %w", err)
	}

	return &SelectiveAggregatorListener{
		ObserveLastStateRootUpdateAggregatedCb: func(rollupId uint32, blockNumber uint64) {
			lastStateRootUpdateAggregated.WithLabelValues(fmt.Sprintf("%d", rollupId)).Set(float64(blockNumber))
		},
		ObserveLastStateRootUpdateReceivedCb: func(rollupId uint32, blockNumber uint64) {
			lastStateRootUpdateReceived.WithLabelValues(fmt.Sprintf("%d", rollupId)).Set(float64(blockNumber))
		},
		ObserveLastOperatorSetUpdateAggregatedCb: func(operatorSetUpdateId uint64) {
			lastOperatorSetUpdateAggregated.Set(float64(operatorSetUpdateId))
		},
		ObserveLastOperatorSetUpdateReceivedCb: func(operatorSetUpdateId uint64) {
			lastOperatorSetUpdateReceived.Set(float64(operatorSetUpdateId))
		},
		IncExpiredMessagesCb: func() {
			expiredMessages.Inc()
		},
		IncExpiredTasksCb: func() {
			expiredTasks.Inc()
		},
		IncErroredSubmissionsCb: func() {
			erroredSubmissions.Inc()
		},
		IncAggregatorInitializationsCb: func() {
			aggregatorInitializations.Inc()
		},
		ObserveLastCheckpointReferenceSentCb: func(referenceId uint32) {
			lastCheckpointReferenceSent.Set(float64(referenceId))
		},
		ObserveLastCheckpointTaskReferenceReceivedCb: func(referenceId uint32) {
			lastCheckpointTaskReferenceReceived.Set(float64(referenceId))
		},
		ObserveLastCheckpointTaskReferenceAggregatedCb: func(referenceId uint32) {
			lastCheckpointTaskReferenceAggregated.Set(float64(referenceId))
		},
	}, nil
}
