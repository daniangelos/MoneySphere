package repositories

import (
	"github.com/MoneySphere/constants"
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

func (tr *TransactionRepository) CreateTransaction(clientID int, value int, transactionType constants.TypeEnum, description string) (*models.Transaction, error) {
	transaction := models.Transaction{
		ClientID:               clientID,
		TransactionValue:       value,
		TransactionType:        transactionType,
		TransactionDescription: description,
	}

	err := tr.dbConn.Create(&transaction).Error

	return &transaction, err
}

func (tr *TransactionRepository) GetTransactionsByClientID(clientID int) ([]*models.Transaction, error) {
	transactions := []*models.Transaction{}

	err := tr.dbConn.Find(&transactions).Where("client_id = ?", clientID).Error

	return transactions, err
}
