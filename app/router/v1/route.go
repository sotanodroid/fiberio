package routerv1

import (
	"fiberio/app/handler"
	fiber "github.com/gofiber/fiber/v2"
	repo "fiberio/app/repo"
)

// SetupRoutes ...
func SetupRoutes(app *fiber.App, db *repo.Db) {
	api := app.Group("/v1")

	baseHandler := handler.NewBasehandler(db)

	api.Get("/", baseHandler.Hello)
	api.Get("/:id/", baseHandler.GetByID)
}
