package helpers

import "github.com/gofiber/fiber/v2"

func ContentType(c *fiber.Ctx) string {
	contentType := c.Get(fiber.HeaderContentType)
	return contentType
}
