package controllers

import (
	"github.com/astaxie/beego"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin bool
}

var (
	Cfg          = beego.AppConfig
	RootName     string
	RootPass     string
	BlogEmail    string
	BlogTitle    string
	BlogIconPath string
)

func init() {
	RootName = Cfg.String("root_name")
	RootPass = Cfg.String("root_pass")
	BlogEmail = Cfg.String("blog_email")
	BlogTitle = Cfg.String("blog_title")
	BlogIconPath = Cfg.String("blog_icon_path")
}

func (this *BaseController) Prepare() {
	this.Data["BlogEmail"] = BlogEmail
	this.Data["BlogTitle"] = BlogTitle
	this.Data["BlogIconPath"] = BlogIconPath
	this.AssignIsAdmin()
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	name := this.Ctx.GetCookie("loginname")
	password := this.Ctx.GetCookie("password")
	if name == "" || password == "" {
		this.IsAdmin = false
		return
	}

	if name != RootName || password != RootPass {
		this.IsAdmin = false
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin
}
