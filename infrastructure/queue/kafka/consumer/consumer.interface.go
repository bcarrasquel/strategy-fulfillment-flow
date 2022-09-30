package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerInterface interface {
	SuscribeTopic(ctx context.Context) *kafka.Consumer
	Close() error
}
