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
	if err := repo.NewDB(cfg.DBURL); err != nil {
		return err
	}
	fib, err := newApp()
	if err != nil {
		return err
	}
	if err := fib.app.Listen(cfg.AppPort); err != nil {
		return err
	}
	return nil
}

// NewApp ...
func newApp() (*fiberio, error) {
	a := &fiberio{
		app: fiber.New(),
	}
	a.setup()

	return a, nil
}

func (f *fiberio) setup() {
	routerv1.SetupRoutes(f.app)
	middleware.SetUpLogger(f.app)
}
