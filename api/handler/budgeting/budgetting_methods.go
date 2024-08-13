package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// @Summary Create budget
// @Description Create budget
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body budgeting.CreateBudgetReq true "Budget"
// @Success 200 {object} budgeting.CreateBudgetResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /budgets [post]
func (h *budgettingHandlerImpl) CreateBudgetHandler(ctx *gin.Context) {
	var budget pb.CreateBudgetReq
	if err := ctx.ShouldBindJSON(&budget); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.budgetManagement.CreateBudget(ctx, &budget)
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

// @Summary get budget
// @Description get budget
// @Tags budget
// @Accept json
// @Produce json
// @Param id path string true "budget id"
// @Success 200 {object} budgeting.GetBudgetResp
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /budgets/{id} [get]
func (h *budgettingHandlerImpl) GetBudgetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	budget, err := h.budgetManagement.GetBudget(ctx, &pb.GetBudgetReq{Id: id})
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, budget)
}

// @summary update budget
// @Description update budget
// @tags budget
// @Accept json
// @Produce json
// @param id path string true "budget id"
// @param budget body budgeting.UpdateBudgetReq true "budget"
// @Success 200 {object} models.Response
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /budgets/{id} [put]
func (h *budgettingHandlerImpl) UpdateBudgetHandler(ctx *gin.Context) {
	var budget pb.UpdateBudgetReq
	if err := ctx.ShouldBindJSON(&budget); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	budget.Id = ctx.Param("id")
	// resp, err := h.budgetManagement.UpdateBudget(ctx, &budget)
	// if err != nil {
	// 	ctx.JSON(500, models.ErrorResponse{
	// 		Status:  500,
	// 		Message: "Internal Server Error",
	// 		Error:   err.Error(),
	// 	})
	// 	return
	// }
	data, err := json.Marshal(&budget)
	if err != nil {
		h.logger.Error("Failed to marshal budget", "error", err)
		ctx.JSON(400, models.ErrorResponse{
            Status:  400,
            Message: "Invalid request body",
            Error:   err.Error(),
        })
	}	
	err = h.producer.ProducerMessage("budgets", data)
	if err != nil {
		h.logger.Error("Failed to produce message", "error", err)
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
	}
	ctx.JSON(200, models.Response{
		Status: "success",
        Message: "Budget updated successfully",
	})
}

// @summary delete budget
// @description delete budget
// @tags budget
// @Accept json
// @produce json
// @param id path string true "budget id"
// @Success 200 {object} budgeting.DeleteBudgetResp
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /budgets/{id} [delete]
func (h *budgettingHandlerImpl) DeleteBudgetHandler(ctx *gin.Context) {
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
		ctx.AbortWithStatusJSON(403, models.ErrorResponse{
			Status:  403,
			Message: "Invalid access token",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	resp, err := h.budgetManagement.DeleteBudget(ctx, &pb.DeleteBudgetReq{
		Id:     id,
		UserId: claims.GetId(),
	})
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

// @summary list budgets
// @Description list budgets
// @Tags budget
// @Accept json
// @Produce json
// @Param UserId query string false "user id"
// @Param CategoryId query string false "category id"
// @Param Period query string false "period" (enum: day, week, month, year)
// @Param StartDate query string false "start date"
// @Param EndDate query string false "end date"
// @Success 200 {object} budgeting.GetBudgetsResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /budgets [get]
func (h *budgettingHandlerImpl) GetBudgetsListHandler(ctx *gin.Context) {
	var budgets pb.GetBudgetsReq
	if err := ctx.BindQuery(&budgets); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request query parameters",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.budgetManagement.GetBudgetsList(ctx, &budgets)
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
