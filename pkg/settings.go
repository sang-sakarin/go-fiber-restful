package pkg

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2"
)

func (extendApp *ExtendApp) ApplySetting(app *fiber.App) {
	cfg := Config{}
	env.Parse(&cfg)

	// Use database
	extendApp.UseDatabase(cfg)

	// Use migration
	extendApp.UseMigration()
}
