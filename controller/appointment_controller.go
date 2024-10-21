package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
)

type AppointmentController interface {
	CreateSchedule(ctx *gin.Context)
}
type AppointmentControllerImpl struct {
	serviceAppointment service.AppointmentService
	auth               auth.Jwt
}

func NewAppointmentController(serviceAppointment service.AppointmentService, auth auth.Jwt) AppointmentController {
	return &AppointmentControllerImpl{
		auth:               auth,
		serviceAppointment: serviceAppointment,
	}
}

func (controller AppointmentControllerImpl) CreateSchedule(ctx *gin.Context) {
	request := web.AppointmentCreateScheduleRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := api.APIResponse("create schedule is failed", http.StatusBadRequest, "BadRequest", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if !AllowReadPatient(ctx) || !PrivilegePatient(ctx, request.PatientID) {
		return
	}
	appointment, err := controller.serviceAppointment.CreateSchedule(request)
	if err != nil {
		response := api.APIResponse(err.Error(), http.StatusConflict, "Conflict", nil)
		ctx.JSON(http.StatusConflict, response)
		return
	}
	response := api.APIResponse("create schedule is success", http.StatusOK, "OK", appointment)
	ctx.JSON(http.StatusOK, response)
	return
}
