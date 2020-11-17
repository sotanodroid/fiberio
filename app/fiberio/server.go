package server

import (
	middleware "fiberio/app/middleware"
	repo "fiberio/app/repo"
	routerv1 "fiberio/app/router/v1"
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
func newApp() (*fiberio) {
	return &fiberio{
		app: fiber.New(),
	}
}

func (f *fiberio) setup(DBURL string) error {
	db, err := repo.NewDB(DBURL)
	if err != nil {
		return err
	}

	routerv1.SetupRoutes(f.app, db)
	middleware.SetUpLogger(f.app)

	return nil
}
