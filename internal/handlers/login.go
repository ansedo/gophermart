package handlers

import (
	"fmt"
	"net/http"

	"github.com/ansedo/gophermart/internal/auth"
	"github.com/ansedo/gophermart/internal/helpers"
	"github.com/ansedo/gophermart/internal/middlewares"
	"github.com/ansedo/gophermart/internal/models"
	"github.com/ansedo/gophermart/internal/services/gophermart"
)

func Login(g *gophermart.Gophermart) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := middlewares.GetUserFromCtx(r.Context())
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}

		dbUser, err := g.Storage.GetUserByLogin(r.Context(), user)
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}
		if !auth.Authenticate(dbUser, user) {
			helpers.HTTPError(w, models.ErrInvalidLoginAttempt)
			return
		}

		token, err := auth.GenerateToken(user)
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
		w.WriteHeader(http.StatusOK)
	}
}
