package controllers

import (
	m "blog/models"
	mArticle "blog/models/article"
	mCategory "blog/models/category"
	mLoginLog "blog/models/loginlog"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type BlogController struct {
	AdminController
}

/*

同时大家注意到新版本里面增加了URLMapping这个函数，
这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (this *BlogController) URLMapping() {
	this.Mapping("Edit", this.Edit)
	this.Mapping("AddArticle", this.AddArticle)
	this.Mapping("DelArticle", this.DelArticle)
	this.Mapping("EditArticle", this.EditArticle)
	this.Mapping("DelCategoryById", this.DelCategoryById)
	this.Mapping("UpdateCategory", this.UpdateCategory)
	this.Mapping("QueryChildCount", this.QueryChildCount)
	this.Mapping("LoginLogs", this.LoginLogs)
	this.Mapping("ClearLoginLog", this.ClearLoginLog)
}

// @router /blog/edit [get]
func (this *BlogController) Edit() {
	if !this.IsAdmin {
		this.Redirect("/blog/index/1", 302)
	} else {
		categorys := mCategory.GetCategoryAll()
		this.Data["Categorys"] = categorys
		this.Data["IsAdmin"] = this.IsAdmin
		this.Layout = "main.html"
		this.TplName = "edit.html"
	}
}

// @router /blog/addArticle [post]
func (this *BlogController) AddArticle() {
	this.Ctx.Request.ParseForm()
	article := m.Article{}
	article.Id, _ = strconv.ParseInt(this.Ctx.Request.Form.Get("id"), 10, 64)
	article.Title = this.Ctx.Request.Form.Get("title")
	article.Summary = this.Ctx.Request.Form.Get("summary")
	article.Content = this.Ctx.Request.Form.Get("content")
	category := &m.Category{}

	var id int64
	id, _ = strconv.ParseInt(this.Ctx.Request.Form.Get("categoryId"), 10, 64)
	var cname = strings.TrimSpace(this.Ctx.Request.Form.Get("categoryAdded"))

	if id == 0 && cname != "" {
		// 1.id为空,parentid为空，表示初始化分类
		category.Name = cname

	}
	if id != 0 && cname != "" {
		tempC := mCategory.GetCategoryById(id)
		if tempC.ParentId == 0 {
			// 2.id为空，parentid不为空，表示新增次级分类
			category.ParentId = id
			category.Name = cname
		} else {
			category.Id = id //限制只能有二级分类
		}

	}
	if id != 0 && cname == "" {
		// 3.id不为空，parentid为空，表示不新增分类
		category.Id = id

	}
	article.Time = time.Now()
	article.Category = category
	var articleId int64
	if article.Id == 0 {
		articleId = mArticle.InsertArticle(article)
	} else {
		articleId = mArticle.UpdateArticle(article)
	}

	this.Ctx.Redirect(302, "/blog/article/a/"+strconv.FormatInt(articleId, 10))

}

// @router /blog/delArticle [post]
func (this *BlogController) DelArticle() {
	id, _ := strconv.ParseInt(this.Ctx.Request.Form.Get("id"), 10, 64)
	res := this.Ctx.ResponseWriter
	if mArticle.DelArticleById(id) {
		io.WriteString(res, "true")
	} else {
		io.WriteString(res, "删除失败！")
	}
}

// @router /blog/editArticle/:id [get]
func (this *BlogController) EditArticle() {
	if !this.IsAdmin {
		this.Redirect("/blog/index/1", 302)
	} else {

		categorys := mCategory.GetCategoryAll()
		this.Data["Categorys"] = categorys

		var article = &m.Article{}
		var category = &m.Category{}
		aid, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
		article, _ = mArticle.GetArticleById("a", aid)
		if article != nil {
			this.Data["Article"] = article
			category = article.Category
			if category != nil {
				this.Data["Category"] = category
			}
		}
		this.Data["IsAdmin"] = this.IsAdmin
		this.Layout = "main.html"
		this.TplName = "edit.html"
	}
}

// @router /blog/delCategory [post]
func (this *BlogController) DelCategoryById() {
	id, _ := strconv.ParseInt(this.Ctx.Request.Form.Get("id"), 10, 64)
	res := this.Ctx.ResponseWriter
	if cnt := mArticle.GetArticleByCategoryId(id); cnt != 0 {
		io.WriteString(res, "删除失败，该分类已经在使用，不能删除！")
	} else if cnt := mCategory.GetCategoryByParent(id); cnt != 0 {
		io.WriteString(res, "删除失败，该分类有二级分类，不能删除！")
	} else {
		if mCategory.DelCategoryById(id) {
			io.WriteString(res, "true")
		} else {
			io.WriteString(res, "删除失败！")
		}
	}

}

// @router /blog/updateCategory [post]
func (this *BlogController) UpdateCategory() {
	c := m.Category{}
	err := this.ParseForm(&c)
	category := mCategory.GetCategoryById(c.ParentId)
	childCount := mCategory.GetCategoryChildCount(c.Id)
	if category.ParentId == 0 && childCount == 0 {
		if err == nil {
			mCategory.UpdateCategory(c)
		}
	} else {
		fmt.Println("修改分类失败，存在子分类或者只限二级分类")
	}
	id := strconv.FormatInt(c.Id, 10)
	this.Redirect("/blog/guidang/"+id+"/1", 302)
}

// @router /blog/queryChildCount [post]
func (this *BlogController) QueryChildCount() {
	id, _ := strconv.ParseInt(this.Ctx.Request.Form.Get("id"), 10, 64)
	childCount := mCategory.GetCategoryChildCount(id)
	io.WriteString(this.Ctx.ResponseWriter, strconv.FormatInt(childCount, 10))
}

// @router /blog/loginLogs [get]
func (this *BlogController) LoginLogs() {
	loginLogs := mLoginLog.GetLoginLogAll()
	this.Data["LoginLogs"] = loginLogs
	this.Layout = "main.html"
	this.TplName = "loginlog.html"
}

// @router /blog/clearLoginLog [get]
func (this *BlogController) ClearLoginLog() {
	mLoginLog.ClearLoginLog()
	this.Redirect("/blog/loginLogs/", 302)
}
