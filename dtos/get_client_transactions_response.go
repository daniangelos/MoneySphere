package dtos

import (
	"time"

	"github.com/MoneySphere/constants"
	"github.com/MoneySphere/models"
)

type GetClientTransactionsResponse struct {
	Balance            ClientBalance       `json:"saldo"`
	LatestTransactions []ClientTransaction `json:"ultimas_transacoes"`
}

type ClientBalance struct {
	Total       int       `json:"total"`
	CurrentTime time.Time `json:"data_extrato"`
	Limit       int       `json:"limite"`
}

type ClientTransaction struct {
	Value       int                `json:"valor"`
	Type        constants.TypeEnum `json:"tipo"`
	Description string             `json:"descricao"`
	CreatedAt   time.Time          `json:"realizada_em"`
}

func ConvertToGetClientTransactionsReponse(client *models.Client, transactions []*models.Transaction) GetClientTransactionsResponse {
	var latestTransactions []ClientTransaction

	for _, transaction := range transactions {
		clientTransaction := ClientTransaction{
			Value:       transaction.TransactionValue,
			Type:        transaction.TransactionType,
			Description: transaction.TransactionDescription,
			CreatedAt:   transaction.CreatedAt,
		}

		latestTransactions = append(latestTransactions, clientTransaction)
	}

	return GetClientTransactionsResponse{
		Balance: ClientBalance{
			Total:       client.Balance,
			CurrentTime: time.Now(),
			Limit:       client.AccountLimit,
		},
		LatestTransactions: latestTransactions,
	}
}
