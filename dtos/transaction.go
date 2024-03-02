package dtos

import "github.com/MoneySphere/constants"

type TransactionCreateRequest struct {
	Value       int                `json:"valor"`
	Type        constants.TypeEnum `json:"tipo"`
	Description string             `json:"descricao"`
}
