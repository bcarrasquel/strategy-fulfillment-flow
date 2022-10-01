package strategy

type PickingFinalizedStrategy struct{}

func (strategy PickingFinalizedStrategy) BusinessAdaptersExecute(message []byte) ([]byte, string, error) {
	return nil, "", nil
}
