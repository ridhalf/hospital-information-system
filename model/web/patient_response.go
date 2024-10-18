package web

import (
	"hospital-information-system/model/domain"
	"time"
)

type PatientRegisterResponse struct {
	Id          int                  `json:"id"`
	Name        string               `json:"name"`
	DateOfBirth time.Time            `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	Address     string               `json:"address"`
	Phone       string               `json:"phone"`
	User        UserRegisterResponse `json:"user"`
}

func ToPatientRegisterResponse(user domain.User, patient domain.Patient, token string) PatientRegisterResponse {
	return PatientRegisterResponse{
		Id:          patient.ID,
		Name:        patient.Name,
		DateOfBirth: patient.DateOfBirth,
		Address:     patient.Address,
		Phone:       patient.Phone,
		User: UserRegisterResponse{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Token:    token,
		},
	}
}

type PatientFindByIdResponse struct {
	Id          int                  `json:"id"`
	Name        string               `json:"name"`
	DateOfBirth time.Time            `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	Address     string               `json:"address"`
	Phone       string               `json:"phone"`
	User        UserFindByIdResponse `json:"user"`
}

func ToPatientFindByIdResponse(patient domain.Patient) PatientFindByIdResponse {
	return PatientFindByIdResponse{
		Id:          patient.ID,
		Name:        patient.Name,
		DateOfBirth: patient.DateOfBirth,
		Address:     patient.Address,
		Phone:       patient.Phone,
		User: UserFindByIdResponse{
			Id:       patient.User.ID,
			Username: patient.User.Username,
			Email:    patient.User.Email,
			Role:     patient.User.Role,
		},
	}
}
