package ui

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/wallet/application"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/command"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func PostWalletHandler(
	commandBus command.Bus,
	logger *zap.Logger,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["user_id"]

		var c application.CreateWalletCommand

		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if userId != c.UserId {
			http.Error(w, "Invalid user", http.StatusForbidden)
			return
		}

		err := commandBus.Dispatch(r.Context(), c)

		switch err.(type) {
		case nil:
			w.WriteHeader(http.StatusAccepted)
		case *sharedUserId.InvalidUserId:
			w.WriteHeader(http.StatusBadRequest)
		case *domain.WalletAlreadyExist:
			w.WriteHeader(http.StatusConflict)
		default:
			logger.Error("Unexpected error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
