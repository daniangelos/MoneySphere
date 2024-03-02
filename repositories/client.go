package repositories

import "github.com/MoneySphere/models"

type ClientRepository struct {
}

func NewClientRepository() ClientRepository {
	return ClientRepository{}
}

func (cr *ClientRepository) GetClientByID(id int) (*models.Client, error) {
	return nil, nil
}

func (cr *ClientRepository) UpdateClientBalance(id int, newBalance int) error {
	return nil
}
