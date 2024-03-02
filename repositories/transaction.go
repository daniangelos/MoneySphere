package repositories

import (
	"github.com/MoneySphere/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	dbConn *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		dbConn: dbConn,
	}
}

func (tr *TransactionRepository) CreateTransaction(clientID int, value int, transactionType string, description string) (*models.Transaction, error) {
	return nil, nil
}

func (tr *TransactionRepository) GetTransactionsByClientID() ([]*models.Transaction, error) {
	return nil, nil
}
