package handler

import (
	fiber "github.com/gofiber/fiber/v2"
)

// GetByIDV1 ...
func (h *BaseHandler) GetByIDV1(c *fiber.Ctx) error {

	res, err := h.db.GetAll()
	if err != nil {
		return err
	}
	c.SendString(res)

	return nil
}

// HelloV1 ...
func (h *BaseHandler) HelloV1(c *fiber.Ctx) error {
	c.SendString("Hello, world!")
	return nil
}
