package service

import (
	"errors"
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
	"hospital-information-system/repository"
)

type PrescriptionService interface {
	Create(request web.PrescriptionCreateRequest) (domain.Prescription, error)
}
type PrescriptionServiceImpl struct {
	prescriptionRepository repository.PrescriptionRepository
	patientRepository      repository.PatientRepository
	userRepository         repository.UserRepository
}

func NewPrescriptionService(prescriptionRepository repository.PrescriptionRepository, patientRepository repository.PatientRepository, userRepository repository.UserRepository) PrescriptionService {
	return &PrescriptionServiceImpl{
		prescriptionRepository: prescriptionRepository,
		patientRepository:      patientRepository,
		userRepository:         userRepository,
	}
}

func (service PrescriptionServiceImpl) Create(request web.PrescriptionCreateRequest) (domain.Prescription, error) {
	count, err := service.userRepository.Count(request.DoctorID)
	if err != nil {
		return domain.Prescription{}, errors.New("an error occurred while processing your request")
	}
	if count == 0 {
		return domain.Prescription{}, errors.New("doctor is not found")
	}
	count, err = service.patientRepository.Count(request.PatientID)
	if err != nil {
		return domain.Prescription{}, errors.New("an error occurred while processing your request")
	}
	if count == 0 {
		return domain.Prescription{}, errors.New("patient is not found")
	}

	prescription := domain.Prescription{
		DoctorID:   request.DoctorID,
		PatientID:  request.PatientID,
		Medication: request.Medication,
		Dosage:     request.Dosage,
	}

	save, err := service.prescriptionRepository.Save(prescription)
	if err != nil {
		return domain.Prescription{}, errors.New("an error occurred while processing your request")
	}
	return save, nil
}
