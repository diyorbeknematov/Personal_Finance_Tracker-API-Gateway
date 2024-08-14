package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary Create a new category
// @description Creates a new category
// @tags category
// @accept json
// @produce json
// @param category body budgeting.CreateCategoryReq true "New category"
// @success 200 {object} budgeting.CreateCategoryResp
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @failure 404 {object} map[string]interface{}
// @router /categories [post]
func (h *budgettingHandlerImpl) CreateCategoryHandler(ctx *gin.Context) {
	var category pb.CreateCategoryReq
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	resp, err := h.budgetManagement.CreateCategory(ctx, &category)
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

// @summary Get a category by id
// @description Gets a category by id
// @tags category
// @accept json
// @produce json
// @param id path string true "Category ID"
// @success 200 {object} budgeting.GetCategoryResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /categories/{id} [get]
func (h *budgettingHandlerImpl) GetCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.budgetManagement.GetCategory(ctx, &pb.GetCategoryReq{Id: id})
    if err!= nil {
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
    }

    ctx.JSON(200, resp)
}


func (h *budgettingHandlerImpl) UpdateCategoryHandler(ctx *gin.Context) {
	var category pb.UpdateCategoryReq
    if err := ctx.ShouldBindJSON(&category); err!= nil {
        ctx.JSON(400, models.ErrorResponse{
            Status:  400,
            Message: "Invalid request body",
            Error:   err.Error(),
        })
        return
    }

	category.Id = ctx.Param("id")
    resp, err := h.budgetManagement.UpdateCategory(ctx, &category)
    if err!= nil {
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
    }

    ctx.JSON(200, resp)
}

// @summary Delete a category
// @description Deletes a category
// @tags category
// @accept json
// @produce json
// @param id path string true "Category ID"
// @success 200 {object} budgeting.DeleteCategoryResp
// @failure 404 {object} map[string]interface{}
// @failure 400 {object} models.ErrorResponse
// @failure 500 {object} models.ErrorResponse
// @router /categories/{id} [delete]
func (h *budgettingHandlerImpl) DeleteCategoryHandler(ctx *gin.Context) {
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
	if err!= nil {
		ctx.JSON(403, models.ErrorResponse{
            Status:  403,
            Message: "Invalid access token",
            Error:   err.Error(),
        })
        return
	}
	id := ctx.Param("id")
    resp, err := h.budgetManagement.DeleteCategory(ctx, &pb.DeleteCategoryReq{Id: id, UserId: claims.GetId()})
    if err!= nil {
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
    }

    ctx.JSON(200, resp)
}

// @summary Get all categories
// @description Gets all categories
// @tags category
// @accept json
// @produce json
// @Param Limit query int false "Limit" default(10)
// @Param Page query int false "Offset" default(1)
// @Param UserId query string false "User ID"
// @Param Name query string false "Name"
// @Param Type query string false "Type" "type" (enum: [expence, income])
// @Success 200 {object} budgeting.GetCategoriesResp
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} map[string]interface{}
// @router /categories [get]
func (h *budgettingHandlerImpl) GetCategoriesListHandler(ctx *gin.Context) {
	var query pb.GetCategoriesReq
	if err := ctx.ShouldBindQuery(&query); err!= nil {
		ctx.JSON(400, models.ErrorResponse{
            Status:  400,
            Message: "Invalid request body",
            Error:   err.Error(),
        })
        return
	}
	resp, err := h.budgetManagement.GetCategoriesList(ctx, &query)
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
	}

	ctx.JSON(http.StatusOK, resp)
}
