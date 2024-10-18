package domain

import "time"

type PharmacyTransaction struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"` // Primary key with auto-increment
	PrescriptionID int       `json:"prescription_id" gorm:"not null"`    // Foreign key referencing the Prescriptions table
	MedicationID   int       `json:"medication_id" gorm:"not null"`      // Foreign key referencing the Medications table
	Quantity       int       `json:"quantity" gorm:"not null"`           // Quantity of medication dispensed
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`   // Timestamp for creation time
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`   // Timestamp for creation time
}
