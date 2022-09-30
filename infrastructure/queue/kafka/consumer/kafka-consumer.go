package consumer

import (
	"context"
	"os"
	kafkaclient "strategy-fulfillment-flow/infrastructure/queue/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetConsumerClient(version string, topic string) ConsumerInterface {
	config := kafkaclient.GetKafkConfigByVersion(version)
	config.SetKey("group.id", os.Getenv("CONSUMER_GROUP_ID"))
	config.SetKey("auto.offset.reset", "earliest")

	return &Client{
		topic:  topic,
		config: config,
	}
}

func (client *Client) SuscribeTopic(ctx context.Context) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&client.config)
	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{client.topic}, nil)
	return consumer
}

func (client *Client) Close() error {
	return client.consumer.Close()
}
