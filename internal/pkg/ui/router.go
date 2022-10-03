package ui

import (
	balanceUi "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/ui"
	transactionUi "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/ui"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/infrastructure"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildRouter(services *infrastructure.Services) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(
		"/users/{user_id}/balance",
		balanceUi.GetUserBalanceHandler(services.BalanceRepository(), services.Logger()),
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/users/{user_id}/transactions",
		transactionUi.PostTransactionHandler(
			services.TransactionRepository(),
			services.TimeProvider(),
			services.DomainEventPublisher(),
			services.Logger(),
		),
	).Methods(http.MethodPost)

	return r
}
