package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/hoisie/web"
	"html/template"
	"log"
)

func serverWrong(ctx *web.Context) {
	ctx.Abort(500, "sory! server wrong")
}
func render(ctx *web.Context, htmlName string, data interface{}) {
	t, err := template.ParseFiles(viewsDir + htmlName + htmlExt)
	if err != nil {
		log.Fatal(err)
		serverWrong(ctx)
	}
	t.Execute(ctx.ResponseWriter, data)
}

func encryptPwd(pwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}
