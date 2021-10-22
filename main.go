package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/pkg"
	"go-fiber-restful/pkg/routers"
)

func main() {
	app := fiber.New()

	extendApp := pkg.ExtendApp{}

	extendApp.ApplySetting(app)

	routers.Routers(app)

	app.Listen(":3000")
}
