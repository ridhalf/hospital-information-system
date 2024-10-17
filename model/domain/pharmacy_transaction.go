package domain

import "time"

type PharmacyTransaction struct {
	Id             int          `json:"id"`
	PrescriptionId int          `json:"prescription_id"`
	MedicationId   int          `json:"medication_id"`
	Quantity       int          `json:"quantity"`
	CreatedAt      time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	Prescription   Prescription `json:"prescription"`
	Medication     Medication   `json:"medication"`
}
