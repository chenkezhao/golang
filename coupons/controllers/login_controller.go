package controllers

import (
	"coupons/utils/strtool"
	// "fmt"
)

type LoginController struct {
	BaseController
}

/*

同时大家注意到新版本里面增加了URLMapping这个函数，
这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *LoginController) URLMapping() {
	this.Mapping("Login", this.Login)
	this.Mapping("SignIn", this.SignIn)
}

// @router /login [get]
func (this *LoginController) Login() {
	this.Data["LoginMsg"] = this.LoginMsg
	this.TplName = "login.html"
}

// @router /signIn [post]
func (this *LoginController) SignIn() {
	//请求的是登陆数据，那么执行登陆的逻辑判断
	this.Ctx.Request.ParseForm()
	loginName := this.Ctx.Request.Form.Get("loginName")
	password := this.Ctx.Request.Form.Get("password")
	this.Ctx.SetCookie("loginName", loginName, 0, "/")
	this.Ctx.SetCookie("password", strtool.Md5(password), 0, "/")
	this.Redirect("/admin/main", 302)
}
