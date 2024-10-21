package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`                                                                                                   // Primary key with auto-increment
	Username  string    `json:"username" gorm:"size:50;not null"`                                                                                                     // Username with a maximum length of 50
	Password  string    `json:"password" gorm:"size:255;not null"`                                                                                                    // Hashed password with a maximum length of 255
	Email     string    `json:"email" gorm:"size:100;unique;not null"`                                                                                                // Email must be unique and not null
	Role      string    `json:"role" gorm:"type:ENUM('patient', 'doctor', 'nurse', 'admin', 'receptionist', 'pharmacist', 'lab_technician', 'radiologist');not null"` // Role with enum type
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`                                                                                                     // Timestamp for creation time
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`                                                                                                     // Timestamp for creation time
}

func (User) TableName() string {
	return "Users"
}
