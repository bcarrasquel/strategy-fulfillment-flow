package strategy

import (
	"context"
	"strategy-fulfillment-flow/infrastructure/queue/kafka/consumer"
	"strategy-fulfillment-flow/infrastructure/queue/kafka/publisher"
	"strategy-fulfillment-flow/logger"
	"strategy-fulfillment-flow/metrics"
)

var log = logger.NewLogger()

func InitStrategy(s Strategy, input_queue_version string, output_queue_version string, input_topic string) *strategy {
	return &strategy{
		Strategy:             s,
		input_queue_version:  input_queue_version,
		output_queue_version: output_queue_version,
		input_topic:          input_topic,
	}
}

func (s *strategy) Execute(eventMetrics map[string]metrics.CounterInterface) {
	var producer publisher.PublisherInterface = publisher.GetProducerClient(s.output_queue_version)
	var consumerInterface consumer.ConsumerInterface = consumer.GetConsumerClient(s.input_queue_version, s.input_topic)
	consumer := consumerInterface.SuscribeTopic(context.Background())
	defer consumer.Close()

	for {
		message, errorConsume := consumer.ReadMessage(-1)
		if errorConsume != nil {
			metrics.Inc(eventMetrics["error"], *message.TopicPartition.Topic, "consumer-error")
			log.Error("Consumer error", logger.Parameters{Error: errorConsume.Error()})
			panic(errorConsume)
		}

		metrics.Inc(eventMetrics["read"], *message.TopicPartition.Topic, "")
		_, errorCommit := consumer.Commit()

		if errorCommit != nil {
			metrics.Inc(eventMetrics["error"], *message.TopicPartition.Topic, "consumer-commit-error")
			log.Error("Error committing message", logger.Parameters{Error: errorCommit.Error()})
		} else {
			metrics.Inc(eventMetrics["process"], *message.TopicPartition.Topic, "")
		}

		output_message, output_topic, err := s.Strategy.BusinessAdaptersExecute(message.Value)
		if err != nil {
			producer.Produce(output_message, output_topic)
			metrics.Inc(eventMetrics["publish"], output_topic, "")
		}
	}
}
