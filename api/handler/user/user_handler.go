package uhandler

import (
	"api-gateway/generated/user"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserProfileHandler(c *gin.Context)
	UpdateUserProfileHandler(c *gin.Context)
	ChangePasswordHandler(c *gin.Context)
	GetUsersListHandler(c *gin.Context)
}

type userHandlerImpl struct {
	userClient user.AuthServiceClient
	logger     *slog.Logger
}

func NewUserHandler(userClient user.AuthServiceClient, logger *slog.Logger) UserHandler {
	return &userHandlerImpl{
		userClient: userClient,
		logger:     logger,
	}
}
