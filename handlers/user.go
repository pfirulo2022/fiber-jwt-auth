package handlers

import "github.com/gofiber/fiber/v2"

func User(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
