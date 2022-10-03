package main

import (
	"fmt"
	"os"
	"strategy-fulfillment-flow/metrics"
	"strategy-fulfillment-flow/strategy"
)

var execute_strategy strategy.Strategy

func main() {

	switch os.Getenv("STRATEGY") {
	case "RECEIVE_ORDERS":
		execute_strategy = strategy.ReceptionOrderStrategy{}
	case "PICKING_FINALIZED":
		execute_strategy = strategy.ReceptionOrderStrategy{}
	default:
		fmt.Println("STRATEGY DOES NO DEFINED")
		panic(1)
	}

	eventMetrics := metrics.InitMetrics()
	go strategy.InitStrategy(
		execute_strategy,
		os.Getenv("INPUT_QUEUE_VERSION"),
		os.Getenv("OUTPUT_QUEUE_VERSION"),
		os.Getenv("INPUT_TOPIC"),
		eventMetrics).Execute()

	metrics.Expose()
}
