package strategy

import "strategy-fulfillment-flow/metrics"

type strategy struct {
	Strategy             Strategy
	input_queue_version  string
	output_queue_version string
	input_topic          string
	metrics              map[string]metrics.CounterInterface
}

type Strategy interface {
	BusinessAdaptersExecute(message []byte) ([]byte, string, error)
}
