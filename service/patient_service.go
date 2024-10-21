package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
	"hospital-information-system/repository"
	"strings"
	"time"
)

type PatientService interface {
	Register(request web.PatientRegisterRequest) (domain.User, domain.Patient, error)
	FindById(request web.PatientFindByIdRequest) (domain.Patient, error)
	FindByUserId(userId int) (domain.Patient, error)
}
type PatientServiceImpl struct {
	patientRepository repository.PatientRepository
	userRepository    repository.UserRepository
}

func NewPatientService(patientRepository repository.PatientRepository, userRepository repository.UserRepository) PatientService {
	return &PatientServiceImpl{
		patientRepository: patientRepository,
		userRepository:    userRepository,
	}
}
func (service PatientServiceImpl) Register(request web.PatientRegisterRequest) (domain.User, domain.Patient, error) {
	if request.User.Password != request.User.PasswordConfirmation {
		return domain.User{}, domain.Patient{}, errors.New("password incorrect")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, domain.Patient{}, errors.New("failed to generate password hash")
	}
	user := domain.User{
		Username: strings.Split(request.User.Email, "@")[0],
		Password: string(password),
		Email:    request.User.Email,
		Role:     request.User.Role,
	}
	register, err := service.userRepository.Save(user)
	if err != nil {
		return domain.User{}, domain.Patient{}, errors.New("failed to save user")
	}
	if request.DateOfBirth == "" {
		return register, domain.Patient{}, errors.New("invalid date of birth")
	}
	dateOfBirth, err := time.Parse("2006-04-02", request.DateOfBirth)
	if err != nil {
		return domain.User{}, domain.Patient{}, errors.New("invalid date of birth")
	}

	patient := domain.Patient{
		UserID:      register.ID,
		Name:        request.Name,
		DateOfBirth: dateOfBirth,
		Address:     request.Address,
		Phone:       request.Phone,
	}
	save, err := service.patientRepository.Save(patient)
	if err != nil {
		return domain.User{}, domain.Patient{}, errors.New("failed to save patient")
	}
	return register, save, nil
}

func (service PatientServiceImpl) FindById(request web.PatientFindByIdRequest) (domain.Patient, error) {
	patient, err := service.patientRepository.FindById(request.Id, true)
	if err != nil {
		return domain.Patient{}, errors.New("failed to find patient")
	}
	return patient, nil

}
func (service PatientServiceImpl) FindByUserId(userId int) (domain.Patient, error) {
	patient, err := service.patientRepository.FindByUserId(userId, true)
	if err != nil {
		return domain.Patient{}, errors.New("failed to find patient")
	}
	return patient, nil
}
