package main

import (
	"blog/Initialization"
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	Initialization.InitBlog()
	beego.Run()
}
