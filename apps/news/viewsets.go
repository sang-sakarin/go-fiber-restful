package news

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/pkg"
	//"go-fiber-restful/pkg"
)

// List godoc
// @Summary
// @Description List news
// @Tags News
// @Router /v1/news/ [get]
// @Success 200 {string} Token "Test"
func List(c *fiber.Ctx) error {
	var news []News
	pkg.DB.Find(&news)

	return c.JSON(news)
}

func Create(c *fiber.Ctx) error {
	return c.Send([]byte("Create News"))
}

func Retrieve(c *fiber.Ctx) error {
	return c.Send([]byte("News 1"))
}

func Update(c *fiber.Ctx) error {
	return c.Send([]byte("Update News 1"))
}

func Delete(c *fiber.Ctx) error {
	return c.Send([]byte("Deleted News 1"))
}
