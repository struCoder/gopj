package main

import (
	"controller"
	"github.com/hoisie/web"
	"net/http"
)

func initRouter() {
	//静态资源
	web.Get("/static", http.FileServer(http.Dir("/Users/David/gopj")))

	//登入登出
	web.Get("/login", controller.GetLogin)
	web.Post("/login", controller.DoLogin)
	web.Get("/logout", controller.LogOut)
	//主页
	web.Get("/", controller.GetHome)

	//投诉管理
	web.Get("/complain", controller.GetComplain)
	web.Post("/complain", controller.AddComplain)
	web.Post("/complain", controller.DelComplain)
}
func main() {
	initRouter()
	web.Run(":6070")
}
