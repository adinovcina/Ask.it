package controller

import (
	"net/http"
	"strconv"

	"github.com/adinovcina/entity"
	"github.com/adinovcina/helper"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	GetAll(*gin.Context)
	Insert(*gin.Context)
}

type postController struct {
	postService service.PostService
	jwtService  service.JWTService
}

func NewPostController(post service.PostService, jwt service.JWTService) PostController {
	return &postController{
		postService: post,
		jwtService:  jwt,
	}
}

func (c *postController) GetAll(context *gin.Context) {
	var posts []entity.Post = c.postService.GetAll()
	res := helper.BuildResponse(true, "OK", posts)
	context.JSON(http.StatusOK, res)
}

func (c *postController) Insert(context *gin.Context) {
	var post entity.Post
	err := context.ShouldBind(&post)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := service.GetUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseInt(userID, 10, 64)
		if err == nil {
			post.UserId = int(convertedUserID)
		}
		result := c.postService.Insert(post)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
