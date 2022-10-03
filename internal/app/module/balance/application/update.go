package application

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

const eventToListen = "payments_api_demo.transaction_created"

type UpdateBalanceCommand struct {
	UserId   string
	Amount   int64
	Currency string
}

type UpdateBalanceCommandHandler struct {
	repository domain.BalanceRepository
}

func (h UpdateBalanceCommandHandler) Handle(event sharedDomain.DomainEvent) error {
	if event.Name() != eventToListen {
		return nil
	}

	stringUserId := event.Data()["destination_user_id"].(string)

	userId, err := sharedUserId.New(stringUserId)
	if err != nil {
		return err
	}

	balance, err := h.repository.Find(*userId)
	if err != nil {
		return err
	}

	int64Amount := event.Data()["amount"].(int64)
	stringCurrency := event.Data()["currency"].(string)

	balance.Update(money.NewFromPrimitives(int64Amount, stringCurrency))

	return h.repository.Save(balance)
}

func NewUpdateBalanceCommandHandler(repository domain.BalanceRepository) *UpdateBalanceCommandHandler {
	return &UpdateBalanceCommandHandler{repository: repository}
}
