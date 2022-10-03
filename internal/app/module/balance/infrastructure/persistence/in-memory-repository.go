package persistence

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

type InMemoryRepository struct {
	items map[sharedUserId.UserId]*domain.Balance
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		items: map[sharedUserId.UserId]*domain.Balance{},
	}
}

func (r *InMemoryRepository) Find(userId sharedUserId.UserId) (*domain.Balance, error) {
	item, ok := r.items[userId]
	if ok {
		return item, nil
	}

	return domain.NewBalance(userId, money.NewSGD(money.NewZeroAmount())), nil
}
