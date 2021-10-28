package controller

import (
	"net/http"
	"strconv"

	"github.com/adinovcina/entity"
	"github.com/adinovcina/helper"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type AnswerPostController interface {
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Insert(context *gin.Context)
}

type answerpostController struct {
	answerPostService service.AnswerPostService
	jwtService        service.JWTService
	answerService     service.AnswerService
}

func AnswerUserPostController(answerPost service.AnswerPostService, jwt service.JWTService, postser service.AnswerService) AnswerPostController {
	return &answerpostController{
		answerPostService: answerPost,
		jwtService:        jwt,
		answerService:     postser,
	}
}

func (c *answerpostController) GetAll(context *gin.Context) {
	var posts []entity.AnswerPost = c.answerPostService.GetAll()
	res := helper.BuildResponse(true, "OK", posts)
	context.JSON(http.StatusOK, res)
}

func (c *answerpostController) Update(context *gin.Context) {
	var answerPost entity.AnswerPost
	err := context.ShouldBind(&answerPost)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			answerPost.UserId = int(convertedUserID)
		}
		res := c.answerPostService.VerifyIfGradeExist(answerPost)

		if res.Id != 0 {
			result := c.answerPostService.UpdateAnswerMark(answerPost)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		} else {
			res := helper.BuildErrorResponse("Failed to process request", "Bad request", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
			return
		}
	}
}

func (c *answerpostController) Insert(context *gin.Context) {
	var answerPost entity.AnswerPost
	err := context.ShouldBind(&answerPost)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			answerPost.UserId = int(convertedUserID)
		}

		response := c.answerPostService.VerifyIfDataExist(answerPost)

		if response {
			res := helper.BuildErrorResponse("Failed to process request", "Bad request", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
			return
		} else {
			result := c.answerPostService.Insert(answerPost)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}
