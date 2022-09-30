package publisher

type PublisherInterface interface {
	Produce(msg []byte, topic string) error
}
