package user_id

import "fmt"

type InvalidUserId struct {
	message string
}

func NewInvalidEmptyUserId() *InvalidUserId {
	return &InvalidUserId{
		message: "User ID can not be empty.",
	}
}

func NewInvalidReceivedUserId(value string) *InvalidUserId {
	return &InvalidUserId{
		message: fmt.Sprintf("Invalid user ID. Received value: <%s>.", value),
	}
}

func (e InvalidUserId) Error() string {
	return e.message
}
