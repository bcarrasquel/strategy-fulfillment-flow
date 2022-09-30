package strategy

import (
	"strategy-fulfillment-flow/metrics"
)

func Init() {
	metrics := metrics.InitMetrics()

	go newConsumptionAndDispatchQueueFlow(ReceptionOrderFlow{}).ExecuteFlow(metrics)
	go newConsumptionAndDispatchQueueFlow(PickingFinalizedFlow{}).ExecuteFlow(metrics)
}

func newConsumptionAndDispatchQueueFlow(flowProcess FlowProcess) *flow {
	return &flow{
		FlowProcess: flowProcess,
	}
}
