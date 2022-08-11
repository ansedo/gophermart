package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ansedo/gophermart/internal/helpers"
	"github.com/ansedo/gophermart/internal/middlewares"
	"github.com/ansedo/gophermart/internal/services/gophermart"
)

func GetWithdrawals(g *gophermart.Gophermart) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := middlewares.GetUserFromCtx(r.Context())
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}

		withdrawals, err := g.Storage.GetWithdrawalsByUID(r.Context(), user.ID)
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}
		res, err := json.Marshal(withdrawals)
		if err != nil {
			http.Error(w, err.Error(), helpers.GetStatusByError(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
