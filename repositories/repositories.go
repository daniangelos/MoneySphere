package repositories

type RepositoryContainer struct {
	ClientRepository      IClientRepository
	TransactionRepository ITransactionRepository
}

func NewRepositoryContainer() RepositoryContainer {
	return RepositoryContainer{
		ClientRepository:      NewClientRepository(),
		TransactionRepository: NewTransactionRepository(),
	}
}
