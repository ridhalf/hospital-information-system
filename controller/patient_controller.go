package controller

import "github.com/gin-gonic/gin"

type PatientController interface {
	RegisterPatient(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
