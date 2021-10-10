package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/apps/users"
)

func Routers(app *fiber.App) {

	route := app.Group("v1")

	route.Get("/users", users.GetAll)
}
