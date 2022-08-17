package handlers

import (
	"fmt"
	"net/http"

	"github.com/ansedo/gophermart/internal/auth"
	"github.com/ansedo/gophermart/internal/helpers"
	"github.com/ansedo/gophermart/internal/middlewares"
	"github.com/ansedo/gophermart/internal/services/gophermart"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(g *gophermart.Gophermart) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := middlewares.GetUserFromCtx(r.Context())
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			helpers.HTTPError(w, err)
			return
		}

		user.ID = uuid.NewString()
		user.Password = string(hashedPass)

		if err = g.Storage.AddUser(r.Context(), user); err != nil {
			helpers.HTTPError(w, err)
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
