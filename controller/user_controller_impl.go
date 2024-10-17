package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
)

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
		response := api.APIResponse("register is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := controller.userService.Register(request)
	if err != nil {
		response := api.APIResponse("register is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := controller.auth.GenerateToken(user.Id)
	if err != nil {
		response := api.APIResponse("register is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := web.ToRegisterResponse(user, token)
	apiResponse := api.APIResponse("register is success", http.StatusOK, "Success", response)
	ctx.JSON(http.StatusOK, apiResponse)
}
func (controller UserControllerImpl) FindById(ctx *gin.Context) {
	request := web.UserFindByIdRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		response := api.APIResponse("find is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := controller.userService.FindById(request)
	userDto := web.ToFindByIdResponse(user)
	if err != nil {
		response := api.APIResponse("find is failed", http.StatusBadRequest, "BadRequest", userDto)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := api.APIResponse("find is success", http.StatusOK, "Success", userDto)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller UserControllerImpl) Login(ctx *gin.Context) {
	request := web.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := api.APIResponse("login is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := controller.userService.Login(request)
	if err != nil {
		response := api.APIResponse("login is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := controller.auth.GenerateToken(user.Id)
	if err != nil {
		response := api.APIResponse("login is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userResponse := web.ToUserLoginResponse(user, token)
	response := api.APIResponse("login is success", http.StatusOK, "Success", userResponse)
	ctx.JSON(http.StatusOK, response)
}
