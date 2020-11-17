package router

import (
	"fiberio/app/handler"
	fiber "github.com/gofiber/fiber/v2"
)

func setupV1Routes(r fiber.Router, h *handler.BaseHandler) {
	r.Get("/", h.HelloV1)
	r.Get("/:id/", h.GetByIDV1)
}
