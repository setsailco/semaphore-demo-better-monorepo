package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Initial release
func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// start the application on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	app.Use(logger.New())

	// Register the index route with a simple
	// JSON response. It should return status
	// code 200
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Hello from go-service-2!",
		})
	})

	// Return the configured app
	return app
}
