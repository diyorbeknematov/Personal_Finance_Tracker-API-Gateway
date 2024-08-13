package uhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/user"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Profile
// @Description This API is used to get user profile
// @tags user
// @Accept json
// @Produce json
// @Success 200 {object} user.UserProfile
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /user/profile [get]
func (h *userHandlerImpl) GetUserProfileHandler(ctx *gin.Context) {
	val, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(400, gin.H{"error": "Missing token claims"})
		return
	}

	claims, err := token.TokenClaimsParse(val)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userClient.GetUserProfile(ctx, &pb.GetUserProfileReq{Id: claims.GetId()})
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Update User Profile
// @Description This API is used to update user profile
// @tags user
// @Accept json
// @Produce json
// @Param userProfile body user.UpdateUserProfileReq true "User Profile"
// @Success 200 {object} user.UpdateUserProfileResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /user/profile [put]
func (h *userHandlerImpl) UpdateUserProfileHandler(ctx *gin.Context) {
	var userProfile pb.UpdateUserProfileReq
	if err := ctx.ShouldBindJSON(&userProfile); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.userClient.UpdateUserProfile(ctx, &userProfile)
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Change Password
// @Description This API is used to change password
// @tags user
// @Accept json
// @Produce json
// @param changePassword body user.ChangePasswordReq true "Change Password"
// @Success 200 {object} user.ChangePasswordResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404	{object} map[string]interface{}
// @router /user/password [put]
func (h *userHandlerImpl) ChangePasswordHandler(ctx *gin.Context) {
	var changePassword pb.ChangePasswordReq
	if err := ctx.ShouldBindJSON(&changePassword); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.userClient.ChangePassword(ctx, &changePassword)
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @Summary Get Users List
// @Description This API is used to get users list
// @tags user
// @Accept json
// @Produce json
// @param Limit query int false "Limit" default(10)
// @param Page query int false "Page" default(1)
// @param FirstName query string false "FirstName"
// @param LastName query string false "LastName"
// @param Email query string false "Email"
// @param Role query string false "Role"
// @success 200 {object} user.GetUsersListResp
// @failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /users [get]
func (h *userHandlerImpl) GetUsersListHandler(ctx *gin.Context) {
	var req pb.GetUsersListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request query parameters",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.userClient.GetUsersList(ctx, &req)
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}
