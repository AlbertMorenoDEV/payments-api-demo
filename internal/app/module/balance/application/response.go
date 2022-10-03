package application

import balance_domain "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"

type BalanceResponse struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func NewResponse(balance *balance_domain.Balance) *BalanceResponse {
	return &BalanceResponse{
		Amount:   balance.Amount().Amount().Int64(),
		Currency: balance.Amount().Currency().String(),
	}
}
