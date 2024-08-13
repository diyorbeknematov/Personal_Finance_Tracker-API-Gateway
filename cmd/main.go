package main

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/pkg/logs"
	"api-gateway/queue/kafka/producer"
	"api-gateway/queue/rabbitmq/producermq"
	"api-gateway/service"
	"log"
)

func main() {
	log.Println("Starting server")
	logger := logs.InitLogger()
	cfg := config.Load()

	kafkaProducer := producer.NewKafkaProducer([]string{"localhost:9092"})
	defer kafkaProducer.Close()

	rabbitMQProducer, err := producermq.NewRabbitMQProducer("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitMQProducer.Close()

	serviceManager, err := service.NewServiceManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	controller := api.NewController()
	controller.SetupRoutes(serviceManager, kafkaProducer, rabbitMQProducer, logger)
	err = controller.Start(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
