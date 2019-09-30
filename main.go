package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"go_mod/models"
	"go_mod/routers"
)

//初始化
func initApp() (app *iris.Application) {
	app = iris.New()
	//日志
	app.Use(logger.New())
	//错误处理 -- 404错误与500错误的统一处理
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON("Not Found Http Exception")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON("Server Error")
	})
	//连接数据库
	db, _ := gorm.Open("mysql","root:root@/go_test?charset=utf8&parseTime=True&loc=Local")
	//数据迁移
	db.AutoMigrate(&models.News{},&models.Users{})
	//跨域处理
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	//路由
	routers.LoadApiRoute(app,crs)

	return
}

func main() {
	app := initApp()

	app.Run(iris.Addr("localhost:8080"))
}
