package producer

import (
	"github.com/IBM/sarama"
)

type EventProducer interface {
	Publish(topic string, message []byte) error
	Close() error
}

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (EventProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer: producer}, nil
}

// Publish sends a message to a Kafka topic
func (k *KafkaProducer) Publish(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := k.producer.SendMessage(msg)
	return err
}

// Shuts down the Kafka producer
func (k *KafkaProducer) Close() error {
	return k.producer.Close()
}
