package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/config"
	"api-gateway/queue/kafka/producer"
	"api-gateway/queue/rabbitmq/producermq"
	"api-gateway/service"
	"fmt"
	"log/slog"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	Start(cfg *config.Config) error
	SetupRoutes(service service.ServiceManager, kafkaProducer producer.KafkaProducer, rabbitmqProducer producermq.RabbitMQProducer, logger *slog.Logger)
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
		c.port = fmt.Sprintf(":%d", cfg.HTTP_PORT)
	}

	return c.router.Run(c.port)
}

// @title Api Gateway
// @version 1.0
// @description Api Gateway service
// @BasePath /api/v1
// @schemes http
// @in header
// @name Authorization
func (c *controllerImpl) SetupRoutes(service service.ServiceManager, kafkaProducer producer.KafkaProducer, rabbitmqProducer producermq.RabbitMQProducer, logger *slog.Logger) {

	h := handler.NewMainHandler(
		service,
		logger,
		kafkaProducer,
		rabbitmqProducer,
	)

	c.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c.router.Use(middleware.IsAuthenticated(service.User()))
	router := c.router.Group("/api/v1")

	// User routes
	user := router.Group("/users")
	{
		user.GET("/profile", h.UserHandler().GetUserProfileHandler)
		user.PUT("/profile", h.UserHandler().UpdateUserProfileHandler)
		user.PUT("/password", h.UserHandler().ChangePasswordHandler)
		user.GET("/", h.UserHandler().GetUsersListHandler)
	}

	account := router.Group("/accounts")
	{
		account.POST("/", h.BudgetHandler().CreateAccountHandler)
		account.GET("/:id", h.BudgetHandler().GetAccountHandler)
		account.PUT("/:id", h.BudgetHandler().UpdateAccountHandler)
		account.DELETE("/:id", h.BudgetHandler().DeleteAccountHandler)
		account.GET("/", h.BudgetHandler().GetAccountsListHandler)
	}

	budget := router.Group("/budgets")
	{
		budget.POST("/", h.BudgetHandler().CreateBudgetHandler)
		budget.GET("/:id", h.BudgetHandler().GetBudgetHandler)
		budget.PUT("/:id", h.BudgetHandler().UpdateBudgetHandler)
		budget.DELETE("/:id", h.BudgetHandler().DeleteBudgetHandler)
		budget.GET("/", h.BudgetHandler().GetBudgetsListHandler)
	}

	category := router.Group("/categories")
	{
		category.POST("/", h.BudgetHandler().CreateCategoryHandler)
		category.PUT("/:id", h.BudgetHandler().UpdateCategoryHandler)
		category.DELETE("/:id", h.BudgetHandler().DeleteCategoryHandler)
		category.GET("/", h.BudgetHandler().GetCategoriesListHandler)
		category.GET("/:id", h.BudgetHandler().GetCategoryHandler)
	}

	transaction := router.Group("/transactions")
	{
		transaction.POST("/", h.BudgetHandler().CreateTransactionHandler)
		transaction.GET("/:id", h.BudgetHandler().GetTransactionHandler)
		transaction.PUT("/:id", h.BudgetHandler().UpdateTransactionHandler)
		transaction.DELETE("/:id", h.BudgetHandler().DeleteTransactionHandler)
		transaction.GET("/", h.BudgetHandler().GetTransactionsListHandler)
	}

	goal := router.Group("/goals")
	{
		goal.POST("/", h.BudgetHandler().CreateGoalsHandler)
		goal.GET("/:id", h.BudgetHandler().GetGoalHandler)
		goal.PUT("/:id", h.BudgetHandler().UpdateGoalHandler)
		goal.DELETE("/:id", h.BudgetHandler().DeleteGoalHandler)
		goal.GET("/", h.BudgetHandler().GetGoalsListHandler)
	}

	reporting := router.Group("/reporting")
	{
		reporting.GET("/sepending", h.BudgetHandler().GetSependingReportHandler)
        reporting.GET("/income", h.BudgetHandler().GetIncomeReportHandler)
        reporting.GET("/budget-performance", h.BudgetHandler().GetBudgetPerformanceHandler)
        reporting.GET("/goal-progress", h.BudgetHandler().GetGoalProgressHandler)
	}
}
