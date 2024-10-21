package service

import (
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
	"hospital-information-system/repository"
)

type MedicalRecordService interface {
	FindByPatientID(request web.MedicalRecordFindByPatientIDRequest) ([]domain.MedicalRecord, error)
}
type MedicalRecordServiceImpl struct {
	medicalRecordRepository repository.MedicalRecordRepository
}

func NewMedicalRecordService(medicalRecordRepository repository.MedicalRecordRepository) MedicalRecordService {
	return &MedicalRecordServiceImpl{
		medicalRecordRepository: medicalRecordRepository,
	}
}

func (service MedicalRecordServiceImpl) FindByPatientID(request web.MedicalRecordFindByPatientIDRequest) ([]domain.MedicalRecord, error) {
	medicalRecords, err := service.medicalRecordRepository.FindByPatientId(request.PatientID, true)
	if err != nil {
		return nil, err
	}
	return medicalRecords, nil
}
