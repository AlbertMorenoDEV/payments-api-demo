package domain

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

type Balance struct {
	userId sharedUserId.UserId
	amount money.Money
}

func (b *Balance) UserId() sharedUserId.UserId {
	return b.userId
}

func (b *Balance) Amount() money.Money {
	return b.amount
}

func (b *Balance) Update(amount money.Money) {
	b.amount = b.amount.Add(amount)
}

func NewBalance(userId sharedUserId.UserId, amount money.Money) *Balance {
	return &Balance{userId: userId, amount: amount}
}

func NewBalanceFromPrimitives(stringUserId string, amount int64, currency string) (*Balance, error) {
	userId, err := sharedUserId.New(stringUserId)
	if err != nil {
		return nil, err
	}

	return &Balance{
		userId: *userId,
		amount: money.NewFromPrimitives(amount, currency),
	}, nil
}
