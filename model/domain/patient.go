package domain

import "time"

type Patient struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User        User      `json:"user"`
}
