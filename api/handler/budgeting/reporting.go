package bhandler

import (
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @summare 	Get sepending report
// @description This endpoint will return the sepending report
// @tags 		reporting
// @accepts 	json
// @produce 	json
// @param 		UserId query string false "User ID"
// @param 		Yearly query bool false "Yearly report"
// @param 		Monthly query bool false "Monthly report"
// @success 	200 {object} budgeting.GetSependingResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @failure 	500 {object} models.ErrorResponse
// @router 		/reporting/sepending [get]
func (h *budgettingHandlerImpl) GetSependingReportHandler(ctx *gin.Context) {
	var request pb.GetSependingReq
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.reportingManagement.GetSepending(ctx, &request)
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

// @summare 	Get income report
// @description This endpoint will return the income report
// @tags 		reporting
// @accepts 	json
// @produce 	json
// @param 		UserId query string false "User ID"
// @param 		Year query int false "Year report" default(2024)
// @param 		Month query int false "Monthly report" default(1)
// @success 	200 {object} budgeting.GetBudgetPerformanceResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @failure 	500 {object} models.ErrorResponse
// @router 		/reporting/budget-performance [get]
func (h *budgettingHandlerImpl) GetBudgetPerformanceHandler(ctx *gin.Context) {
	var request pb.GetBudgetPerformanceReq
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.reportingManagement.GetBudgetPerformance(ctx, &request)
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

// @summare 	Get goal progress report
// @description This endpoint will return the goal progress report
// @tags 		reporting
// @accepts 	json
// @produce 	json
// @param 		UserId query string false "User ID"
// @success 	200 {object} budgeting.GetGoalProgressResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @failure 	500 {object} models.ErrorResponse
// @router 		/reporting/goal-progress [get]
func (h *budgettingHandlerImpl) GetGoalProgressHandler(ctx *gin.Context) {
	var request pb.GetGoalProgressReq
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.reportingManagement.GoalProgress(ctx, &request)
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

// @summare 	get income report
// @description This endpoint will return the income report
// @tags 		reporting
// @accepts 	json
// @produce 	json
// @param 		UserId query string false "User ID"
// @param 		Yearly query bool false "Yearly report"
// @param 		Monthly query bool false "Monthly report"
// @success 	200 {object} budgeting.GetIncomeReportResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @failure 	500 {object} models.ErrorResponse
// @router 		/reporting/income [get]
func (h *budgettingHandlerImpl) GetIncomeReportHandler(ctx *gin.Context) {
	var request pb.GetIncomeReportReq
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.reportingManagement.GetIncome(ctx, &request)
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
