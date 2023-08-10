package entity

import "time"

type HistoryTopUp struct {
	Id        string
	Username  string
	Amount    float64
	CreatedAt time.Time
}
