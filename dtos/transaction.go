package dtos

type TransactionCreateRequest struct {
	Value           int    `json:"valor"`
	TransactionType string `json:"tipo"`
	Description     string `json:"descricao"`
}
