package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponCategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponFeedbackController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponPiclungunController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponRefTagsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:CouponsTagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"] = append(beego.GlobalControllerRouter["coupons/controllers/public:PublicParamsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
