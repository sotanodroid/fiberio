package handler

import (
	"context"
	repo "fiberio/internal/app/repo"
	fiber "github.com/gofiber/fiber/v2"
)

// GetByID ...
func GetByID(c *fiber.Ctx) error {
	ctx := context.Background()
	conn, err := repo.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	conn.Query(ctx, "SELECT 1;")
	defer conn.Release()

	c.SendString("Selected 1")

	return nil
}

// Hello ...
func Hello(c *fiber.Ctx) error {
	c.SendString("Hello, world!")
	return nil
}
