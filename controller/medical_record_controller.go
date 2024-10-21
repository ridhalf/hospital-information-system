package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
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
		HandleBindError(ctx)
		return
	}
	if !AllowReadPatient(ctx) || !PrivilegePatient(ctx, request.PatientID) {
		return
	}
	medicalRecords, err := controller.medicalRecordService.FindByPatientID(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	medicalRecordResponses := web.ToMedicalRecordResponses(medicalRecords)
	HandleRequestSuccess(ctx, "medical record found is successfully", medicalRecordResponses)
	return

}
