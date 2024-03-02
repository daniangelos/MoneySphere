package models

import "time"

type Transaction struct {
	ID                     int
	ClientID               int
	TransactionValue       int
	TransactionType        string
	TransactionDescription string
	CreatedAt              time.Time
}
