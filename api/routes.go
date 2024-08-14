package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/config"
	"api-gateway/queue/kafka/producer"
	"api-gateway/service"
	"fmt"
	"log/slog"

	_ "api-gateway/api/docs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	Start(cfg *config.Config) error
	SetupRoutes(service.ServiceManager, producer.KafkaProducer, *casbin.Enforcer, *slog.Logger)
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
		c.port = fmt.Sprintf("api_app:%d", cfg.HTTP_PORT)
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
func (c *controllerImpl) SetupRoutes(service service.ServiceManager, kafkaProducer producer.KafkaProducer, enforcer *casbin.Enforcer, logger *slog.Logger) {

	h := handler.NewMainHandler(
		service,
		logger,
		kafkaProducer,
		// rabbitmqProducer,
	)

	c.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := c.router.Group("/api/v1")
	// c.router.Use(middleware.IsAuthenticated(service.User()))

	// User routes

	user := router.Group("/users")
	user.Use(middleware.IsAuthenticated(service.User()))
	user.Use(middleware.IsAuthorize(enforcer))

	{
		user.GET("/profile", h.UserHandler().GetUserProfileHandler)
		user.PUT("/profile", h.UserHandler().UpdateUserProfileHandler)
		user.PUT("/password", h.UserHandler().ChangePasswordHandler)
		user.GET("/", h.UserHandler().GetUsersListHandler)
	}

	account := router.Group("/accounts")
	account.Use(middleware.IsAuthenticated(service.User()))
	account.Use(middleware.IsAuthorize(enforcer))
	{
		account.POST("/", h.BudgetHandler().CreateAccountHandler)
		account.GET("/:id", h.BudgetHandler().GetAccountHandler)
		account.PUT("/:id", h.BudgetHandler().UpdateAccountHandler)
		account.DELETE("/:id", h.BudgetHandler().DeleteAccountHandler)
		account.GET("/", h.BudgetHandler().GetAccountsListHandler)
	}

	budget := router.Group("/budgets")
	budget.Use(middleware.IsAuthenticated(service.User()))
	budget.Use(middleware.IsAuthorize(enforcer))
	{
		budget.POST("/", h.BudgetHandler().CreateBudgetHandler)
		budget.GET("/:id", h.BudgetHandler().GetBudgetHandler)
		budget.PUT("/:id", h.BudgetHandler().UpdateBudgetHandler)
		budget.DELETE("/:id", h.BudgetHandler().DeleteBudgetHandler)
		budget.GET("/", h.BudgetHandler().GetBudgetsListHandler)
	}

	category := router.Group("/categories")
	category.Use(middleware.IsAuthenticated(service.User()))
	category.Use(middleware.IsAuthorize(enforcer))
	{
		category.POST("/", h.BudgetHandler().CreateCategoryHandler)
		category.PUT("/:id", h.BudgetHandler().UpdateCategoryHandler)
		category.DELETE("/:id", h.BudgetHandler().DeleteCategoryHandler)
		category.GET("/", h.BudgetHandler().GetCategoriesListHandler)
		category.GET("/:id", h.BudgetHandler().GetCategoryHandler)
	}

	transaction := router.Group("/transactions")
	transaction.Use(middleware.IsAuthenticated(service.User()))
	transaction.Use(middleware.IsAuthorize(enforcer))
	{
		transaction.POST("/", h.BudgetHandler().CreateTransactionHandler)
		transaction.GET("/:id", h.BudgetHandler().GetTransactionHandler)
		transaction.PUT("/:id", h.BudgetHandler().UpdateTransactionHandler)
		transaction.DELETE("/:id", h.BudgetHandler().DeleteTransactionHandler)
		transaction.GET("/", h.BudgetHandler().GetTransactionsListHandler)
	}

	goal := router.Group("/goals")
	goal.Use(middleware.IsAuthenticated(service.User()))
	goal.Use(middleware.IsAuthorize(enforcer))
	{
		goal.POST("/", h.BudgetHandler().CreateGoalsHandler)
		goal.GET("/:id", h.BudgetHandler().GetGoalHandler)
		goal.PUT("/:id", h.BudgetHandler().UpdateGoalHandler)
		goal.DELETE("/:id", h.BudgetHandler().DeleteGoalHandler)
		goal.GET("/", h.BudgetHandler().GetGoalsListHandler)
	}

	reporting := router.Group("/reporting")
	reporting.Use(middleware.IsAuthenticated(service.User()))
	reporting.Use(middleware.IsAuthorize(enforcer))
	{
		reporting.GET("/sepending", h.BudgetHandler().GetSependingReportHandler)
		reporting.GET("/income", h.BudgetHandler().GetIncomeReportHandler)
		reporting.GET("/budget-performance", h.BudgetHandler().GetBudgetPerformanceHandler)
		reporting.GET("/goal-progress", h.BudgetHandler().GetGoalProgressHandler)
	}

	notification := router.Group("/notification")
	notification.Use(middleware.IsAuthenticated(service.User()))
	notification.Use(middleware.IsAuthorize(enforcer))
	{
		notification.POST("/", h.BudgetHandler().SendNotificationHandler)
		notification.GET("/", h.BudgetHandler().GetNotificationListHandler)
		notification.GET("/:id", h.BudgetHandler().GetNotificationHandler)
		notification.PUT("/:id/read", h.BudgetHandler().UpdateNotificationHandler)
		notification.DELETE("/:id", h.BudgetHandler().DeleteNotificationHandler)
	}
}
