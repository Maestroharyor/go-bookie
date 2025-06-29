package response

import "github.com/gofiber/fiber/v2"

func Succcess(c *fiber.Ctx, message string, data interface{}, statusCode ...int) error {
	code := 200

	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, message string, data interface{}, statusCode ...int) error {
	code := 400

	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"data":    data,
	})
}
