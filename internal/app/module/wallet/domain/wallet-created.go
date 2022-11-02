package domain

import (
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"time"
)

type WalletCreated struct {
	eventId      sharedDomain.Uuid
	userId       sharedUserId.UserId
	currency     money.Currency
	creationDate time.Time
}

func (e WalletCreated) Id() string {
	return e.eventId.String()
}

func (e WalletCreated) AggregateId() string {
	return e.userId.Value()
}

func (e WalletCreated) Currency() money.Currency {
	return e.currency
}

func (e WalletCreated) Name() string {
	return "payments_api_demo.wallet_created"
}

func (e WalletCreated) CreatedAt() time.Time {
	return e.creationDate
}

func (e WalletCreated) Data() map[string]interface{} {
	return map[string]interface{}{
		"user_id":  e.userId.Value(),
		"currency": e.currency.String(),
	}
}

func NewWalletCreated(userId sharedUserId.UserId, currency money.Currency, now time.Time) *WalletCreated {
	return &WalletCreated{
		eventId:      sharedDomain.NewUuid(),
		userId:       userId,
		currency:     currency,
		creationDate: now,
	}
}
