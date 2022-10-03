package application

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

type FindBalance struct {
	UserId string
}

func (q FindBalance) Id() string {
	return "find_balance"
}

type FindBalanceHandler struct {
	repository domain.BalanceRepository
}

func NewFindBalanceHandler(repository domain.BalanceRepository) FindBalanceHandler {
	return FindBalanceHandler{repository: repository}
}

func (h FindBalanceHandler) Handle(query FindBalance) (*BalanceResponse, error) {
	userId, err := sharedUserId.New(query.UserId)
	if err != nil {
		return nil, err
	}

	balance, err := h.repository.Find(*userId)
	if err != nil {
		return nil, err
	}

	return NewResponse(balance), nil
}
