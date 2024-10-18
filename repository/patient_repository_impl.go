package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type PatientRepositoryImpl struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &PatientRepositoryImpl{
		db: db,
	}
}
func (repository PatientRepositoryImpl) Save(patient domain.Patient) (domain.Patient, error) {
	err := repository.db.Save(&patient).Error
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (repository PatientRepositoryImpl) FindById(id int) (domain.Patient, error) {
	patient := domain.Patient{}
	err := repository.db.
		Joins("JOIN users u ON u.id = patients.user_id").
		Select("patients.*, u.username, u.email, u.role").
		Where("patients.id = ?", id).
		First(&patient).Error
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}
