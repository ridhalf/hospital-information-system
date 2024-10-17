package domain

import "time"

type Medication struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
