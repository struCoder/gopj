package controller

import (
	// "github.com/gorilla/sessions"
	"github.com/hoisie/web"
)

func GetComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	currentUser := getCurrentUser(ctx)
	render(ctx, "complaint/add", map[string]interface{}{
		"user": currentUser,
	})
}

func AddComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	inputName := ctx.Params["inputName"]
	inputTel := ctx.Params["inputTel"]
	beInputName := ctx.Params["beInputName"]
	beComplaintHome := ctx.Params["beComplaintHome"]
	inputReason := ctx.Params["inputReason"]
	dealPerson := ctx.Params["dealPerson"]
}
