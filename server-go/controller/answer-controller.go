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
	MostAnswers(context *gin.Context)
	EditAnswer(context *gin.Context)
	DeleteAnswer(context *gin.Context)
	Update(*gin.Context)
}

type answerController struct {
	answerService     service.AnswerService
	jwtService        service.JWTService
	answerpostService service.AnswerPostService
}

func NewAnswerController(answer service.AnswerService, jwt service.JWTService, srvc service.AnswerPostService) AnswerController {
	return &answerController{
		answerService:     answer,
		jwtService:        jwt,
		answerpostService: srvc,
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

func (c *answerController) MostAnswers(context *gin.Context) {
	var answers []entity.MostAnswers = c.answerService.MostAnswers()
	res := helper.BuildResponse(true, "OK", answers)
	context.JSON(http.StatusOK, res)
}

func (c *answerController) EditAnswer(context *gin.Context) {
	var answer entity.Answer
	err := context.ShouldBind(&answer)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.answerService.EditAnswer(answer)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *answerController) DeleteAnswer(context *gin.Context) {
	_id := context.Param("id")
	stringToId, err := strconv.Atoi(_id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.answerService.DeleteAnswer(stringToId)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *answerController) Update(context *gin.Context) {
	var userAnswer entity.AnswerPost
	err := context.ShouldBind(&userAnswer)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	authHeader := context.GetHeader("Authorization")
	userID := service.GetUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseInt(userID, 10, 64)
	if err == nil {
		userAnswer.UserId = int(convertedUserID)
	}
	grade := userAnswer.Grade
	ifExist := c.answerpostService.Verify(userAnswer)

	if grade == 1 {
		if !ifExist {
			c.answerService.Update(entity.Answer{Id: userAnswer.AnswerId, Likes: 1})
			r := c.answerService.UpdateGrade("dislike", userAnswer.AnswerId)
			response := helper.BuildResponse(true, "OK", r)
			context.JSON(http.StatusOK, response)
		} else {
			res := helper.BuildResponse(true, "OK", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
		}
	} else {
		if !ifExist {
			c.answerService.Update(entity.Answer{Id: userAnswer.AnswerId, Dislikes: 1})
			r := c.answerService.UpdateGrade("like", userAnswer.AnswerId)
			response := helper.BuildResponse(true, "OK", r)
			context.JSON(http.StatusOK, response)
		} else {
			res := helper.BuildResponse(true, "OK", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
		}
	}
}
