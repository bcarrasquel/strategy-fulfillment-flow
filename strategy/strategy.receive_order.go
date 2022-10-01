package strategy

import (
	"strategy-fulfillment-flow/adapters"
)

type ReceptionOrderStrategy struct{}

func (strategy ReceptionOrderStrategy) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	return adapters.AdaptOrderReception(message)
}
