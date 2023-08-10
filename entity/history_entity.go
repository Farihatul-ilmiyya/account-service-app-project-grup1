package entity

import "time"

type History struct {
	ID          string
	Username    string
	PhoneNumber string
	Amount      float64
	CreatedAt   time.Time
	IsSender    bool
}
