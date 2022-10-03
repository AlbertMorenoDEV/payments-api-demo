package user_id

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/uuid"
)

type UserId struct {
	value string
}

func (ui UserId) Value() string {
	return ui.value
}

func New(value string) (*UserId, error) {
	if value == "" {
		return nil, NewInvalidEmptyUserId()
	}

	if uuid.IsValid(value) {
		return &UserId{value: value}, nil
	}

	return nil, NewInvalidReceivedUserId(value)
}
