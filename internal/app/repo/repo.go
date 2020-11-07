package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB ...
var DB *pgxpool.Pool

// NewDB ...
func NewDB(dbURL string) error {
	var err error
	DB, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return err
	}

	return nil
}
