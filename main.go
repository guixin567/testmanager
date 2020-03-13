package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"test1/config"
	"test1/controller"
	"test1/datasource"
	"test1/service"
)

func main() {
	app := iris.New()
	//1，设置Log级别
	app.Logger().SetLevel("debug")
	//2，注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manager/static", "./static")
	app.HandleDir("/img", "./static/img")
	//3，注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	//4,字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//5,错误配置,请求404,500
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(
			iris.Map{
				"message": "not found",
				"result":  "",
				"code":    iris.StatusNotFound,
			})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(
			iris.Map{
				"message": "server error",
				"result":  "",
				"code":    iris.StatusInternalServerError,
			})
	})
	//6,MVC和Session配置
	session := sessions.New(sessions.Config{})
	engine := datasource.New()
	//Admin登陆
	adminService := service.GetAdminService(engine)
	adminMVC := mvc.New(app.Party("/admin"))
	adminMVC.Register(adminService, session.Start)
	adminMVC.Handle(new(controller.AdminController))

	//用户
	userMVC := mvc.New(app.Party("/user"))
	userMVC.Register(session.Start)
	userMVC.Handle(new(controller.UserController))
	//7,服务器，端口信息配置
	initConfig := config.InitConfig()
	_ = app.Run(iris.Addr(initConfig.Port))
}
