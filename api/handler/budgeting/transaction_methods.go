package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary Create transaction
// @description Create transaction
// @tags transaction
// @Accept json
// @produce json
// @param transaction body budgeting.CreateTransactionReq true "transaction"
// @Success 200 {object} models.Response
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /transactions [post]
func (h *budgettingHandlerImpl) CreateTransactionHandler(ctx *gin.Context) {
	var transaction pb.CreateTransactionReq
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	// resp, err := h.financeManagement.CreateTransaction(ctx, &transaction)
	// if err != nil {
	// 	h.logger.Error("Internal Server Error", "error", err)
	// 	ctx.JSON(500, models.ErrorResponse{
	// 		Status:  500,
	// 		Message: "Internal Server Error",
	// 		Error:   err.Error(),
	// 	})
	// 	return
	// }
	data, err := json.Marshal(&transaction)
	if err != nil {
		h.logger.Error("Failed to marshal transaction", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	err = h.producer.ProducerMessage("transactions", data)
	if err != nil {
		h.logger.Error("Failed to produce message", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Message: "Transaction created successfully",
		Status:  "success",
	})
}

// @summary get transaction
// @description get transaction
// @tags transaction
// @Accept json
// @produce json
// @param id path string true "transaction id"
// @success 200 {object} budgeting.GetTransactionResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /transactions/{id} [get]
func (h *budgettingHandlerImpl) GetTransactionHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := h.financeManagement.GetTransaction(ctx, &pb.GetTransactionReq{
		Id:     id,
	})
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

// @summary update transaction
// @description update transaction
// @tags transaction
// @accept json
// @produce json
// @param id path string true "transaction id"
// @param transaction body budgeting.UpdateTransactionReq true "transaction"
// @success 200 {object} budgeting.UpdateTransactionResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /transactions/{id} [put]
func (h *budgettingHandlerImpl) UpdateTransactionHandler(ctx *gin.Context) {
	var transaction pb.UpdateTransactionReq
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	transaction.Id = ctx.Param("id")
	resp, err := h.financeManagement.UpdateTransaction(ctx, &transaction)
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary delete transaction
// @Description delete transaction
// @Tags transaction
// @accept json
// @produce json
// @Param id path string true "transaction id"
// @success 200 {object} budgeting.DeleteTransactionResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /transactions/{id} [delete]
func (h *budgettingHandlerImpl) DeleteTransactionHandler(ctx *gin.Context) {
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
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid access token",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	resp, err := h.financeManagement.DeleteTransaction(ctx, &pb.DeleteTransactionReq{
		Id:     id,
		UserId: claims.GetId(),
	})

	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary get transactions list
// @description get transactions list
// @Tags transaction
// @accept json
// @produce json
// @param Limit query int false "limit" default(10)
// @param Page query int false "offset" default(1)
// @param AccountName query string false "Account name"
// @param Amount query float64 false "Amount"
// @param CategoryName query string false "Category name"
// @param Type query string false "Type" "type" (enum: [expense, income])
// @param Description query string false "Description"
// @param DateFrom query string false "Date from"
// @param DateTo query string false "Date to"
// @success 200 {object} budgeting.GetTransactionsListResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /transactions [get]
func (h *budgettingHandlerImpl) GetTransactionsListHandler(ctx *gin.Context) {
	var query pb.GetTransactionsListReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.financeManagement.GetTransactionsList(ctx, &query)
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
