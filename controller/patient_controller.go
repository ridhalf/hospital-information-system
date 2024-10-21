package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
)

type PatientController interface {
	RegisterPatient(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
type PatientControllerImpl struct {
	patientService service.PatientService
	auth           auth.Jwt
}

func NewPatientController(patientService service.PatientService, auth auth.Jwt) PatientController {
	return &PatientControllerImpl{
		patientService: patientService,
		auth:           auth,
	}
}

func (controller PatientControllerImpl) RegisterPatient(ctx *gin.Context) {
	request := web.PatientRegisterRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx, "register patient is failed")
		return
	}

	register, patient, err := controller.patientService.Register(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	token, err := controller.auth.GenerateToken(register.ID)
	if err != nil {
		HandleGenerateTokenError(ctx)
		return
	}
	userResponse := web.ToPatientRegisterResponse(register, patient, token)
	HandleRequestSuccess(ctx, "register patient is success", userResponse)
	return

}

func (controller PatientControllerImpl) FindById(ctx *gin.Context) {
	request := web.PatientFindByIdRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		HandleBindError(ctx, "find patient is failed")
		return
	}
	if !AllowReadPatient(ctx) || !PrivilegePatient(ctx, request.Id) {
		return
	}
	patient, err := controller.patientService.FindById(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	patientResponse := web.ToPatientFindByIdResponse(patient)
	HandleRequestSuccess(ctx, "find patient is success", patientResponse)
	return
}
