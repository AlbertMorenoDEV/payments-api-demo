package persistence

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
)

type InMemoryRepository struct {
	items map[transactionId.TransactionId]*domain.Transaction
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		items: map[transactionId.TransactionId]*domain.Transaction{},
	}
}

func (r *InMemoryRepository) Save(transaction *domain.Transaction) error {
	r.items[transaction.TransactionId()] = transaction

	return nil
}
