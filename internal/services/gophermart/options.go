package gophermart

import (
	"context"

	"github.com/ansedo/gophermart/internal/storages/postgres"
)

type Option func(g *Gophermart)

func WithPostgreStorage(ctx context.Context) Option {
	return func(g *Gophermart) {
		g.Storage = postgres.New(ctx)
	}
}

func WithDefaultStorage(ctx context.Context) Option {
	return WithPostgreStorage(ctx)
}
