package domain

import (
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

type BalanceRepository interface {
	Find(userId sharedUserId.UserId) (*Balance, error)
	Save(balance *Balance) error
}
