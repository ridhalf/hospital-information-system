package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
)

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
		response := api.APIResponse("register patient is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	register, patient, err := controller.patientService.Register(request)
	if err != nil {
		response := api.APIResponse("register patient is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := controller.auth.GenerateToken(register.ID)
	if err != nil {
		response := api.APIResponse("register is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	userResponse := web.ToPatientRegisterResponse(register, patient, token)
	response := api.APIResponse("register patient is success", http.StatusOK, "Success", userResponse)
	ctx.JSON(http.StatusOK, response)

}

func (controller PatientControllerImpl) FindById(ctx *gin.Context) {
	request := web.PatientFindByIdRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		response := api.APIResponse("find patient is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	patient, err := controller.patientService.FindById(request)
	if err != nil {
		response := api.APIResponse("find patient is failed", http.StatusBadRequest, "BadRequest", request)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	patientResponse := web.ToPatientFindByIdResponse(patient)
	response := api.APIResponse("find patient is success", http.StatusOK, "Success", patientResponse)

	ctx.JSON(http.StatusOK, response)
}
