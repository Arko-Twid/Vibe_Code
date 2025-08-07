package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	ConnectDB()
	DB.AutoMigrate(&Product{}, &User{})

	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	setupRoutes(app)

	app.Listen(":8080")
}
