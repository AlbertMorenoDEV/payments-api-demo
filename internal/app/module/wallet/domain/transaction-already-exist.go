package domain

import (
	"fmt"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
)

type TransactionAlreadyExist struct {
	transactionId transactionId.TransactionId
}

func NewTransactionAlreadyExist(transactionId transactionId.TransactionId) *TransactionAlreadyExist {
	return &TransactionAlreadyExist{transactionId: transactionId}
}

func (e TransactionAlreadyExist) Error() string {
	return fmt.Sprintf("Transaction <%s> already exist", e.transactionId)
}
