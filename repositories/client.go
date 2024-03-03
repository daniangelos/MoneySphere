package repositories

import (
	"github.com/MoneySphere/models"
	"gorm.io/gorm"
)

type ClientRepository struct {
	dbConn *gorm.DB
}

func NewClientRepository(dbConn *gorm.DB) *ClientRepository {
	return &ClientRepository{
		dbConn: dbConn,
	}
}

func (cr *ClientRepository) GetClientByID(id int) (*models.Client, error) {
	var client models.Client

	err := cr.dbConn.First(&client, id).Error

	return &client, err
}

func (cr *ClientRepository) UpdateClientBalance(id int, newBalance int) error {
	return cr.dbConn.Model(&models.Client{}).Where("id = ?", id).Update("balance", newBalance).Error
}
