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
	repoContainer repositories.RepositoryContainer
}

func NewClientController(repositoryContainer repositories.RepositoryContainer) *ClientController {
	return &ClientController{
		repoContainer: repositoryContainer,
	}
}

func (cc *ClientController) CreateClientTransaction(id int, transactionData dtos.TransactionCreateRequest) (*models.Transaction, error) {
	client, err := cc.repoContainer.ClientRepository.GetClientByID(id)
	if err != nil {
		fmt.Println("failed to get client from db", err)
		return nil, err
	}

	err = client.UpdateBalance(transactionData.Type, transactionData.Value)
	if err != nil {
		return nil, err
	}

	err = cc.repoContainer.ClientRepository.UpdateClientBalance(id, client.Balance)
	if err != nil {
		return nil, err
	}

	transaction, err := cc.repoContainer.TransactionRepository.CreateTransaction(id, transactionData.Value, transactionData.Type, transactionData.Description)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
