package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["coupons/controllers/administrator:AdminController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:AdminController"],
		beego.ControllerComments{
			Method: "ToMain",
			Router: `/admin/main`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:AdminController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:AdminController"],
		beego.ControllerComments{
			Method: "Exit",
			Router: `/admin/exit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysCategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "GetAllList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "Correlate",
			Router: `/deptCorrelate`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/del/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "QueryNotThreeMenu",
			Router: `/queryThreeMenu`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "EditForm",
			Router: `/editFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "DetailFrom",
			Router: `/detailFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysDeptController"],
		beego.ControllerComments{
			Method: "AddForm",
			Router: `/addFrom`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysLoginlogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/del/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "QueryNotThreeMenu",
			Router: `/queryNotThreeMenu`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "QueryAllMenu",
			Router: `/queryAllMenu`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "GetAllList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "EditForm",
			Router: `/editFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "DetailFrom",
			Router: `/detailFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResController"],
		beego.ControllerComments{
			Method: "AddForm",
			Router: `/addFrom`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysResRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "GetAllList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/del/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "QueryNotThreeMenu",
			Router: `/queryThreeMenu`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "EditForm",
			Router: `/editFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "DetailFrom",
			Router: `/detailFrom/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUnitController"],
		beego.ControllerComments{
			Method: "AddForm",
			Router: `/addFrom`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"] = append(beego.GlobalControllerRouter["coupons/controllers/administrator:SysUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
