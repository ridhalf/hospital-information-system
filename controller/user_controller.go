package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Login(ctx *gin.Context)
}
