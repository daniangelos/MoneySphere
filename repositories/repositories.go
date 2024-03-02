package repositories

import "gorm.io/gorm"

type RepositoryContainer struct {
	ClientRepository      IClientRepository
	TransactionRepository ITransactionRepository
}

func NewRepositoryContainer(dbConn *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		ClientRepository:      NewClientRepository(dbConn),
		TransactionRepository: NewTransactionRepository(dbConn),
	}
}
