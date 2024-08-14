package bhandler

import (
	"api-gateway/generated/budgeting"
	"api-gateway/queue/kafka/producer"
	"api-gateway/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type BudgettingHandler interface {
	CreateAccountHandler(ctx *gin.Context)
	GetAccountHandler(ctx *gin.Context)
	UpdateAccountHandler(ctx *gin.Context)
	DeleteAccountHandler(ctx *gin.Context)
	GetAccountsListHandler(ctx *gin.Context)

	// budget methods
	CreateBudgetHandler(ctx *gin.Context)
	GetBudgetsListHandler(ctx *gin.Context)
	UpdateBudgetHandler(ctx *gin.Context)
	DeleteBudgetHandler(ctx *gin.Context)
	GetBudgetHandler(ctx *gin.Context)

	// Category
	CreateCategoryHandler(ctx *gin.Context)
	UpdateCategoryHandler(ctx *gin.Context)
	DeleteCategoryHandler(ctx *gin.Context)
	GetCategoriesListHandler(ctx *gin.Context)
	GetCategoryHandler(ctx *gin.Context)

	// Goal
	CreateGoalsHandler(ctx *gin.Context)
	UpdateGoalHandler(ctx *gin.Context)
	DeleteGoalHandler(ctx *gin.Context)
	GetGoalsListHandler(ctx *gin.Context)
	GetGoalHandler(ctx *gin.Context)

	// Transaction
	CreateTransactionHandler(ctx *gin.Context)
	UpdateTransactionHandler(ctx *gin.Context)
	DeleteTransactionHandler(ctx *gin.Context)
	GetTransactionsListHandler(ctx *gin.Context)
	GetTransactionHandler(ctx *gin.Context)

	// Reporting
	GetSependingReportHandler(ctx *gin.Context)
	GetBudgetPerformanceHandler(ctx *gin.Context)
	GetGoalProgressHandler(ctx *gin.Context)
	GetIncomeReportHandler(ctx *gin.Context)

	// Notification
	SendNotificationHandler(ctx *gin.Context)
	GetNotificationListHandler(ctx *gin.Context)
	GetNotificationHandler(ctx *gin.Context)
	UpdateNotificationHandler(ctx *gin.Context)
	DeleteNotificationHandler(ctx *gin.Context)
}

type budgettingHandlerImpl struct {
	financeManagement   budgeting.FinanceManagementServiceClient
	budgetManagement    budgeting.BudgetingServiceClient
	goalsManagement     budgeting.GoalsManagemenServiceClient
	reportingManagement budgeting.ReportingNotificationServiceClient
	producer            producer.KafkaProducer
	logger              *slog.Logger
}

func NewBudgettingHandler(serviceManger service.ServiceManager, logger *slog.Logger, producer producer.KafkaProducer) BudgettingHandler {
	return &budgettingHandlerImpl{
		financeManagement:   serviceManger.FinanceManagement(),
		budgetManagement:    serviceManger.BudgetManagement(),
		goalsManagement:     serviceManger.GoalsManagement(),
		reportingManagement: serviceManger.ReportingManagement(),
		producer:            producer,
		logger:              logger,
	}
}
