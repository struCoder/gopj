package controller

import (
	// "github.com/gorilla/sessions"
	"github.com/hoisie/web"
)

type test struct {
	Name string
}

func GetComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	te := test{"David"}
	render(ctx, "complaint/index", te)
}
