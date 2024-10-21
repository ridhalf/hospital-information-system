package domain

import "time"

type MedicalRecord struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"` // Primary key with auto-increment
	PatientID  int       `json:"patient_id" gorm:"not null"`         // Foreign key referencing the Patients table
	Diagnosis  string    `json:"diagnosis" gorm:"type:text"`         // Diagnosis of the patient
	Treatment  string    `json:"treatment" gorm:"type:text"`         // Treatment given to the patient
	Medication string    `json:"medication" gorm:"type:text"`        // Medication prescribed to the patient
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Patient    Patient   `json:"patient" gorm:"foreignKey:patient_id"`
}

func (MedicalRecord) TableName() string {
	return "MedicalRecords"
}
