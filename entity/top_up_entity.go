package entity

import "time"

type TopUp struct {
	ID        string
	UserID    string
	Amount    float64
	CreatedAt time.Time
}
