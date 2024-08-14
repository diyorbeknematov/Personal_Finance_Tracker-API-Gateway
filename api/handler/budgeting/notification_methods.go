package bhandler

import (
	"api-gateway/api/token"
	pb "api-gateway/generated/budgeting"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @summary 	Send notification to a user
// @description Send notification to a user
// @tags 		notification
// @accept 		json
// @produce 	json
// @param 		sendNotificationReq body budgeting.SendNotificationReq true "Send notification request"
// @success 	200 {object} budgeting.SendNotificationResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	500 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @route 		/notification/send [post]
func (h *budgettingHandlerImpl) SendNotificationHandler(ctx *gin.Context) {
	var sendNotificationReq pb.SendNotificationReq
	if err := ctx.ShouldBindJSON(&sendNotificationReq); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}
	resp, err := h.reportingManagement.SendNotification(ctx, &sendNotificationReq)
	if err != nil {
		h.logger.Error("Error sending notification", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status:  500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		})
		return
	}
	// data, err := json.Marshal(&sendNotificationReq)
	// if err != nil {
	// 	h.logger.Error("Error marshaling response", "error", err)
	//     ctx.JSON(500, models.ErrorResponse{
	//         Status:  500,
	//         Message: "Internal Server Error",
	//         Error:   err.Error(),
	//     })
	//     return
	// }
	// err = h.producer.ProducerMessage("notifications", data)
	// if err!= nil {
	// 	h.logger.Error("Failed to produce message", "error", err)
	//     ctx.JSON(500, models.ErrorResponse{
	//         Status:  500,
	//         Message: "Internal Server Error",
	//         Error:   err.Error(),
	//     })
	//     return
	// }
	ctx.JSON(200, resp)
}

// @summary 	Get notification list
// @description Get notification list
// @tags 		notification
// @accept 		json
// @produce 	json
// @success 	200 {object} budgeting.GetNotificationsListResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	500 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @route 		/notification/list [get]
func (h *budgettingHandlerImpl) GetNotificationListHandler(ctx *gin.Context) {
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
		h.logger.Error("Invalid access token")
        ctx.JSON(403, models.ErrorResponse{
            Status:  403,
            Message: "Invalid access token",
            Error:   err.Error(),
        })
        return
	}
	resp, err := h.reportingManagement.GetNotificationList(ctx, &pb.GetNotificationsListReq{
		UserId: claims.GetId(),
	})
	if err!= nil {
		h.logger.Error("Error getting notification list", "error", err)
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
	}
	ctx.JSON(200, resp)
}

// @summary 	Get notification
// @description Get notification
// @tags 		notification
// @accept 		json
// @produce 	json
// @param 		id path string true "Notification ID"
// @success 	200 {object} budgeting.GetNotificationResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	500 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @router		/notification/{id} [get]
func (h *budgettingHandlerImpl) GetNotificationHandler(ctx *gin.Context) {
	id := ctx.Param("id")
    resp, err := h.reportingManagement.GetNotification(ctx, &pb.GetNotificationReq{Id: id})
    if err!= nil {
        h.logger.Error("Error getting notification", "error", err)
        ctx.JSON(500, models.ErrorResponse{
            Status:  500,
            Message: "Internal Server Error",
            Error:   err.Error(),
        })
        return
    }
    ctx.JSON(200, resp)
}

// @summary 	Update notification
// @description Update notification
// @tags 		notification
// @accept 		json
// @produce 	json
// @param 		id path string true "Notification ID"
// @param 		notification body budgeting.UpdateNotificationReq true "Updated notification"
// @success 	200 {object} budgeting.UpdateNotificationResp
// @failure 	400 {object} models.ErrorResponse
// @failure 	500 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @router		/notification/{id} [put]
func (h *budgettingHandlerImpl) UpdateNotificationHandler(ctx *gin.Context) {
	var updateNotificationReq pb.UpdateNotificationReq
	if err := ctx.ShouldBindJSON(&updateNotificationReq); err != nil {
		h.logger.Error("Invalid request body", "error", err)
		ctx.JSON(400, models.ErrorResponse{
			Status: 400,
			Message: "Invalid request body",
			Error: err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	updateNotificationReq.Id = id
	resp, err := h.reportingManagement.UpdateNotification(ctx, &updateNotificationReq)
	if err != nil {
		h.logger.Error("Error updating notification", "error", err)
		ctx.JSON(500, models.ErrorResponse{
			Status: 500,
			Message: "Internal server error",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}

// @summary 	Delete notification
// @description Delete notification
// @tags		notification
// @accept 		json
// @produce 	json
// @param 		id path string true "Notification ID"
// @success 	200 {object} budgeting.DeleteNotificationResp
// @failure 	400 {object} models.ErrorResponse
// @failure		500 {object} models.ErrorResponse
// @failure 	404 {object} map[string]interface{}
// @router		/notification/{id} [delete]
func (h *budgettingHandlerImpl) DeleteNotificationHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.reportingManagement.DeleteNotification(ctx, &pb.DeleteNotificationReq{Id: id})
	if err!= nil {
		h.logger.Error("Error deleting notification", "error", err)
        ctx.JSON(500, models.ErrorResponse{
            Status: 500,
            Message: "Internal server error",
            Error: err.Error(),
        })
        return
	}
	ctx.JSON(200, resp)
}