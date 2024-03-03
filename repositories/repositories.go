package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	ClientRepository      IClientRepository
	TransactionRepository ITransactionRepository
	dbConn                *gorm.DB
}

func NewRepositoryContainer(dbConn *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		ClientRepository:      NewClientRepository(dbConn),
		TransactionRepository: NewTransactionRepository(dbConn),
		dbConn:                dbConn,
	}
}

func (r *RepositoryContainer) WithDBTransaction(f func(repositoryContainer RepositoryContainer) error) (err error) {
	db := r.dbConn
	tx := db.Begin()
	txContainer := NewRepositoryContainer(tx)

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("panic during transaction", rec)
			tx.Rollback()
			panic(rec)
		}

		if err == nil {
			tx.Commit()
		}

		if err != nil {
			fmt.Println("error during transaction", err)
			tx.Rollback()
			return
		}
	}()

	err = f(txContainer)

	return err
}
