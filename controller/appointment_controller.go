package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
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
		HandleBindError(ctx)
		return
	}
	if !AllowReadPatient(ctx) || !PrivilegePatient(ctx, request.PatientID) {
		return
	}
	appointment, err := controller.serviceAppointment.CreateSchedule(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	HandleRequestSuccess(ctx, "create schedule is success", appointment)
	return
}
