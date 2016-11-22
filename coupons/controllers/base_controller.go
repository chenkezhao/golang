package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type T interface{}
type Checker interface {
	CheckLogin()
}
type InitUploadPath interface {
	InitUploadPath()
}

type BaseController struct {
	beego.Controller
	IsAdmin        bool
	LoginMsg       string
	UploadBasePath string
}

var (
	Cfg   = beego.AppConfig
	Email string
)

func init() {
	Email = Cfg.String("blog_email")
}

func (this *BaseController) Prepare() {
	this.Data["Email"] = Email
	this.UploadBasePath = Cfg.String("upload_base_path")
	if app, ok := this.AppController.(InitUploadPath); ok {
		app.InitUploadPath()
	}
}

func (this *BaseController) ResponseData(success bool, message string) map[string]interface{} {
	response := make(map[string]interface{})
	response["success"] = success
	response["message"] = message
	return response
}

/*
指定页面，并且返回公共参数
*/
func (this *BaseController) Show(url string) {
	this.Data["staticUrl"] = beego.AppConfig.String("staticUrl")
	this.TplName = url
}
func (this *BaseController) ClientShow(url string) {
	this.Data["staticUrl"] = beego.AppConfig.String("staticUrl")
	this.Layout = "public/index.html"
	this.TplName = url
}

/*

同时大家注意到新版本里面增加了URLMapping这个函数，
这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *BaseController) URLMapping() {
	this.Mapping("ShowHtml", this.ShowHtml)
}

//加载页面
// @router /showHtml/:resPath [get]
func (this *BaseController) ShowHtml() {
	resPath := this.Ctx.Input.Param(":resPath")
	fmt.Println("............>>>>>>>>>>>>>>>>>", resPath)
	resPath = strings.Replace(resPath, "+++", "/", -1)
	this.Show(resPath)
}
