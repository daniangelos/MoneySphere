package models

import (
	"time"

	"github.com/MoneySphere/constants"
)

type Transaction struct {
	ID                     int
	ClientID               int
	TransactionValue       int
	TransactionType        constants.TypeEnum
	TransactionDescription string
	CreatedAt              time.Time
}
