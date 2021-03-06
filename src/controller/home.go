package controller

import (
	// "fmt"
	"github.com/hoisie/web"
)

func GetHome(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	render(ctx, "index", map[string]interface{}{
		"title": "主页",
	})
}
