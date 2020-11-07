package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetUpLogger ...
func SetUpLogger(app *fiber.App) {
	// Default middleware config
	app.Use(logger.New())
}
