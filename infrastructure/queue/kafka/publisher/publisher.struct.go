package publisher

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Client struct {
	producer *kafka.Producer
}
