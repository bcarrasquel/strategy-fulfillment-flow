package main

import (
	"strategy-fulfillment-flow/metrics"
	"strategy-fulfillment-flow/strategy"
)

func main() {
	strategy.Init()
	metrics.Expose()
}
