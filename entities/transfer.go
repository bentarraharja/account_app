package entities

import "time"

type Transfer struct {
	ID                int
	AccountIdSender   int
	AccountIdReceiver int
	Amount            int
	CreatedAt         time.Time
}
