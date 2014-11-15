package controller

import (
	// "github.com/gorilla/sessions"
	"github.com/hoisie/web"
)

type te struct {
	Name string
	Id   string
}

func GetComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	u := te{"david", "22"}
	render(ctx, "complaint/index", map[string]interface{}{
		"user": u,
	})
}
