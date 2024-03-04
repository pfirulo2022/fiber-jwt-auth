package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/fiber-jwt-auth/models"
)

func User(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
