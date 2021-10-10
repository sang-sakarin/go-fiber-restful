package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/pkg/routers"
)

func main() {
	app := fiber.New()

	routers.Routers(app)

	app.Listen(":3000")
}
