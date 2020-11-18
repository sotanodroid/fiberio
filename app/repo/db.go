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
func NewDB(dbURL string) (*Db, error) {
	ctx := context.Background()
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

// Ping ...
func (db *Db) Ping() (int, error) {
	ctx := context.Background()
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var result int
	if err := conn.QueryRow(ctx, "SELECT 1").Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}
