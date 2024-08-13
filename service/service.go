package service

import (
	"api-gateway/config"
	"api-gateway/generated/budgeting"
	"api-gateway/generated/user"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	FinanceManagement() budgeting.FinanceManagementServiceClient
	BudgetManagement() budgeting.BudgetingServiceClient
	GoalsManagement() budgeting.GoalsManagemenServiceClient
	User() user.AuthServiceClient
}

type serviceManagerImpl struct {
	financeManagement budgeting.FinanceManagementServiceClient
	budgetManagement  budgeting.BudgetingServiceClient
	goalsManagement   budgeting.GoalsManagemenServiceClient
	user              user.AuthServiceClient
}

func NewServiceManager(cfg *config.Config) (ServiceManager, error) {
	connBudget, err := grpc.NewClient(
		fmt.Sprintf("localhost:%d", cfg.BUDGETING_SERVICE_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Println("error connecting to gRPC server: ", err)
		return nil, err
	}
	
	connUser, err := grpc.NewClient(
		fmt.Sprintf("localhost:%d", cfg.USER_SERVICE_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("error connecting to gRPC server: ", err)
		return nil, err
	}

	return &serviceManagerImpl{
		financeManagement: budgeting.NewFinanceManagementServiceClient(connBudget),
		budgetManagement:  budgeting.NewBudgetingServiceClient(connBudget),
		goalsManagement:   budgeting.NewGoalsManagemenServiceClient(connBudget),
		user:              user.NewAuthServiceClient(connUser),
	}, nil
}

func (s *serviceManagerImpl) FinanceManagement() budgeting.FinanceManagementServiceClient {
	return s.financeManagement
}

func (s *serviceManagerImpl) BudgetManagement() budgeting.BudgetingServiceClient {
	return s.budgetManagement
}

func (s *serviceManagerImpl) GoalsManagement() budgeting.GoalsManagemenServiceClient {
	return s.goalsManagement
}

func (s *serviceManagerImpl) User() user.AuthServiceClient {
	return s.user
}
