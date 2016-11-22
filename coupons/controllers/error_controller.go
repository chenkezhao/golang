package controllers

import ()

/*
通过上面的例子我们可以看到，所有的函数都是有一定规律的，都是Error开头，后面的名字就是我们调用Abort的名字，例如Error404函数其实调用对应的就是Abort("404")

我们就只要在beego.Run之前采用beego.ErrorController注册这个错误处理函数就可以了
package main

import (
    _ "btest/routers"
    "btest/controllers"

    "github.com/astaxie/beego"
)

func main() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Run()
}
*/
type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
