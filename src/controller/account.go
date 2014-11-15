package controller

import (
	// "fmt"
	"github.com/gorilla/sessions"
	"github.com/hoisie/web"
	"model"
)

var store = sessions.NewCookieStore([]byte(secret))

func GetLogin(ctx *web.Context) {
	if checkLogin(ctx) {
		ctx.Redirect(302, "/")
		return
	}
	render(ctx, "account/login", nil)
}

/*
 *用户登入, 设置保存cookie, 跳转到主页面
**/
func DoLogin(ctx *web.Context) {
	var status loginStatus
	if checkLogin(ctx) {
		ctx.Redirect(302, "/")
		return
	}
	userName := ctx.Params["email"]
	userPwd := ctx.Params["pwd"]
	if !isEmail(userName) || isEmpty(userPwd) {
		status = loginStatus{0, "请确保登录信息完整"}
		returnJson(ctx, status)
		return
	}
	md5Pwd := encryptPwd(userPwd)
	isUser, userStruct := model.FindUser(userName, md5Pwd)
	if isUser {
		session, _ := store.Get(ctx.Request, "user")
		session.Values["userId"] = userStruct["user"].Id
		session.Values["userName"] = userStruct["user"].Name
		session.Save(ctx.Request, ctx.ResponseWriter)
		status = loginStatus{1, "/"}
		returnJson(ctx, status)
		return
	}
	status = loginStatus{0, "确保您的账号密码无误"}
	returnJson(ctx, status)
}

/*
 *用户登出, 清除cookie, 跳转到登入页面
**/
func LogOut(ctx *web.Context) {
	session, _ := store.Get(ctx.Request, "user")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(ctx.Request, ctx.ResponseWriter)
	ctx.Redirect(302, "/login")
}
