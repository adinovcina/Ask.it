package controller

import (
	"net/http"
	"strconv"

	"github.com/adinovcina/entity"
	"github.com/adinovcina/helper"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type AnswerController interface {
	GetAll(context *gin.Context)
	Insert(context *gin.Context)
	// MostAnswers(context *gin.Context)
	// UpdateAnswerMark(context *gin.Context)
}

type answerController struct {
	answerService service.AnswerService
	jwtService    service.JWTService
}

func NewAnswerController(answer service.AnswerService, jwt service.JWTService) AnswerController {
	return &answerController{
		answerService: answer,
		jwtService:    jwt,
	}
}

func (c *answerController) GetAll(context *gin.Context) {
	var answers []entity.Answer = c.answerService.GetAll()
	res := helper.BuildResponse(true, "OK", answers)
	context.JSON(http.StatusOK, res)
}

func (c *answerController) Insert(context *gin.Context) {
	var answer entity.Answer
	err := context.ShouldBind(&answer)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			answer.UserId = int(convertedUserID)
		}
		result := c.answerService.Insert(answer)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// func (c *answerController) MostAnswers(context *gin.Context) {
// 	var answers []models.MostAnswers = c.answerService.MostAnswers()
// 	res := helper.BuildResponse(true, "OK", answers)
// 	context.JSON(http.StatusOK, res)
// }
