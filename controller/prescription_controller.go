package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
)

type PrescriptionController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}
type PrescriptionControllerImpl struct {
	prescriptionService service.PrescriptionService
}

func NewPrescriptionController(prescriptionService service.PrescriptionService) PrescriptionController {
	return &PrescriptionControllerImpl{
		prescriptionService: prescriptionService,
	}
}

func (controller PrescriptionControllerImpl) Create(ctx *gin.Context) {
	request := web.PrescriptionCreateRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx)
	}
	if !AllowReadDoctor(ctx) || !PrivilegeDoctor(ctx, request.DoctorID) {
		return
	}
	prescription, err := controller.prescriptionService.Create(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	response := web.ToPrescriptionCreateResponse(prescription)
	HandleRequestSuccess(ctx, "your data has been successfully saved.", response)
	return
}
func (controller PrescriptionControllerImpl) Update(ctx *gin.Context) {
	request := web.PrescriptionUpdateRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx)
	}
	if !AllowReadDoctor(ctx) || !PrivilegeDoctor(ctx, request.DoctorID) {
		return
	}
	prescription, err := controller.prescriptionService.Update(request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	response := web.ToPrescriptionCreateResponse(prescription)
	HandleRequestSuccess(ctx, "your data has been successfully saved.", response)
	return
}
