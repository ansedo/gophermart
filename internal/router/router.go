package router

import (
	"context"

	"github.com/ansedo/gophermart/internal/handlers"
	"github.com/ansedo/gophermart/internal/middlewares"
	"github.com/ansedo/gophermart/internal/services/gophermart"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(ctx context.Context) chi.Router {
	r := chi.NewRouter()
	g := gophermart.New(ctx)

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Compress(5),
		middlewares.Decompress,
	)
	r.Route("/api/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.TokenValidation())

			r.Route("/balance", func(r chi.Router) {
				r.Get("/", handlers.GetBalance(g))
				r.Post("/withdraw", handlers.Withdraw(g))
			})
			r.Route("/orders", func(r chi.Router) {
				r.Get("/", handlers.GetOrders(g))
				r.Post("/", handlers.ProcessOrder(g))
			})
			r.Get("/withdrawals", handlers.GetWithdrawals(g))
		})
		r.Group(func(r chi.Router) {
			r.Use(middlewares.UserValidation)

			r.Post("/login", handlers.Login(g))
			r.Post("/register", handlers.Register(g))
		})
	})
	return r
}
