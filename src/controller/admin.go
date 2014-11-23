package controller

import (
	"github.com/hoisie/web"
)

func GetAdmin(ctx *web.Context) {
	if checkLogin(ctx) {
		ctx.Redirect(302, "/")
		return
	}

	if !isAdmin(ctx) {
		ctx.Abort(403, "YOU HAVE NO RIGHT")
	}
}
