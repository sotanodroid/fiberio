package handler

import (
	repo "fiberio/app/repo"
	fiber "github.com/gofiber/fiber/v2"
)

// GetByID ...
func GetByID(c *fiber.Ctx) error {

	res, err := repo.Db.GetAll()
	if err != nil {
		return err
	}
	c.SendString(res)

	return nil
}

// Hello ...
func Hello(c *fiber.Ctx) error {
	c.SendString("Hello, world!")
	return nil
}
