package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func UseDatabase(cfg Config) {

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	//DB.Migrator().AutoMigrate(&users.User{})
	//DB.Migrator().AutoMigrate(&news.News{})
}
