package ui

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/application"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	sharedApplication "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/application"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func PostTransactionHandler(
	repository domain.TransactionRepository,
	timeProvider sharedApplication.TimeProvider,
	domainEventPublisher sharedDomain.DomainEventPublisher,
	logger *zap.Logger,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["user_id"]

		var c application.CreateTransactionCommand

		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if userId != c.UserId {
			http.Error(w, "Invalid user", http.StatusForbidden)
			return
		}

		ch := application.NewCreateTransactionCommandHandler(repository, timeProvider, domainEventPublisher)

		err := ch.Handle(c)

		switch err.(type) {
		case nil:
			w.WriteHeader(http.StatusAccepted)
		case *transactionId.InvalidTransactionId:
		case *sharedUserId.InvalidUserId:
			w.WriteHeader(http.StatusBadRequest)
		case *domain.TransactionAlreadyExist:
			w.WriteHeader(http.StatusConflict)
		default:
			logger.Error("Unexpected error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
