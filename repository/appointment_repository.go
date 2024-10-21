package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type AppointmentRepository interface {
	Save(appointment domain.Appointment) (domain.Appointment, error)
	Count(appointment domain.Appointment) (int64, error)
}
type AppointmentRepositoryImpl struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &AppointmentRepositoryImpl{
		db: db,
	}
}

func (repository AppointmentRepositoryImpl) Count(appointment domain.Appointment) (int64, error) {
	var count int64
	err := repository.db.Model(&domain.Appointment{}).Where("doctor_id = ? AND status = ? AND appointment_date = ? AND appointment_time = ?",
		appointment.DoctorID, appointment.Status, appointment.AppointmentDate, appointment.AppointmentTime).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repository AppointmentRepositoryImpl) Save(appointment domain.Appointment) (domain.Appointment, error) {
	err := repository.db.Save(&appointment).Error
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}
