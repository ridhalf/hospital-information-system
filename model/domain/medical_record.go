package domain

import "time"

type MedicalRecord struct {
	Id         int       `json:"id"`
	PatientId  int       `json:"patient_id"`
	Diagnosis  string    `json:"diagnosis"`
	Treatment  string    `json:"treatment"`
	Medication string    `json:"medication"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Patient    Patient   `json:"patient"`
}
