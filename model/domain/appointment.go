package domain

import "time"

type Appointment struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`                                                // Primary key with auto-increment
	DoctorID        int       `json:"doctor_id" gorm:"not null"`                                                         // Foreign key referencing the Users table (doctor)
	PatientID       int       `json:"patient_id" gorm:"not null"`                                                        // Foreign key referencing the Patients table
	AppointmentDate string    `json:"appointment_date" gorm:"not null"`                                                  // Appointment date
	AppointmentTime string    `json:"appointment_time" gorm:"not null"`                                                  // Appointment time
	Status          string    `json:"status" gorm:"type:ENUM('scheduled', 'completed', 'canceled');default:'scheduled'"` // Status of the appointment
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Timestamp for creation time
}
