package routerv1

import (
	"fiberio/internal/app/handler"
	fiber "github.com/gofiber/fiber/v2"
)

// SetupRoutes ...
func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1")

	api.Get("/", handler.Hello)
	api.Get("/:id/", handler.GetByID)
}
