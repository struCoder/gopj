package controller

import (
	// "encoding/json"
	// "fmt"
	// "github.com/gorilla/sessions"
	"github.com/hoisie/web"
	"model"
)

func GetLogin(ctx *web.Context) {
	render(ctx, "account/login", nil)
}

func DoLogin(ctx *web.Context) {
	userName := ctx.Params["email"]
	userPwd := ctx.Params["pwd"]
	md5Pwd := encryptPwd(userPwd)
	isUser, userInfoArr := model.FindUser(userName, md5Pwd)
	if isUser {
		ctx.WriteString(userInfoArr[0])
		return
	}

	ctx.WriteString("sdfsdf")

}
