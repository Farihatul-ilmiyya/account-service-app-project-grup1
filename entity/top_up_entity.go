package entity

import "time"

type TopUp struct {
	ID        string
	UserID    string
	Amount    int
	CreatedAt time.Time
}
