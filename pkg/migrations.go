package pkg

import (
	"go-fiber-restful/apps/news"
	"go-fiber-restful/apps/users"
	"go-fiber-restful/pkg/databases"
)

func (extendApp *ExtendApp) UseMigration() {
	databases.DBConn.Migrator().AutoMigrate(&users.User{})
	databases.DBConn.Migrator().AutoMigrate(&news.News{})
}
