package controller

import (
	"github.com/gorilla/sessions"
	"github.com/hoisie/web"
)

func GetComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
	}
}
