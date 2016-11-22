package controllers

import (
	m "blog/models"
	mAdmin "blog/models/admin"
	mLoginLog "blog/models/loginlog"
	"blog/utils"
	"blog/utils/strtool"
	"fmt"
)

var loginMsg string = ""

type LoginController struct {
	BaseController
}

/*

同时大家注意到新版本里面增加了URLMapping这个函数，
这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *LoginController) URLMapping() {
	this.Mapping("Admin", this.Admin)
	this.Mapping("Logon", this.Logon)
	this.Mapping("ExitEdit", this.ExitEdit)
}

// @router /blog/admin [get]
func (this *LoginController) Admin() {
	// this.Redirect("/blog/index", 302)
	this.Data["loginMsg"] = loginMsg
	this.TplName = "login.html"
	loginMsg = ""
}

// @router /blog/logon [post]
func (this *LoginController) Logon() {
	//请求的是登陆数据，那么执行登陆的逻辑判断
	this.Ctx.Request.ParseForm()
	loginName := this.Ctx.Request.Form.Get("loginName")
	password := this.Ctx.Request.Form.Get("password")
	admin := m.SysAdmin{LoginName: loginName, Password: strtool.Md5(password)}
	if mAdmin.Login(admin) {
		//登录成功设置session
		// sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		// sess.Set("uid", "1")
		// sess.Set("uname", "admin")

		this.Ctx.SetCookie("loginname", loginName, 0, "/")
		this.Ctx.SetCookie("password", strtool.Md5(password), 0, "/")
		// this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password="+strtool.Md5(password)+"; Max-Age=3600; Path=/; httponly")
		ip := this.Ctx.Input.IP()
		address := utils.GetAddressByIP(ip)
		loginlog := m.LoginLog{LoginName: loginName, Ip: ip, Address: address, Token: ""}
		mLoginLog.Insert(loginlog)

		this.Redirect("/blog/edit", 302)
	} else {
		this.Redirect("/blog/admin", 302)
		loginMsg = "登录名或密码不正确."
		fmt.Println("loginName:", loginName)
		fmt.Println("password:", password)
		fmt.Println("登录名或密码不正确.")
		fmt.Println("newPass", strtool.Md5("000000"))
	}

}

// @router /blog/exitEdit [get]
func (this *LoginController) ExitEdit() {
	this.Ctx.SetCookie("loginname", "", 0, "/")
	this.Ctx.SetCookie("password", "", 0, "/")
	// this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password=; Max-Age=0; Path=/; httponly")
	this.Ctx.Redirect(302, "/blog/index/1")
}
