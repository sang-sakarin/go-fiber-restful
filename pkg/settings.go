package pkg

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func (extendApp *ExtendApp) ApplySetting(app *fiber.App) {
	cfg := Config{}
	env.Parse(&cfg)
	var err error

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
}
