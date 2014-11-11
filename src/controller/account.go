package controller

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hoisie/web"
	"model"
)

//存储用户登陆成功与否
type loginStatus struct {
	Code int32
	Msg  string
}

var store = sessions.NewCookieStore([]byte(secret))

func GetLogin(ctx *web.Context) {
	render(ctx, "account/login", nil)
}

/*
 *用户登入, 设置保存cookie, 跳转到主页面
**/
func DoLogin(ctx *web.Context) {
	session, _ := store.Get(ctx.Request, "user")
	fmt.Println(session.Values["userId"])
	userName := ctx.Params["email"]
	userPwd := ctx.Params["pwd"]
	md5Pwd := encryptPwd(userPwd)
	isUser, userStruct := model.FindUser(userName, md5Pwd)
	if isUser {

		session.Values["userId"] = userStruct["user"].Id
		session.Values["userName"] = userStruct["user"].Name
		session.Save(ctx.Request, ctx.ResponseWriter)
		ctx.Redirect(200, "/")
		return
	}
	status := loginStatus{0, "fail"}
	returnJson(ctx, status)
}

/*
 *用户登出, 清除cookie, 跳转到登入页面
**/
func LogOut(ctx *web.Context) {
	session, _ := store.Get(ctx.Request, "user")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(ctx.Request, ctx.ResponseWriter)
	render(ctx, "account/login", nil)
}
