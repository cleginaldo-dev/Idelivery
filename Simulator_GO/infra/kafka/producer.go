package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.server": os.Getenv("KafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Panicln(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Producer(message, nil)
	if err != nil {
		return err
	}
	return nil
}
