package entity

import "time"

type Transfer struct {
	ID              string
	UserIdSender    string
	UserIdRecipient string
	Amount          float64
	CreatedAt       time.Time
}
