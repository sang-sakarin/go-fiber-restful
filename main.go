package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/pkg"
)

func main() {
	app := fiber.New()

	extendApp := pkg.ExtendApp{}

	extendApp.ApplySetting(app)

	pkg.Routers(app)

	app.Listen(":3000")
}
