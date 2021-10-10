package users

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    "first data",
	})
}
