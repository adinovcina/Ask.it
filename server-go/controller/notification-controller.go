package controller

import (
	"net/http"
	"strconv"

	"github.com/adinovcina/helper"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type NotificationController interface {
	GetAll(context *gin.Context)
}

type notificationController struct {
	notificationService service.NotificationService
	jwtService          service.JWTService
}

func NewNotificationController(notification service.NotificationService,
	jwt service.JWTService) NotificationController {
	return &notificationController{
		notificationService: notification,
		jwtService:          jwt,
	}
}

func (c *notificationController) GetAll(context *gin.Context) {
	var q = context.Query("userId")
	stringToId, _ := strconv.Atoi(q)
	if stringToId != 0 {
		notifications := c.notificationService.GetAll(stringToId)
		res := helper.BuildResponse(true, "OK", notifications)
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Failed to process request", "Bad request", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
}
