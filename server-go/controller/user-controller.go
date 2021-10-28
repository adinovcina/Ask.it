package controller

import (
	"net/http"

	"github.com/adinovcina/entity"
	"github.com/adinovcina/helper"
	"github.com/adinovcina/models"
	"github.com/adinovcina/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(auth service.UserService, jwt service.JWTService) UserController {
	return &userController{
		userService: auth,
		jwtService:  jwt,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var loggedUser entity.User
	errDTO := ctx.ShouldBind(&loggedUser)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.userService.VerifyCredential(loggedUser.Email, loggedUser.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.Id)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *userController) Register(ctx *gin.Context) {
	var newUser entity.User
	errDTO := ctx.ShouldBind(&newUser)
	if errDTO != nil || len(newUser.FirstName) == 0 || len(newUser.LastName) == 0 {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsDuplicatedEmail(newUser.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)

	} else {
		createdUser := c.userService.CreateUser(newUser)
		service.SendEmail("New account", "Ask.it", createdUser.Email,
			createdUser.FirstName+" "+createdUser.LastName, newUser.Password)
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *userController) ChangePassword(ctx *gin.Context) {
	var newUser models.UserProfile
	errDTO := ctx.ShouldBind(&newUser)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsOldPasswordCorrect(newUser) {
		response := helper.BuildErrorResponse("Failed to process request", "Incorrect old password", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Password successfully changed", newUser.Email)
		ctx.JSON(http.StatusCreated, response)
	}
}
