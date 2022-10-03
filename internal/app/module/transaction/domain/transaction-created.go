package domain

import (
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	sharedMoney "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"time"
)

type TransactionCreated struct {
	eventId           sharedDomain.Uuid
	transactionId     transactionId.TransactionId
	userId            sharedUserId.UserId
	destinationUserId sharedUserId.UserId
	amount            sharedMoney.Money
	creationDate      time.Time
}

func (e TransactionCreated) Id() sharedDomain.Uuid {
	return e.eventId
}

func (e TransactionCreated) AggregateId() sharedDomain.Uuid {
	return sharedDomain.Uuid(e.transactionId.Value())
}

func (e TransactionCreated) Name() string {
	return "payments_api_demo.transaction_created"
}

func (e TransactionCreated) CreatedAt() time.Time {
	return e.creationDate
}

func (e TransactionCreated) Data() map[string]interface{} {
	return map[string]interface{}{
		"transaction_id":      e.transactionId.Value(),
		"user_id":             e.userId.Value(),
		"destination_user_id": e.destinationUserId.Value(),
		"amount":              e.amount.Amount().Int64(),
		"currency":            e.amount.Currency().String(),
	}
}

func NewTransactionCreated(transaction Transaction) *TransactionCreated {
	return &TransactionCreated{
		eventId:           sharedDomain.NewUuid(),
		transactionId:     transaction.TransactionId(),
		userId:            transaction.UserId(),
		destinationUserId: transaction.DestinationUserId(),
		amount:            transaction.Amount(),
		creationDate:      transaction.CreationDate(),
	}
}
