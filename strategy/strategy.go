package strategy

type strategy struct {
	Strategy             Strategy
	input_queue_version  string
	output_queue_version string
	input_topic          string
}

type Strategy interface {
	BusinessAdaptersExecute(message []byte) ([]byte, string, error)
}
