package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"html/template"
	"log"
	"model"
	"regexp"
)

//服务器错误友好提示
func serverWrong(ctx *web.Context) {
	ctx.Abort(500, "sory! server wrong")
}

//模板渲染
func render(ctx *web.Context, htmlName string, data map[string]interface{}) {
	t := template.Must(template.ParseFiles(viewsDir+htmlName+htmlExt, "views/header.tmpl", "views/footer.tmpl"))
	err := t.Execute(ctx.ResponseWriter, data)
	if err != nil {
		log.Fatal(err)
		serverWrong(ctx)
	}
}

//返回md5加密后的字符串
func encryptPwd(pwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}

//返回json字符串
func returnJson(ctx *web.Context, input interface{}) {
	b, _ := json.Marshal(input)
	ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
	ctx.WriteString(string(b))
}

//判断是否用户已经登录
func checkLogin(ctx *web.Context) bool {
	session, _ := store.Get(ctx.Request, "user")
	if session.Values["userId"] != nil {
		return true
	}
	return false
}

//判读是否为Email
func isEmail(email string) bool {
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")
	return regEmail.MatchString(email)
}

// 判断字符串是否为空
func isEmpty(str string) bool {
	if len(str) < 1 {
		return true
	}
	return false
}

// 判断是否为管理员
func isAdmin(ctx *web.Context) bool {
	return false
}

//得到当前用户
func getCurrentUser(ctx *web.Context) model.UserInfo {
	session, _ := store.Get(ctx.Request, "user")
	// currentUser := make(map[string]string)
	var user model.UserInfo
	user.Id = (session.Values["userId"]).(string)
	user.Name = (session.Values["userName"]).(string)
	//其实在这更安全的方法就是得到用户Id后进行一次数据库查询,
	//目前就先这样吧
	return user
}
