package pkg

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go-fiber-restful/apps/news"
	"go-fiber-restful/apps/users"
	_ "go-fiber-restful/docs"
)

func Routers(app *fiber.App) {

	app.Get("/docs/*", swagger.Handler)

	v1 := app.Group("v1")

	v1.Get("/users", users.GetAll)

	_news := v1.Group("news")
	_news.Get("", news.List)
	_news.Post("", news.Create)
	_news.Get(":id", news.Retrieve)
	_news.Patch(":id", news.Update)
	_news.Delete(":id", news.Delete)

}
