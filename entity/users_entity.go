package entity

import "time"

type Users struct {
	ID          string
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	DateOfBirth time.Time
	Address     string
	Balance     float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
