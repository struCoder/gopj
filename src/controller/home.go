package controller

import (
	"github.com/hoisie/web"
)

func GetHome(ctx *web.Context) {
	render(ctx, "index", nil)
}
