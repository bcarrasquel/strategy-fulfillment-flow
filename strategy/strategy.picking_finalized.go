package strategy

import (
	"fmt"
)

type PickingFinalizedStrategy struct{}

func (strategy PickingFinalizedStrategy) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	fmt.Println("Procesando el picking finalizado" + string(message))
	return nil, "", nil
}
