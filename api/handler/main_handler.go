package handler

import (
	uhandler "api-gateway/api/handler/user"
	"api-gateway/queue/kafka/producer"
	"api-gateway/queue/rabbitmq/producermq"
	"api-gateway/service"
	"log/slog"
)

type MainHandler interface {
	UserHandler() uhandler.UserHandler
}

type mainHandlerImpl struct {
	service          service.ServiceManager
	logger           *slog.Logger
	kafkaProducer    producer.KafkaProducer
	rabbitmqProducer producermq.RabbitMQProducer
}

func NewMainHandler(service service.ServiceManager, logger *slog.Logger, kafkaProducer producer.KafkaProducer, rabbitmqProducer producermq.RabbitMQProducer) MainHandler {
	return &mainHandlerImpl{
		service:          service,
		logger:           logger,
		kafkaProducer:    kafkaProducer,
		rabbitmqProducer: rabbitmqProducer,
	}
}

func (mh *mainHandlerImpl) UserHandler() uhandler.UserHandler {
	return uhandler.NewUserHandler(mh.service.User(), mh.logger)
}
