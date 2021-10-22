package pkg

import (
	"fmt"
	"go-fiber-restful/apps/news"
	"go-fiber-restful/apps/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func (extendApp *ExtendApp) UseDatabase(cfg Config) {

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
}

func (extendApp *ExtendApp) UseMigration() {
	DB.Migrator().AutoMigrate(&users.User{})
	DB.Migrator().AutoMigrate(&news.News{})
}
