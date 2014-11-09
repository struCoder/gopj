package controller

import (
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
