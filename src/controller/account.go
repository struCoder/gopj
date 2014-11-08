package controller

import (
	// "fmt"
	"crypto/md5"
	"github.com/hoisie/web"
	"html/template"
	"model"
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
	md5Pwd = md5.Sum([]byte(userPwd))
	isUser := model.FindUser(userName, md5Pwd)
	if isUser {
		ctx.Redirect(302, "http://baidu.com")
		return
	}
	ctx.WriteString("Sdfsfsd")

}
