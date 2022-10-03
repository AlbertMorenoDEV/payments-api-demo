package domain

import (
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"time"
)

type Transaction struct {
	transactionId     transactionId.TransactionId
	userId            sharedUserId.UserId
	destinationUserId sharedUserId.UserId
	amount            money.Money
	creationDate      time.Time

	eventRecorder *sharedDomain.EventRecorder
}

func (t Transaction) TransactionId() transactionId.TransactionId {
	return t.transactionId
}

func (t Transaction) UserId() sharedUserId.UserId {
	return t.userId
}

func (t Transaction) DestinationUserId() sharedUserId.UserId {
	return t.destinationUserId
}

func (t Transaction) Amount() money.Money {
	return t.amount
}

func (t Transaction) CreationDate() time.Time {
	return t.creationDate
}

func (t Transaction) PullDomainEvents() sharedDomain.DomainEvents {
	return t.eventRecorder.Pull()
}

func NewTransaction(
	transactionId transactionId.TransactionId,
	userId sharedUserId.UserId,
	destinationUserId sharedUserId.UserId,
	amount money.Money,
	creationDate time.Time,
) *Transaction {
	t := &Transaction{
		transactionId:     transactionId,
		userId:            userId,
		destinationUserId: destinationUserId,
		amount:            amount,
		creationDate:      creationDate,
		eventRecorder:     sharedDomain.InitEventRecorder(),
	}

	t.eventRecorder.Record(NewTransactionCreated(*t))

	return t
}
