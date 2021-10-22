package pkg

import (
	"fmt"
	"go-fiber-restful/pkg/databases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func (extendApp *ExtendApp) UseDatabase(cfg Config) {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	var err error

	databases.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
}
