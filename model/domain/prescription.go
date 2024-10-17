package domain

import "time"

//resep dokter yang dibuat dokter untuk pasien

type Prescription struct {
	Id         int       `json:"id"`
	DoctorId   int       `json:"doctor_id"`
	PatientId  int       `json:"patient_id"`
	Medication string    `json:"medication"`
	Dosage     string    `json:"dosage"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Doctor     User      `json:"doctor"`
	Patient    Patient   `json:"patient"`
}
