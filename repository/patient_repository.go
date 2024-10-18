package repository

import "hospital-information-system/model/domain"

type PatientRepository interface {
	Save(patient domain.Patient) (domain.Patient, error)
	FindById(id int, withRelation bool) (domain.Patient, error)
}
