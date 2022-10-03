package test

import (
	"fmt"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserRetrieveHisBalanceSuccessfully(t *testing.T) {
	tests := []struct {
		testName         string
		userId           string
		expectedResponse string
	}{
		{
			"empty wallet",
			"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
			`{
                "amount": 0,
				"currency": "SGD"
            }`,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			setUp()

			res := executeRequest(
				http.MethodGet,
				fmt.Sprintf("/users/%s/balance", test.userId),
				nil,
			)

			assert.Equal(t, http.StatusOK, res.Code)
			jsonassert.New(t).Assertf(res.Body.String(), test.expectedResponse)
		})
	}
}

func TestGetUserBalanceWrongMethod(t *testing.T) {
	tests := []struct {
		testName string
		userId   string
		method   string
	}{
		{
			"empty wallet",
			"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
			http.MethodPost,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res := executeRequest(
				test.method,
				fmt.Sprintf("/users/%s/balance", test.userId),
				nil,
			)

			assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
		})
	}
}
