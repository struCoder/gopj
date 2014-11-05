package controller

import (
	// "fmt"
	"github.com/hoisie/web"
	"html/template"
)

func GetLogin(ctx *web.Context) {
	t, err := template.ParseFiles("views/account/login.html")
	if err != nil {
		return
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func DoLogin(ctx *web.Context) {
	userName := ctx.Params["email"]
	userPwd := ctx.Params["pwd"]
}
