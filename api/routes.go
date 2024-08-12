package api

import (
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/queue/kafka/producer"
	"api-gateway/queue/rabbitmq/producermq"
	"api-gateway/service"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
}

type controllerImpl struct {
	port   string
	router *gin.Engine
}

func NewController() Controller {
	router := gin.Default()

	return &controllerImpl{
		router: router,
	}
}

func (c *controllerImpl) Start(cfg *config.Config) error {
	if c.port == "" {
		c.port = fmt.Sprintf("localhost:%d", cfg.HTTP_PORT)
	}

	return c.router.Run(c.port)
}

// @Title API Gateway
// @Summary This is a simple API Gateway
// @Description This is a simple API Gateway
// @version 1.0
// @in header
// @name Authorization
// @BasePath /api/v1
// @schemes http
func (c *controllerImpl) SetupAPIRoutes(service service.ServiceManager, kafkaProducer producer.KafkaProducer, rabbitmqProducer producermq.RabbitMQProducer, logger *slog.Logger) {

	h := handler.NewMainHandler(
		service,
		logger,
		kafkaProducer,
		rabbitmqProducer,
	)

	c.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := c.router.Group("/api/v1")

	// User routes
	user := router.Group("/user")
	{
		user.GET("/profile", h.UserHandler().GetUserProfileHandler)
		user.PUT("/profile", h.UserHandler().UpdateUserProfileHandler)
		user.PUT("/password", h.UserHandler().ChangePasswordHandler)
		user.GET("/users", h.UserHandler().GetUsersListHandler)
	}
}
