package entity

import "time"

type Users struct {
	ID          string
	Username    string
	Email       string
	Password    string
	PhoneNumber int
	DateOfBirth time.Time
	Address     string
	Balance     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
