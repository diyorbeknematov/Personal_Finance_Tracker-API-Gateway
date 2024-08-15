package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create account
// @description Create account
// @Tags account
// @Accept json
// @Produce json
// @Param account body budgeting.CreateAccountReq true "account"
// @Success 200 {object} budgeting.CreateAccountResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /accounts [post]
func (h *budgettingHandlerImpl) CreateAccountHandler(ctx *gin.Context) {
	var account pb.CreateAccountReq

	if err := ctx.ShouldBindJSON(&account); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.financeManagement.CreateAccount(ctx, &account)
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(200, resp)
}

// @Summary get account
// @Description get account
// @tags account
// @Accept json
// @produce json
// @param id path string true "account id"
// @Success 200 {object} budgeting.GetAccountResp
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /accounts/{id} [get]
func (h *budgettingHandlerImpl) GetAccountHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	account, err := h.financeManagement.GetAccount(ctx, &pb.GetAccountReq{
		Id:     id,
	})
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, account)
}

// @Summary update account
// @Description update account
// @Tags account
// @Accept json
// @Produce json
// @param id path string true "account id"
// @param account body budgeting.UpdateAccountReq true "account"
// @Success 200 {object} budgeting.UpdateAccountResp
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /accounts/{id} [put]
func (h *budgettingHandlerImpl) UpdateAccountHandler(ctx *gin.Context) {
	var account pb.UpdateAccountReq
	if err := ctx.ShouldBindJSON(&account); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	account.Id = id
	resp, err := h.financeManagement.UpdateAccount(ctx, &account)
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @summary delete account
// @description delete account
// @tags account
// @Accept json
// @Produce json
// @param id path string true "account id"
// @Success 200 {object} budgeting.DeleteAccountResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @Router /accounts/{id} [delete]
func (h *budgettingHandlerImpl) DeleteAccountHandler(ctx *gin.Context) {
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
		h.logger.Error("Invalid access token", "error", err)
		ctx.AbortWithStatusJSON(403, models.ErrorResponse{
			Status:  403,
			Message: "Invalid access token",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	resp, err := h.financeManagement.DeleteAccount(ctx, &pb.DeleteAccountReq{
		Id:     id,
		UserId: claims.GetId(),
	})
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @summary get accounts list
// @description get accounts list
// @tags account
// @Accept json
// @produce json
// @param Offset query int false "offset" default(1)
// @param Limit query int false "limit" default(10)
// @param UserId query string false "user id"
// @param Name query string false "name"
// @param Currency query string false "currency"
// @param Type query string false "type"
// @param Balance query float64 false "balance"
// @param CreatedAt query bool false "created at"
// @Success 200 {object} budgeting.GetAccountsListResp
// @Failure 404 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @router /accounts [get]
func (h *budgettingHandlerImpl) GetAccountsListHandler(ctx *gin.Context) {
	var pbReq pb.GetAccountsListReq
	if err := ctx.ShouldBindQuery(&pbReq); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.financeManagement.GetAccountsList(ctx, &pbReq)
	if err != nil {
		h.logger.Error("Internal Server Error", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(200, resp)
}
