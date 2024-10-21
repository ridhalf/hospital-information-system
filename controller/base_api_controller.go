package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/model/api"
	"hospital-information-system/model/constants"
	"hospital-information-system/model/domain"
	"net/http"
)

func AllowReadPatient(ctx *gin.Context) bool {
	user := ctx.MustGet("user").(domain.User)
	if user.Role == constants.PATIENT || user.Role == constants.ADMIN || user.Role == constants.NURSE {
		return true
	}
	ValidateMessage(ctx)
	return false
}
func PrivilegePatient(ctx *gin.Context, patientId int) bool {
	patient := ctx.MustGet("patient").(domain.Patient)
	if patient.ID == patientId {
		return true
	}
	ValidateMessage(ctx)
	return false
}
func ValidateMessage(ctx *gin.Context) {
	response := api.APIResponse("sorry, but you do not have access to this area. Contact support for further assistance", http.StatusForbidden, "Forbidden", nil)
	ctx.JSON(http.StatusForbidden, response)
	return
}
