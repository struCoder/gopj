package controller

import (
	// "fmt"
	// "crypto/md5"
	"github.com/hoisie/web"
	// "model"
)

func GetLogin(ctx *web.Context) {
	render(ctx, "account/login", nil)
}

// func DoLogin(ctx *web.Context) {
// 	userName := ctx.Params["email"]
// 	userPwd := ctx.Params["pwd"]
// 	md5Pwd := md5.Sum([]byte(userPwd))
// 	isUser := model.FindUser(userName, md5Pwd)
// 	if isUser {
// 		ctx.Redirect(302, "http://baidu.com")
// 		return
// 	}
// 	ctx.WriteString("Sdfsfsd")

// }
