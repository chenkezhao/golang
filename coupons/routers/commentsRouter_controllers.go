package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["coupons/controllers:BaseController"] = append(beego.GlobalControllerRouter["coupons/controllers:BaseController"],
		beego.ControllerComments{
			Method: "ShowHtml",
			Router: `/showHtml/:resPath`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers:ClientController"] = append(beego.GlobalControllerRouter["coupons/controllers:ClientController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers:ClientController"] = append(beego.GlobalControllerRouter["coupons/controllers:ClientController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers:ClientController"] = append(beego.GlobalControllerRouter["coupons/controllers:ClientController"],
		beego.ControllerComments{
			Method: "Tags",
			Router: `/index/:tagid/:stagid/:categoryid/:page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers:LoginController"] = append(beego.GlobalControllerRouter["coupons/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers:LoginController"] = append(beego.GlobalControllerRouter["coupons/controllers:LoginController"],
		beego.ControllerComments{
			Method: "SignIn",
			Router: `/signIn`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
