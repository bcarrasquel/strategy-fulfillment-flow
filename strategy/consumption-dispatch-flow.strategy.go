package strategy

import (
	"context"
	"strategy-fulfillment-flow/infrastructure/queue/kafka/consumer"
	"strategy-fulfillment-flow/infrastructure/queue/kafka/publisher"
	"strategy-fulfillment-flow/logger"
	"strategy-fulfillment-flow/metrics"
)

var log = logger.NewLogger()

func (flow *flow) ExecuteFlow(eventMetrics map[string]metrics.CounterInterface) {

	config := flow.FlowProcess.SetConfig()
	producer := getProducer(config["queue_output_version"])

	var consumerInterface consumer.ConsumerInterface = consumer.GetConsumerClient(config["queue_input_version"], config["input_topic"])
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

		output_message, output_topic, err := flow.FlowProcess.BusinessAdaptersExecute(message.Value)
		if err != nil {
			producer.Produce(output_message, output_topic)
			metrics.Inc(eventMetrics["publish"], output_topic, "")
		}
	}
}

func getProducer(version string) publisher.PublisherInterface {
	var publisherInterface publisher.PublisherInterface = publisher.GetProducerClient(version)
	return publisherInterface
}
