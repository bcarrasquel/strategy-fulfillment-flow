package strategy

type FlowProcess interface {
	SetConfig() map[string]string
	BusinessAdaptersExecute(message []byte) bool
}
