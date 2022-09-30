package strategy

import (
	"fmt"
	kafkaclient "strategy-fulfillment-flow/infrastructure/queue/kafka"
)

// ------------ RECEPTION ORDER FLOW ADAPTERS ------------

func (flow ReceptionOrderFlow) SetConfig() map[string]string {
	config := make(map[string]string)
	config["input_topic"] = "create_order_input_topic"
	config["queue_input_version"] = kafkaclient.KAFKA_VERSION_2_8

	return config
}

func (flow ReceptionOrderFlow) BusinessAdaptersExecute(message []byte) bool {
	fmt.Println("Procesando el picking finalizado" + string(message))
	return true
}

// ------------ PICKING FINALIZED FLOW ADAPTERS ------------

func (flow PickingFinalizedFlow) SetConfig() map[string]string {
	config := make(map[string]string)
	config["input_topic"] = "picking_finalized_input_topic"
	config["queue_input_version"] = kafkaclient.KAFKA_VERSION_2_8

	return config
}

func (flow PickingFinalizedFlow) BusinessAdaptersExecute(message []byte) bool {
	fmt.Println("Procesando el picking finalizado" + string(message))
	return true
}
