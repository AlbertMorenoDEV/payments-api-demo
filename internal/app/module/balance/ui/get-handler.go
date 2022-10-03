package ui

import (
	"encoding/json"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/application"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func GetUserBalanceHandler(repository domain.BalanceRepository, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["user_id"]
		q := application.FindBalance{UserId: userId}
		qh := application.NewFindBalanceHandler(repository)

		res, err := qh.Handle(q)

		switch err.(type) {
		case nil:
			b, err := json.Marshal(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				if _, err := fmt.Fprintf(w, err.Error()); err != nil {
					panic(err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)

			_, err = fmt.Fprintf(w, string(b))
			if err != nil {
				panic(err)
			}

		default:
			logger.Error("Unexpected error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)

			if _, err := fmt.Fprintf(w, err.Error()); err != nil {
				panic(err)
			}
		}
	}
}
