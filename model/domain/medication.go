package domain

import "time"

type Medication struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`     // Primary key with auto-increment
	Name      string    `json:"name" gorm:"not null;type:varchar(100)"` // Name of the medication
	Stock     int       `json:"stock" gorm:"not null"`                  // Stock quantity of the medication
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`       // Timestamp for creation time
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`       // Timestamp for creation time

}
