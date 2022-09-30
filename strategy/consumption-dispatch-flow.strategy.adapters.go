package strategy

import (
	"fmt"
	"strategy-fulfillment-flow/adapters"
	kafkaclient "strategy-fulfillment-flow/infrastructure/queue/kafka"
)

// ------------ RECEPTION ORDER FLOW ADAPTERS ------------

func (flow ReceptionOrderFlow) SetConfig() map[string]string {
	config := make(map[string]string)
	config["input_topic"] = "create_order_input_topic"
	config["queue_input_version"] = kafkaclient.KAFKA_VERSION_2_8
	config["queue_output_version"] = kafkaclient.KAFKA_VERSION_2_8
	return config
}

func (flow ReceptionOrderFlow) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	return adapters.AdaptOrderReception(message)
}

// ------------ PICKING FINALIZED FLOW ADAPTERS ------------

func (flow PickingFinalizedFlow) SetConfig() map[string]string {
	config := make(map[string]string)
	config["input_topic"] = "picking_finalized_input_topic"
	config["queue_input_version"] = kafkaclient.KAFKA_VERSION_2_8
	config["queue_output_version"] = kafkaclient.KAFKA_VERSION_2_8

	return config
}

func (flow PickingFinalizedFlow) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	fmt.Println("Procesando el picking finalizado" + string(message))
	return nil, "", nil
}
