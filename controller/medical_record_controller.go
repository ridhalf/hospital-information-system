package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
)

type MedicalRecordController interface {
	FindByPatientID(ctx *gin.Context)
}
type MedicalRecordControllerImpl struct {
	medicalRecordService service.MedicalRecordService
	auth                 auth.Jwt
}

func NewMedicalRecordController(medicalRecordService service.MedicalRecordService, auth auth.Jwt) MedicalRecordController {
	return &MedicalRecordControllerImpl{
		medicalRecordService: medicalRecordService,
		auth:                 auth,
	}
}

func (controller MedicalRecordControllerImpl) FindByPatientID(ctx *gin.Context) {
	request := web.MedicalRecordFindByPatientIDRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		response := api.APIResponse("find medical record is failed", http.StatusBadRequest, "BadRequest", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if !AllowReadPatient(ctx) || !PrivilegePatient(ctx, request.PatientID) {
		return
	}
	medicalRecords, err := controller.medicalRecordService.FindByPatientID(request)
	if err != nil {
		response := api.APIResponse("find medical record is failed", http.StatusBadRequest, "BadRequest", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	medicalRecordResponses := web.ToMedicalRecordResponses(medicalRecords)
	response := api.APIResponse("find medical record is success", http.StatusOK, "Success", medicalRecordResponses)
	ctx.JSON(http.StatusOK, response)
	return

}
