package controllers

import (
	m "blog/models"
	mArticle "blog/models/article"
	mCategory "blog/models/category"
	"strconv"
)

type MainController struct {
	BaseController
}

/*

同时大家注意到新版本里面增加了URLMapping这个函数，
这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *MainController) URLMapping() {
	this.Mapping("Index", this.Index)
	this.Mapping("Index1", this.Index1)
	this.Mapping("Home", this.Home)
	this.Mapping("Domain", this.Domain)
	this.Mapping("Guidang", this.Guidang)
	this.Mapping("ShowArticle", this.ShowArticle)
}

// @router / [get]
func (this *MainController) Domain() {
	this.Redirect("/blog/index/1", 302)
}

// @router /blog/home [get]
func (this *MainController) Home() {
	this.Redirect("/blog/index/1", 302)
}

// @router /blog/index [get]
func (this *MainController) Index1() {
	this.Redirect("/blog/index/1", 302)
}

// @router /blog/index/:pageNum [get]
func (this *MainController) Index() {
	this.Data["IsAdmin"] = this.IsAdmin
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	articles, maps := mArticle.QueryHomeData(20, pageNum, 10)
	this.Data["Articles"] = articles
	this.Data["AriticleMaps"] = maps
	this.Layout = "main.html"
	this.TplName = "home.html"
}

// @router /blog/guidang/:categoryId/:pageNum [get]
func (this *MainController) Guidang() {
	this.Data["IsAdmin"] = this.IsAdmin
	categorys := mCategory.GetCategoryAll()
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	categoryId, _ := strconv.ParseInt(this.Ctx.Input.Param(":categoryId"), 10, 64)
	articles, maps := mArticle.QueryArticleByCategory(categoryId, 20, pageNum, 10)
	this.Data["Articles"] = articles
	this.Data["AriticleMaps"] = maps
	this.Data["Categorys"] = categorys
	this.Data["currCategory"] = categoryId
	this.Layout = "main.html"
	this.TplName = "guidang.html"
}

// @router /blog/article/:type/:id [get]
func (this *MainController) ShowArticle() {
	this.Data["IsAdmin"] = this.IsAdmin

	var n int64
	var article = &m.Article{}
	var category = &m.Category{}
	aidStr := this.Ctx.Input.Param(":id")
	aid, _ := strconv.ParseInt(aidStr, 10, 64)
	typeStr := this.Ctx.Input.Param(":type")
	article, n = mArticle.GetArticleById(typeStr, aid)
	if n != aid {
		this.Redirect("/blog/article/n/"+strconv.FormatInt(n, 10), 302)
		return
	}

	if article != nil {
		this.Data["Article"] = article
		category = article.Category
		if category != nil {
			this.Data["Category"] = category
		}
	}

	this.Layout = "main.html"
	this.TplName = "article.html"
}
