package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/kiwiframe/morpheus/backend/internal/adapters/postgres/sqlc"
	"github.com/kiwiframe/morpheus/backend/internal/core/ports"
)

type repo struct {
	queries sqlc.Querier
}

func NewPostgresRepository(pool *pgxpool.Pool) ports.Repository {
	return &repo{queries: sqlc.New(pool)}
}
