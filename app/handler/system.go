package handler

import (
	fiber "github.com/gofiber/fiber/v2"
)

// Ping ...
func (h *BaseHandler) Ping(c *fiber.Ctx) error {

	_, err := h.db.Ping()
	if err != nil {
		return err
	}
	c.SendStatus(200)

	return nil
}