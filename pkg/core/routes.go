package core

import "github.com/gofiber/fiber/v2"

func HandleHello(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "your first json response",
	})
}
