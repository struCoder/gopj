package controller

const (
	viewsDir = "views/"
	htmlExt  = ".html"
	secret   = "strucoder"
)

//用来在渲染模板时指定css路径
type cssPath struct {
	Name string
}
