package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"go_mod/controllers"
	"go_mod/middleware"
)

//api路由加载
func LoadApiRoute(app *iris.Application,crs context.Handler) {
	route := app.Party("api/", crs).AllowMethods(iris.MethodOptions)
	{
		route.Get("/news", controllers.News)
		//新闻
		route.PartyFunc("/news", func(news router.Party) {
			news.Use(middleware.JwtHandler().Serve)
			news.Get("/detail/{id:uint}", controllers.NewsDetail)
			news.Get("/list", controllers.NewsList)
			news.Post("/create", controllers.CreateNews)
		})
		//用户
		route.PartyFunc("/users", func(users router.Party) {
			users.Post("/login", controllers.Login)
			users.Post("/register", controllers.UserRegister)
		})
	}

	return
}
