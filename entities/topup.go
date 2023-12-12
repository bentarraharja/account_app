package entities

import "time"

type TopUp struct {
	ID        int
	AccountID int
	Amount    int
	CreatedAt time.Time
}
