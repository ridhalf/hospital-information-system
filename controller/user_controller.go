package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
)

type UserController interface {
	Register(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Login(ctx *gin.Context)
}
type UserControllerImpl struct {
	userService service.UserService
	auth        auth.Jwt
}

func NewUserController(userService service.UserService, auth auth.Jwt) UserController {
	return &UserControllerImpl{
		userService: userService,
		auth:        auth,
	}
}

func (controller UserControllerImpl) Register(ctx *gin.Context) {
	request := web.UserRegisterRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx, "register is failed")
		return
	}
	user, err := controller.userService.Register(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	token, err := controller.auth.GenerateToken(user.ID)
	if err != nil {
		response := api.APIResponse("generate token is failed", http.StatusBadRequest, "BadRequest", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := web.ToRegisterResponse(user, token)
	HandleRequestSuccess(ctx, "register is success", response)
	return
}
func (controller UserControllerImpl) FindById(ctx *gin.Context) {
	request := web.UserFindByIdRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		HandleBindError(ctx, "failed to bind request parameters")
		return
	}
	user, err := controller.userService.FindById(request)
	userDto := web.ToFindByIdResponse(user)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	HandleRequestSuccess(ctx, "user found successfully", userDto)
	return
}

func (controller UserControllerImpl) Login(ctx *gin.Context) {
	request := web.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx, "failed to bind request parameters")
		return
	}

	user, err := controller.userService.Login(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}

	token, err := controller.auth.GenerateToken(user.ID)
	if err != nil {
		response := api.APIResponse("generate token is failed", http.StatusBadRequest, "BadRequest", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userResponse := web.ToUserLoginResponse(user, token)
	HandleRequestSuccess(ctx, "login is success", userResponse)
	return
}
