package service

import (
	"errors"
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
	"hospital-information-system/repository"
)

type AppointmentService interface {
	CreateSchedule(request web.AppointmentCreateScheduleRequest) (domain.Appointment, error)
}
type AppointmentServiceImpl struct {
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(appointmentRepository repository.AppointmentRepository) AppointmentService {
	return &AppointmentServiceImpl{
		appointmentRepository: appointmentRepository,
	}
}

func (service AppointmentServiceImpl) CreateSchedule(request web.AppointmentCreateScheduleRequest) (domain.Appointment, error) {
	appointment := domain.Appointment{
		DoctorID:        request.DoctorID,
		PatientID:       request.PatientID,
		AppointmentDate: request.AppointmentDate,
		AppointmentTime: request.AppointmentTime,
		Status:          request.Status,
	}
	//check schedule
	count, err := service.appointmentRepository.Count(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("failed to count appointments")
	}
	if count > 0 {
		return domain.Appointment{}, errors.New("appointment already exists")
	}

	save, err := service.appointmentRepository.Save(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("failed to save appointment")
	}
	return save, nil
}
