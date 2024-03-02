package repositories

import "github.com/MoneySphere/models"

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (tr *TransactionRepository) CreateTransaction(clientID int, value int, transactionType string, description string) (*models.Transaction, error) {
	return nil, nil
}

func (tr *TransactionRepository) GetTransactionsByClientID() ([]*models.Transaction, error) {
	return nil, nil
}
