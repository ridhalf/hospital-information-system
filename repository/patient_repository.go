package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type PatientRepository interface {
	Save(patient domain.Patient) (domain.Patient, error)
	FindById(id int, withRelation bool) (domain.Patient, error)
	FindByUserId(userId int, withRelation bool) (domain.Patient, error)
	Count(ID int) (int64, error)
}
type PatientRepositoryImpl struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &PatientRepositoryImpl{
		db: db,
	}
}

func (repository PatientRepositoryImpl) FindByUserId(userId int, withRelation bool) (domain.Patient, error) {
	patient := domain.Patient{}
	query := repository.db.Model(&patient).Where("user_id = ?", userId)
	if withRelation {
		query = query.Preload("User")
	}
	err := query.First(&patient).Error
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (repository PatientRepositoryImpl) Save(patient domain.Patient) (domain.Patient, error) {
	err := repository.db.Save(&patient).Error
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (repository PatientRepositoryImpl) FindById(id int, withRelation bool) (domain.Patient, error) {
	patient := domain.Patient{}
	query := repository.db.Model(&patient).Where("id = ?", id)
	if withRelation {
		query = query.Preload("User")
	}
	err := query.First(&patient).Error
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}
func (repository PatientRepositoryImpl) Count(ID int) (int64, error) {
	var count int64
	err := repository.db.Model(&domain.Patient{}).Where("id = ?", ID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
