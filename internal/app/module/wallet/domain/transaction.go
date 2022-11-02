package domain

import (
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"time"
)

const (
	transactionStatusPending  = "PENDING"
	transactionStatusRejected = "REJECTED"
	transactionStatusAccepted = "ACCEPTED"
)

type Transaction struct {
	transactionId     transactionId.TransactionId
	userId            sharedUserId.UserId
	originUserId      sharedUserId.UserId
	destinationUserId sharedUserId.UserId
	amount            money.Money
	creationDate      time.Time
	modificationDate  time.Time
	status            string
}

func (t Transaction) TransactionId() transactionId.TransactionId {
	return t.transactionId
}

func (t Transaction) UserId() sharedUserId.UserId {
	return t.userId
}

func (t Transaction) OriginUserId() sharedUserId.UserId {
	return t.originUserId
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

func (t Transaction) ModificationDate() time.Time {
	return t.modificationDate
}

func (t Transaction) Status() string {
	return t.status
}

func (t Transaction) verify(now time.Time) error {
	t.status = transactionStatusAccepted
	t.modificationDate = now
	return nil
}
