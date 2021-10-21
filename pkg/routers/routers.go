package routers

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/apps/users"
	_ "go-fiber-restful/docs"
)

func Routers(app *fiber.App) {

	app.Get("/swagger/*", swagger.Handler)

	v1 := app.Group("v1")

	v1.Get("/users", users.GetAll)
}
