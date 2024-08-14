package main

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/pkg/logs"
	"api-gateway/queue/kafka/producer"
	"api-gateway/service"
	"log"
)

func main() {
	log.Println("Starting server")
	logger := logs.InitLogger()
	cfg := config.Load()

	kafkaProducer := producer.NewKafkaProducer(cfg.KafkaBrokers)
	defer kafkaProducer.Close()

	// rabbitMQProducer, err := producermq.NewRabbitMQProducer("amqp://guest:guest@localhost:5672/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rabbitMQProducer.Close()
	enforcer, err := config.CasbinEnforcer()
	if err != nil {
		log.Println("error initializing Casbin enforcer: ", err)
        log.Fatal(err)
    }


	serviceManager, err := service.NewServiceManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	controller := api.NewController()
	controller.SetupRoutes(serviceManager, kafkaProducer, enforcer, logger)
	err = controller.Start(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
