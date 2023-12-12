package entities

import "time"

type Account struct {
	ID        int
	FullName  string
	Address   string
	Phone     string
	Email     string
	Password  string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
