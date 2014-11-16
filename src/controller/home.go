package controller

import (
	// "fmt"
	"github.com/hoisie/web"
)

func GetHome(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
	}
	currentUser := getCurrentUser(ctx)
	render(ctx, "index", map[string]interface{}{
		"user": currentUser,
	})
}
