package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Db ...
type Db struct {
	Pool *pgxpool.Pool
}

// NewDB ...
func NewDB(ctx context.Context, dbURL string) (*Db, error) {
	conn, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	return &Db{
		Pool: conn,
	}, nil
}

// GetAll ...
func (db *Db) GetAll() (string, error) {
	ctx := context.Background()
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()

	var result string
	if err := conn.QueryRow(ctx, "SELECT 'Hello from db';").Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
