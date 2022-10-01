package strategy

import (
	"encoding/json"
	"errors"
	"os"
	"strategy-fulfillment-flow/logger"
)

type ReceptionOrderStrategy struct{}

type Order struct {
	WalmartID    string `json:"walmartId"`
	OrderID      string `json:"orderId"`
	BusinessType string `json:"businessType"`
	SalesChannel string `json:"salesChannel"`
}

const (
	ORDER_SOD      = "GR"
	ORDER_CATEX    = "GM"
	EVET_FLOW_NAME = "create_order_flow"
)

var sod_topic = os.Getenv("TOPIC_ORDER_SOD_GENERATED")
var catex_topic = os.Getenv("TOPIC_ORDER_CATEX_GENERATED")

func (s ReceptionOrderStrategy) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	var order Order
	if err := json.Unmarshal(message, &order); err != nil {
		log.Error("Invalid interface for create new fulfillment order", logger.Parameters{Error: err.Error()})
		return nil, "", errors.New("invalid interface for create new fulfillment order")
	}

	if order.BusinessType == ORDER_SOD {
		return message, sod_topic, nil
	} else if order.BusinessType == ORDER_CATEX {
		return message, catex_topic, nil
	}

	log.Info("Event does not match any output topic", logger.Parameters{})
	return nil, "", errors.New("event does not match any output topic")
}
