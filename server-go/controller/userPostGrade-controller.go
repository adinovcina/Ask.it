package controller

import (
	"net/http"
	"strconv"

	"github.com/adinovcina/entity"
	"github.com/adinovcina/helper"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type UserPostController interface {
	GetAll(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
}

type userpostController struct {
	userPostService service.UserPostService
	jwtService      service.JWTService
	postService     service.PostService
}

func NewUserPostController(userPost service.UserPostService, jwt service.JWTService, postser service.PostService) UserPostController {
	return &userpostController{
		userPostService: userPost,
		jwtService:      jwt,
		postService:     postser,
	}
}

func (c *userpostController) GetAll(context *gin.Context) {
	var posts []entity.UserPost = c.userPostService.GetAll()
	res := helper.BuildResponse(true, "OK", posts)
	context.JSON(http.StatusOK, res)
}

func (c *userpostController) Update(context *gin.Context) {
	var userPost entity.UserPost
	err := context.ShouldBind(&userPost)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			userPost.UserId = int(convertedUserID)
		}
		res := c.userPostService.VerifyIfGradeExist(userPost)

		if res.Id != 0 {
			result := c.userPostService.UpdateAnswerMark(userPost)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		} else {
			res := helper.BuildErrorResponse("Failed to process request", "Bad request", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
			return
		}
	}
}

func (c *userpostController) Insert(context *gin.Context) {
	var userPost entity.UserPost
	err := context.ShouldBind(&userPost)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			userPost.UserId = int(convertedUserID)
		}

		response := c.userPostService.VerifyIfDataExist(userPost)

		if response {
			res := helper.BuildErrorResponse("Failed to process request", "Bad request", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
			return
		} else {
			result := c.userPostService.Insert(userPost)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}
}
