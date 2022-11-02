package transaction_id

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/uuid"
)

type TransactionId struct {
	value string
}

func (ti TransactionId) Value() string {
	return ti.value
}

func New(value string) (*TransactionId, error) {
	if value == "" {
		return nil, NewInvalidEmptyTransactionId()
	}

	if uuid.IsValid(value) {
		return &TransactionId{value: value}, nil
	}

	return nil, NewInvalidReceivedTransactionId(value)
}
