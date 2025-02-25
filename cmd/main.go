package main

import (
	"eda/internal/events/consumer"
	"eda/internal/events/handlers"
	"eda/internal/events/producer"
	"eda/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	brokers = []string{"localhost:29092"}
)

func main() {
	kafkaProducer, err := producer.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer func() {
		if err := kafkaProducer.Close(); err != nil {
			log.Printf("Failed to close Kafka producer: %v", err)
		}
	}()

	kafkaConsumer, err := consumer.NewKafkaConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	err = kafkaConsumer.StartConsumer("comments", handlers.LogHanlder)
	if err != nil {
		log.Fatalf("Error consuming messages: %v", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, kafkaProducer)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
