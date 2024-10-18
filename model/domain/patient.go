package domain

import (
	"time"
)

type Patient struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"` // Primary key with auto-increment
	UserID      int       `json:"user_id" gorm:"not null"`            // Foreign key referencing the Users table
	Name        string    `json:"name" gorm:"size:100;not null"`      // Patient's name with a maximum length of 100
	DateOfBirth time.Time `json:"date_of_birth" gorm:"not null"`      // Date of birth of the patient
	Address     string    `json:"address" gorm:"type:text"`           // Patient's address
	Phone       string    `json:"phone" gorm:"size:15"`               // Patient's phone number with a maximum length of 15
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`   // Timestamp for creation time
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`   // Timestamp for last update
	User        User      `json:"user" gorm:"foreignKey:user_id"`
}
