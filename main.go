package main

import (
	"bookie/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", handlers.GetBooksList)
	app.Get("/:bookId", handlers.GetSingleBook)
	app.Post("/", handlers.CreateBook)
	app.Put("/:bookId", handlers.UpdateBook)
	app.Delete("/:bookId", handlers.DeleteBook)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Route not found",
			"data":    nil,
		})
	})

	app.Listen(":3000")
}
