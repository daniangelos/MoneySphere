package dtos

import "github.com/MoneySphere/constants"

type TransactionCreateRequest struct {
	Value       int                `json:"valor" validate:"required,gt=0"`
	Type        constants.TypeEnum `json:"tipo" validate:"required,oneof=c d"`
	Description string             `json:"descricao" validate:"required,min=1,max=10"`
}
