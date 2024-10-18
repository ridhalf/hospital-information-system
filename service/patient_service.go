package service

import (
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
)

type PatientService interface {
	Register(request web.PatientRegisterRequest) (domain.User, domain.Patient, error)
	FindById(request web.PatientFindByIdRequest) (domain.Patient, error)
}
