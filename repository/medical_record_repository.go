package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type MedicalRecordRepository interface {
	FindByPatientId(patientID int, withRelation bool) ([]domain.MedicalRecord, error)
}
type MedicalRecordRepositoryImpl struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository {
	return &MedicalRecordRepositoryImpl{
		db: db,
	}
}

func (repository MedicalRecordRepositoryImpl) FindByPatientId(patientID int, withRelation bool) ([]domain.MedicalRecord, error) {
	var medicalRecords []domain.MedicalRecord
	query := repository.db.Model(&medicalRecords).Where("patient_id = ?", patientID)
	if withRelation {
		query = query.Preload("Patient")
	}
	err := query.Find(&medicalRecords).Error
	if err != nil {
		return nil, err
	}
	return medicalRecords, nil
}
