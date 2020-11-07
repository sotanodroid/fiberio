package handler

import (
	fiber "github.com/gofiber/fiber/v2"
)

// GetByID ...
func GetByID(c *fiber.Ctx) error {
	c.SendString("Hello from getbyid")

	return nil
}

// Hello ...
func Hello(c *fiber.Ctx) error {
	c.SendString("Hello, world!")
	return nil
}
