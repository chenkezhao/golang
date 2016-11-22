package main

import (
	"coupons/initialization"
	_ "coupons/routers"
	"github.com/astaxie/beego"
	"math"
)

func main() {
	initialization.InitData()

	//beego 支持用户定义模板函数，但是必须在 beego.Run() 调用之前
	beego.AddFuncMap("subtraction", Subtraction)
	beego.Run()
}

func Subtraction(a float64, b float64) (out float64) {
	out = math.Abs(a - b)
	return
}
