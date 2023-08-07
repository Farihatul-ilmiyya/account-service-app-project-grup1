package entity

import "time"

type Transfer struct {
	ID              string
	UserIdSender    string
	UserIdRecipient string
	Amount          int
	CreatedAt       time.Time
}
