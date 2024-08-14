package handler

import (
	bhandler "api-gateway/api/handler/budgeting"
	uhandler "api-gateway/api/handler/user"
	"api-gateway/queue/kafka/producer"
	"api-gateway/service"
	"log/slog"
)

type MainHandler interface {
	UserHandler() uhandler.UserHandler
	BudgetHandler() bhandler.BudgettingHandler
}

type mainHandlerImpl struct {
	service       service.ServiceManager
	logger        *slog.Logger
	kafkaProducer producer.KafkaProducer
	// rabbitmqProducer producermq.RabbitMQProducer
}

func NewMainHandler(service service.ServiceManager, logger *slog.Logger, kafkaProducer producer.KafkaProducer) MainHandler {
	return &mainHandlerImpl{
		service:       service,
		logger:        logger,
		kafkaProducer: kafkaProducer,
		// rabbitmqProducer: rabbitmqProducer,
	}
}

func (mh *mainHandlerImpl) UserHandler() uhandler.UserHandler {
	return uhandler.NewUserHandler(mh.service.User(), mh.logger)
}

func (mh *mainHandlerImpl) BudgetHandler() bhandler.BudgettingHandler {
	return bhandler.NewBudgettingHandler(mh.service, mh.logger, mh.kafkaProducer)
}
