package router

import (
	"fiberio/app/handler"
	fiber "github.com/gofiber/fiber/v2"
)

func setupSystemRoutes(r fiber.Router, h *handler.BaseHandler) {
	r.Get("/ping/", h.Ping)
}
