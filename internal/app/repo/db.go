package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	Pool *pgxpool.Pool
}

// Db ...
var Db *db

// NewDB ...
func NewDB(dbURL string) error {
	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return err
	}

	Db = &db{
		Pool: conn,
	}

	return nil
}

func (r *db) GetAll() (string, error) {
	ctx := context.Background()
	conn, err := r.Pool.Acquire(ctx)
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
