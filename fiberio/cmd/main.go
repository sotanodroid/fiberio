package main

import fiber "github.com/gofiber/fiber"

func main() {
	app := fiber.New()
	app.Get("/", hello)
	app.Listen(8000)
  }

// Handler
func hello(c *fiber.Ctx){
	c.send("Hello, world!")
}
