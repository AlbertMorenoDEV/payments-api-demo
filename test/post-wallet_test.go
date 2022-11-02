package test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserCreatesHisWalletSuccessfully(t *testing.T) {
	tests := []struct {
		testName string
		userId   string
		currency string
	}{
		{
			"creates his own wallet",
			"fb3a8c27-594d-493d-ad2b-8c0831a8443e",
			"USD",
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			setUp()

			body := []byte(`{
				"user_id": "` + test.userId + `",
				"currency": "` + test.currency + `"
			}`)

			res := executeRequest(
				http.MethodPost,
				fmt.Sprintf("/wallet/%s", test.userId),
				bytes.NewBuffer(body),
			)

			assert.Equal(t, http.StatusCreated, res.Code)
		})
	}
}
