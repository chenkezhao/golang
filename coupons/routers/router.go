// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coupons/controllers"
	adminControllers "coupons/controllers/administrator"
	publicControllers "coupons/controllers/public"

	"github.com/astaxie/beego"
)

func init() {
	/*
		==============================================================================Base，通用
	*/
	/*
		基础、通用控制器
	*/
	beego.Include(&controllers.BaseController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.ClientController{})
	/*
		==============================================================================系统管理
	*/
	beego.Include(&adminControllers.AdminController{})
	ns2 := beego.NewNamespace("/admin",
		//此处正式版时改为验证加密请求
		// beego.NSCond(func(ctx *context.Context) bool {
		// 	// if ua := ctx.Input.Request.UserAgent(); ua != "" {
		// 	// 	return true
		// 	// }
		// 	// return false
		// 	return true
		// }),
		beego.NSNamespace("/sysCategory",
			beego.NSInclude(
				&adminControllers.SysCategoryController{},
			),
		),

		beego.NSNamespace("/sysDept",
			beego.NSInclude(
				&adminControllers.SysDeptController{},
			),
		),

		beego.NSNamespace("/sysLoginlog",
			beego.NSInclude(
				&adminControllers.SysLoginlogController{},
			),
		),

		beego.NSNamespace("/sysRes",
			beego.NSInclude(
				&adminControllers.SysResController{},
			),
		),

		beego.NSNamespace("/sysResRole",
			beego.NSInclude(
				&adminControllers.SysResRoleController{},
			),
		),

		beego.NSNamespace("/sysRole",
			beego.NSInclude(
				&adminControllers.SysRoleController{},
			),
		),

		beego.NSNamespace("/sysUnit",
			beego.NSInclude(
				&adminControllers.SysUnitController{},
			),
		),

		beego.NSNamespace("/sysUser",
			beego.NSInclude(
				&adminControllers.SysUserController{},
			),
		),
	)

	/*
		==============================================================================优惠券信息
	*/
	beego.Include(&publicControllers.PublicController{})
	ns3 := beego.NewNamespace("/public",
		beego.NSNamespace("/publicParams",
			beego.NSInclude(
				&publicControllers.PublicParamsController{},
			),
		),
		beego.NSNamespace("/couponCategory",
			beego.NSInclude(
				&publicControllers.CouponCategoryController{},
			),
		),

		beego.NSNamespace("/couponFeedback",
			beego.NSInclude(
				&publicControllers.CouponFeedbackController{},
			),
		),

		beego.NSNamespace("/couponInfo",
			beego.NSInclude(
				&publicControllers.CouponInfoController{},
			),
		),

		beego.NSNamespace("/couponPiclungun",
			beego.NSInclude(
				&publicControllers.CouponPiclungunController{},
			),
		),

		beego.NSNamespace("/couponRefTags",
			beego.NSInclude(
				&publicControllers.CouponRefTagsController{},
			),
		),

		beego.NSNamespace("/couponTag",
			beego.NSInclude(
				&publicControllers.CouponTagController{},
			),
		),
	)
	beego.AddNamespace(ns2, ns3)
}
