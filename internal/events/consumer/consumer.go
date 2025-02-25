// internal/events/consumer/consumer.go
package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(brokers []string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{consumer: consumer}, nil
}

// Listens to the topic and sends messages to the handler
func (kc *KafkaConsumer) StartConsumer(topic string, handler func([]byte) error) error {
	partitionConsumer, err := kc.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return fmt.Errorf("failed to start consumer: %w", err)
	}
	defer partitionConsumer.Close()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				err := handler(msg.Value)
				if err != nil {
					fmt.Printf("Error handling message: %v\n", err)
				}
			case err := <-partitionConsumer.Errors():
				fmt.Printf("Consumer error: %v\n", err)
			case <-signChan:
				fmt.Println("Interrupt detected")
				doneCh <- struct{}{}
				return
			}
		}
	}()

	<-doneCh
	return nil
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
