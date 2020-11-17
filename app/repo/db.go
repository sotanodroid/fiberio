package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Db ...
type Db struct {
	Pool *pgxpool.Pool
	ctx context.Context
}

// NewDB ...
func NewDB(ctx context.Context, dbURL string) (*Db, error) {
	conn, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	return &Db{
		Pool: conn,
		ctx: ctx,
	}, nil
}

// GetAll ...
func (db *Db) GetAll() (string, error) {
	conn, err := db.Pool.Acquire(db.ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()

	var result string
	if err := conn.QueryRow(db.ctx, "SELECT 'Hello from db';").Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
