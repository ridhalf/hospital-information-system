package domain

import "time"

// Tabel ini menyimpan informasi tentang janji temu antara pasien dan dokter.
type Appointment struct {
	Id              int       `json:"id"`
	DoctorId        int       `json:"doctor_id"`
	PatientId       int       `json:"patient_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Doctor          User      `json:"doctor"`
	Patient         Patient   `json:"patient"`
}
