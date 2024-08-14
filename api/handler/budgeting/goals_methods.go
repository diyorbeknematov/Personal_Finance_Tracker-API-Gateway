package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @summary Create a new goal
// @description Creates a new goal
// @tags goal
// @accept json
// @produce json
// @param goal body budgeting.CreateGoalReq true "New goal"
// @success 200 {object} budgeting.CreateGoalResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /goals [post]
func (h *budgettingHandlerImpl) CreateGoalsHandler(ctx *gin.Context) {
	var goal pb.CreateGoalReq
	if err := ctx.ShouldBindJSON(&goal); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.goalsManagement.CreateGoal(ctx, &goal)
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

// @summary Get a goal by id
// @description Gets a goal by id
// @tags goal
// @accept json
// @produce json
// @param id path string true "Goal ID"
// @success 200 {object} budgeting.GetGoalResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /goals/{id} [get]
func (h *budgettingHandlerImpl) GetGoalHandler(ctx *gin.Context) {
	val, exists := ctx.Get("claims")
	if !exists {
		h.logger.Error("Missing token claims")
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Missing token claims",
			Error:   "Missing token claims",
		})
		return
	}
	claims, err := token.TokenClaimsParse(val)
	if err != nil {
		ctx.JSON(403, models.ErrorResponse{
			Status:  403,
			Message: "Invalid access token",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	resp, err := h.goalsManagement.GetGoal(ctx, &pb.GetGoalReq{Id: id, UserId: claims.GetId()})
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

// @summary Update a goal
// @description Updates a goal
// @tags goal
// @accept json
// @produce json
// @param id path string true "Goal ID"
// @param goal body budgeting.UpdateGoalReq true "Updated goal"
// @success 200 {object} budgeting.UpdateGoalResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /goals/{id} [put]
func (h *budgettingHandlerImpl) UpdateGoalHandler(ctx *gin.Context) {
	var updateGoal pb.UpdateGoalReq
	if err := ctx.ShouldBindJSON(&updateGoal); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	updateGoal.Id = ctx.Param("id")
	resp, err := h.goalsManagement.UpdateGoal(ctx, &updateGoal)
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

// @summary Delete a goal
// @description Deletes a goal
// @tags goal
// @accept json
// @produce json
// @param id path string true "Goal ID"
// @success 200 {object} budgeting.DeleteGoalResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /goals/{id} [delete]
func (h *budgettingHandlerImpl) DeleteGoalHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.goalsManagement.DeleteGoal(ctx, &pb.DeleteGoalReq{Id: id})
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

// @summary Get all goals
// @description Gets all goals
// @tags goal
// @accept json
// @produce json
// @param Page query int false "Offset" default(1)
// @param Limit query int false "Limit" default(10)
// @param UserId query string false "Filter by user id"
// @param Name query string false "Filter by name"
// @param TargetAmount query float64 false "Filter by target amount"
// @param Status query string false "Filter by status" "status" (enum: [achieved, Failed, Inporgress])
// @param Deadline query string false "Filter by deadline"
// @success 200 {object} budgeting.GetGoalsResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /goals [get]
func (h *budgettingHandlerImpl) GetGoalsListHandler(ctx *gin.Context) {
	var query pb.GetGoalsReq
	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request query parameters",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.goalsManagement.GetGoals(ctx, &query)
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
