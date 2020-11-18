package router

import (
	"fiberio/app/handler"
	repo "fiberio/app/repo"
	fiber "github.com/gofiber/fiber/v2"
)

// SetupRoutes ...
func SetupRoutes(app *fiber.App, db *repo.Db) {
	v1 := app.Group("/api/v1")
	system := app.Group("/")

	baseHandler := handler.NewBasehandler(db)
	setupV1Routes(v1, baseHandler)
	setupSystemRoutes(system, baseHandler)
}
