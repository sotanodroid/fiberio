package server

import (
	"context"
	middleware "fiberio/app/middleware"
	repo "fiberio/app/repo"
	router "fiberio/app/router"
	fiber "github.com/gofiber/fiber/v2"
)

type fiberio struct {
	app *fiber.App
}

// Start ...
func Start(cfg *Config) error {
	fib := newApp()
	if err := fib.setup(cfg.DBURL); err != nil {
		return err
	}
	if err := fib.app.Listen(cfg.AppPort); err != nil {
		return err
	}
	return nil
}

// NewApp ...
func newApp() *fiberio {
	return &fiberio{
		app: fiber.New(),
	}
}

func (f *fiberio) setup(DBURL string) error {
	ctx := context.Background()
	db, err := repo.NewDB(ctx, DBURL)
	if err != nil {
		return err
	}

	router.SetupRoutes(f.app, db)
	middleware.SetUpLogger(f.app)

	return nil
}
