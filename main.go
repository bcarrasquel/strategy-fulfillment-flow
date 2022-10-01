package main

import (
	kafkaclient "strategy-fulfillment-flow/infrastructure/queue/kafka"
	"strategy-fulfillment-flow/metrics"
	"strategy-fulfillment-flow/strategy"
)

func main() {
	eventMetrics := metrics.InitMetrics()

	go strategy.InitStrategy(
		strategy.ReceptionOrderStrategy{},
		kafkaclient.KAFKA_VERSION_2_8,
		kafkaclient.KAFKA_VERSION_2_8,
		"create_order_input_topic",
		eventMetrics).Execute()

	go strategy.InitStrategy(
		strategy.PickingFinalizedStrategy{},
		kafkaclient.KAFKA_VERSION_2_8,
		kafkaclient.KAFKA_VERSION_2_8,
		"picking_finalized_input_topic",
		eventMetrics).Execute()

	metrics.Expose()
}
