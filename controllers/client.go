package controllers

import (
	"fmt"

	"github.com/MoneySphere/dtos"
	"github.com/MoneySphere/models"
	"github.com/MoneySphere/repositories"
)

type IClientController interface {
	CreateClientTransaction(id int, transactionData dtos.TransactionCreateRequest) (*models.Transaction, error)
}

type ClientController struct {
	clientRepository      repositories.IClientRepository
	transactionRepository repositories.ITransactionRepository
}

func NewClientController(repositoryContainer repositories.RepositoryContainer) *ClientController {
	return &ClientController{
		clientRepository:      repositoryContainer.ClientRepository,
		transactionRepository: repositoryContainer.TransactionRepository,
	}
}

func (cc *ClientController) CreateClientTransaction(id int, transactionData dtos.TransactionCreateRequest) (*models.Transaction, error) {
	client, err := cc.clientRepository.GetClientByID(id)
	if err != nil {
		fmt.Println("failed to get client from db", err)
		return nil, err
	}

	err = client.UpdateBalance(transactionData.Type, transactionData.Value)

	return nil, nil
}
