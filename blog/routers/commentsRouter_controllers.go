package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/blog/edit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "AddArticle",
			Router: `/blog/addArticle`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "DelArticle",
			Router: `/blog/delArticle`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "EditArticle",
			Router: `/blog/editArticle/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "DelCategoryById",
			Router: `/blog/delCategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "UpdateCategory",
			Router: `/blog/updateCategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "QueryChildCount",
			Router: `/blog/queryChildCount`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "LoginLogs",
			Router: `/blog/loginLogs`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:BlogController"] = append(beego.GlobalControllerRouter["blog/controllers:BlogController"],
		beego.ControllerComments{
			Method: "ClearLoginLog",
			Router: `/blog/clearLoginLog`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Admin",
			Router: `/blog/admin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Logon",
			Router: `/blog/logon`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog/controllers:LoginController"],
		beego.ControllerComments{
			Method: "ExitEdit",
			Router: `/blog/exitEdit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Domain",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/blog/home`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index1",
			Router: `/blog/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/blog/index/:pageNum`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "Guidang",
			Router: `/blog/guidang/:categoryId/:pageNum`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:MainController"] = append(beego.GlobalControllerRouter["blog/controllers:MainController"],
		beego.ControllerComments{
			Method: "ShowArticle",
			Router: `/blog/article/:type/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
