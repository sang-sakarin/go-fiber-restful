package users

import "github.com/gofiber/fiber/v2"

// GetAll godoc
// @Summary
// @Description Get an item by its ID
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Router /v1/users/ [get]
func GetAll(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    "first data",
	})
}
