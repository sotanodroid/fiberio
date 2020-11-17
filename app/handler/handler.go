package handler

import (
	repo "fiberio/app/repo"
	fiber "github.com/gofiber/fiber/v2"
)

// BaseHandler ...
type BaseHandler struct {
	db *repo.Db
}

// NewBasehandler ...
func NewBasehandler(db *repo.Db) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// GetByID ...
func (h *BaseHandler) GetByID(c *fiber.Ctx) error {

	res, err := h.db.GetAll()
	if err != nil {
		return err
	}
	c.SendString(res)

	return nil
}

// Hello ...
func (h *BaseHandler) Hello(c *fiber.Ctx) error {
	c.SendString("Hello, world!")
	return nil
}
