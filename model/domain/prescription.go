package domain

import "time"

//resep dokter yang dibuat dokter untuk pasien

type Prescription struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"` // Primary key with auto-increment
	DoctorID   int       `json:"doctor_id" gorm:"not null"`          // Foreign key referencing the Users table (doctor)
	PatientID  int       `json:"patient_id" gorm:"not null"`         // Foreign key referencing the Patients table
	Medication string    `json:"medication" gorm:"type:text"`        // Medication prescribed
	Dosage     string    `json:"dosage" gorm:"type:text"`            // Dosage instructions
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`   // Timestamp for creation time
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`   // Timestamp for creation time
}
