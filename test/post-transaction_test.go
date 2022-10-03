package test

import (
	"bytes"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserSendMoneyToAnotherUserSuccessfully(t *testing.T) {
	tests := []struct {
		testName          string
		transactionId     string
		userId            string
		destinationUserId string
		amount            money.Amount
	}{
		{
			"send money to another user",
			"4cdad820-9cb9-4999-a0a6-26c5119df25b",
			"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
			"d543ef31-92c1-458f-8337-ca1ac16eb38d",
			money.Amount(1200),
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			setUp()

			body := []byte(`{
				"transaction_id": "` + test.transactionId + `",
				"user_id": "` + test.userId + `",
				"destination_user_id": "` + test.destinationUserId + `",
                "amount": ` + test.amount.String() + `,
				"currency": "SGD"
			}`)

			res := executeRequest(
				http.MethodPost,
				fmt.Sprintf("/users/%s/transactions", test.userId),
				bytes.NewBuffer(body),
			)

			assert.Equal(t, http.StatusAccepted, res.Code)
		})
	}
}

func TestUserSendMoneyToAnotherUserFail(t *testing.T) {
	tests := []struct {
		testName           string
		transactionId      string
		pathUserId         string
		userId             string
		destinationUserId  string
		amount             money.Amount
		expectedStatusCode int
	}{
		{
			"wrong origin user",
			"4cdad820-9cb9-4999-a0a6-26c5119df25b",
			"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
			"a0bb981a-5ad5-4886-baa0-89004c7f996a",
			"d543ef31-92c1-458f-8337-ca1ac16eb38d",
			money.Amount(1200),
			http.StatusForbidden,
		},
		// ToDo: Find out why this test is failing
		//{
		//	"empty transaction id",
		//	"",
		//	"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
		//	"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
		//	"d543ef31-92c1-458f-8337-ca1ac16eb38d",
		//	money.Amount(1200),
		//	http.StatusBadRequest,
		//},
		{
			"invalid user id",
			"4cdad820-9cb9-4999-a0a6-26c5119df25b",
			"wrong-uuid",
			"wrong-uuid",
			"d543ef31-92c1-458f-8337-ca1ac16eb38d",
			money.Amount(1200),
			http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			setUp()

			body := []byte(`{
				"transaction_id": "` + test.transactionId + `",
				"user_id": "` + test.userId + `",
				"destination_user_id": "` + test.destinationUserId + `",
                "amount": ` + test.amount.String() + `,
				"currency": "SGD"
			}`)

			res := executeRequest(
				http.MethodPost,
				fmt.Sprintf("/users/%s/transactions", test.pathUserId),
				bytes.NewBuffer(body),
			)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}
