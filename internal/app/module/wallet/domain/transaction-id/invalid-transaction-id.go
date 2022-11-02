package transaction_id

import "fmt"

type InvalidTransactionId struct {
	message string
}

func NewInvalidEmptyTransactionId() *InvalidTransactionId {
	return &InvalidTransactionId{
		message: "Transaction ID can not be empty.",
	}
}

func NewInvalidReceivedTransactionId(value string) *InvalidTransactionId {
	return &InvalidTransactionId{
		message: fmt.Sprintf("Invalid transaction ID. Received value: <%s>.", value),
	}
}

func (e InvalidTransactionId) Error() string {
	return e.message
}
