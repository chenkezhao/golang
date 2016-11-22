package controllers

//登录 才能访问些控制器
type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin {
		this.Redirect("/blog/admin", 302)
	}
}
