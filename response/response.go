package response

import "github.com/gofiber/fiber/v2"

func Succcess(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(400).JSON(fiber.Map{
		"success": false,
		"message": message,
		"data":    data,
	})
}
