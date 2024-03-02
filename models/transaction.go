package models

import "time"

type Transaction struct {
	ID          int
	ClientID    int
	Value       int
	Type        string
	Description string
	CreatedAt   time.Time
}
