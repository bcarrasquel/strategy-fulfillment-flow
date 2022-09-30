package consumer

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Client struct {
	consumer *kafka.Consumer
	topic    string
	config   kafka.ConfigMap
}
