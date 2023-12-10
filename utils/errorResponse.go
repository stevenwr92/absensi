package utils

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message":    message,
		"statusCode": statusCode,
	})
}
