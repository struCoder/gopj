package main

import (
	"controller"
	"github.com/hoisie/web"
	"net/http"
)

func initRouter() {
	web.Get("/static", http.FileServer(http.Dir("/Users/David/gopj")))
	web.Get("/login", controller.GetLogin)
	web.Post("/login", controller.DoLogin)
	web.Get("/index", controller.GetHome)
}
func main() {
	initRouter()
	web.Run(":6070")
}
