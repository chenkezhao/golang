package administrator

import (
	"coupons/controllers"
	models "coupons/models/admin"
	"coupons/utils/strtool"
	"fmt"
	"github.com/astaxie/beego"
)

//登录 才能访问些控制器
type AdminController struct {
	controllers.BaseController
	SysUserInfo models.SysUser
}

var (
	Cfg   = beego.AppConfig
	Email string
)

func init() {
	Email = Cfg.String("blog_email")
}
func (this *AdminController) Prepare() {
	this.Data["Email"] = Email

	this.AssignIsAdmin()
	if app, ok := this.AppController.(controllers.Checker); ok {
		app.CheckLogin()
	}
}

func (this *AdminController) AssignIsAdmin() {
	name := this.Ctx.GetCookie("loginName")
	password := this.Ctx.GetCookie("password")
	this.IsAdmin = false
	//查询用户信息
	this.SysUserInfo = SignIn(name, strtool.Md5(password))
	fmt.Println("-----------每个管理员请求，都要验证下遍------user>>>>>>>>>>>>>>>>>>", this.SysUserInfo)
	if this.SysUserInfo.Id != 0 {
		this.IsAdmin = true
	}
	this.Data["IsAdmin"] = this.IsAdmin
	this.Data["AdminInfo"] = this.SysUserInfo
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin {
		this.LoginMsg = "请输入你的用户名和密码，进行登录."
		this.Redirect("/login", 302)
	}
}

/*

	同时大家注意到新版本里面增加了URLMapping这个函数，
	这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
	如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *AdminController) URLMapping() {
	this.Mapping("ToMain", this.ToMain)
	this.Mapping("Exit", this.Exit)
}

// @router /admin/main [get]
func (this *AdminController) ToMain() {
	fmt.Println("-------------->>>>>>>>>>>>>", "登录管理员页面")
	this.TplName = "admin/base/main.html"
}

// @router /admin/exit [get]
func (this *AdminController) Exit() {
	this.Ctx.SetCookie("loginname", "", 0, "/")
	this.Ctx.SetCookie("password", "", 0, "/")
	this.Redirect("/login", 302)
}
