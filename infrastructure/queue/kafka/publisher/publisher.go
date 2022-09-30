package publisher

import (
	"fmt"
	kafkaclient "strategy-fulfillment-flow/infrastructure/queue/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var newProducer func(*kafka.ConfigMap) (*kafka.Producer, error) = kafka.NewProducer

func GetProducerClient(version string) PublisherInterface {
	config := kafkaclient.GetKafkConfigByVersion(version)

	producer, err := newProducer(&config)
	if err != nil {
		fmt.Println("Failed to create kafka producer")
		panic(err)
	}

	return &Client{
		producer: producer,
	}
}

func (client *Client) Produce(msg []byte, topic string) error {
	producerError := client.producer.Produce(createMessage(msg, topic), nil)
	if producerError != nil {
		fmt.Println("Failed to publish event " + producerError.Error())
		panic(producerError)
	}
	event := <-client.producer.Events()
	switch ev := event.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			fmt.Println("Failed to deliver message")
		} else {
			fmt.Println("Successfully produced record to topic")
		}
	}
	return nil
}

func createMessage(message []byte, topic string) *kafka.Message {
	result := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}
	return result
}
