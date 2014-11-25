package controller

import (
	"github.com/hoisie/web"
	"model"
)

//蛋疼的是，暂时没有想到如何写一个中间件检查登陆状态，所以这里
//很费事的写了很多的checklogin
func GetAddComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	render(ctx, "complaint/add", map[string]interface{}{
		"title": "业主投诉",
	})
}

func AddComplain(ctx *web.Context) {
	status := &loginStatus{}
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	inputName := ctx.Params["inputName"]
	inputTel := ctx.Params["inputTel"]
	beInputName := ctx.Params["beInputName"]
	beComplaintHome := ctx.Params["beComplaintHome"]
	inputReason := ctx.Params["inputReason"]
	dealPerson := ctx.Params["dealPerson"]

	isSuccess, errorStr := model.SaveComplaint(inputName, inputTel, beInputName, beComplaintHome, inputReason, dealPerson)
	if isSuccess {
		status.Code = 0
		status.Msg = "信息保存成功"
	} else {
		status.Code = 1
		status.Msg = "保存失败: " + errorStr
	}
	returnJson(ctx, status)
}

func GetComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	complaintMsgMap, count := model.GetMsgByLimit(0, 10)

	render(ctx, "complaint/del", map[string]interface{}{
		"complaint": complaintMsgMap,
		"title":     "撤消投诉",
		"count":     count,
		"start":     0,
		"end":       1,
	})
}

func PagingComplain(ctx *web.Context) {
	currentPage := ctx.Params["page"]
	if !isNum(currentPage) || currentPage < "1" {
		ctx.Abort(400, "非法访问")
		return
	}
	page := parseInt(currentPage)
	endPage := page * 10
	startPage := endPage - 10
	complaintMsgMap, count := model.GetMsgByLimit(startPage, endPage)
	render(ctx, "complaint/del", map[string]interface{}{
		"complaint": complaintMsgMap,
		"title":     "撤消投诉",
		"count":     count,
		"start":     page,
		"end":       page + 1,
	})
}

func DelComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	delId := ctx.Params["id"]
	if !isNum(delId) || delId == "0" {
		ctx.Abort(400, "非法访问")
		return
	}
	status := &dealStatus{}
	isSuccess := model.DoDel(parseInt(delId))
	if isSuccess {
		status.Code = 1
		status.Msg = "删除成功"
	} else {
		status.Code = 0
		status.Msg = "删除失败"
	}
	returnJson(ctx, status)
}

func GetTodoComplain(ctx *web.Context) {
	if !checkLogin(ctx) {
		ctx.Redirect(302, "/login")
		return
	}
	complaintMsgMap, count := model.GetMsgByLimit(0, 10)
	render(ctx, "complaint/deal", map[string]interface{}{
		"complaint": complaintMsgMap,
		"title":     "投诉处理",
		"count":     count,
		"start":     0,
		"end":       1,
	})
}
