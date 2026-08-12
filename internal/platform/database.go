package platform

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool interface is for a pool database connection, meaning it allows for
// connection reuse.
type Pool interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

// DB is the struct containing the database connection pointer.
type DB struct {
	Pool Pool
}

// NewDB creates a new DB struct containing the pgxpool connection
// to the postgres database.
func NewDB(ctx context.Context, env string) (*DB, error) {
	slog.Info("starting database...")

	e := os.Getenv(env)
	if e == "" {
		return nil, fmt.Errorf("can't get database environment variable: %s, not set", e)
	}

	pool, err := pgxpool.New(ctx, e)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &DB{
		Pool: pool,
	}, nil
}
