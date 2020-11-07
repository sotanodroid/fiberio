package main

import (
	middleware "fiberio/middleware"
	routerv1 "fiberio/router/v1"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.SetUpLogger(app)
	routerv1.SetupRoutes(app)

	app.Listen(":8000")
}
