package repositories

import "github.com/MoneySphere/models"

type IClientRepository interface {
	GetClientByID(id int) (*models.Client, error)
	UpdateClientBalance(id int, newBalance int) error
}

type ITransactionRepository interface {
	CreateTransaction(clientID int, value int, transactionType string, description string) (*models.Transaction, error)
	GetTransactionsByClientID() ([]*models.Transaction, error)
}
