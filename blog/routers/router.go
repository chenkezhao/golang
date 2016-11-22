package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//RESTful Controller 路由
	// beego.Router("/", &controllers.MainController{}, "*:Index")
	// beego.Router("/home", &controllers.MainController{}, "*:Home")

	//使用注解路由
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.BlogController{})
}
