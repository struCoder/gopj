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
	web.Get("/complaint/add", controller.GetComplain)
	web.Post("/complaint/add", controller.AddComplain)
	web.Get("/complaint/del", controller.DelComplain)
	// web.Get("/complain/status", controller.StatusComplain)
	// web.Post("/complain/deal", controller.DealComplain)

	//系统管理
	// web.Get("/admin", controller.GetAdmin)
	// web.Post("/admin/add", controller.AddAdmin)
	// web.Get("/admin/profile", controller.GetAdminProfile)
	// web.Post("/admin/del", controller.DelAdmin)
}
func main() {
	initRouter()
	web.Run(":6070")
}
