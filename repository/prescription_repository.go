package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type PrescriptionRepository interface {
	Save(prescription domain.Prescription) (domain.Prescription, error)
}
type PrescriptionRepositoryImpl struct {
	db *gorm.DB
}

func NewPrescriptionRepository(db *gorm.DB) PrescriptionRepository {
	return &PrescriptionRepositoryImpl{
		db: db,
	}
}

func (repository PrescriptionRepositoryImpl) Save(prescription domain.Prescription) (domain.Prescription, error) {
	err := repository.db.Save(&prescription).Error
	if err != nil {
		return domain.Prescription{}, err
	}
	return prescription, nil
}
