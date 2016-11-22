package controllers

import (
	models "coupons/models/public"
	"coupons/utils/convertor"
	"strconv"
	"strings"
)

type ClientController struct {
	BaseController
}
type CategoryInfo struct {
	Category models.CouponCategory
	Count    int64
}

/*

	同时大家注意到新版本里面增加了URLMapping这个函数，
	这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
	如果注册了就会通过interface来进行执行函数，性能上面会提升很多

*/
func (c *ClientController) URLMapping() {
	c.Mapping("Home", c.Home)
	c.Mapping("Index", c.Index)
	c.Mapping("Tags", c.Tags)
}

// @router / [get]
func (c *ClientController) Home() {
	c.Redirect("/index", 302)
}

/*
今日上新tab

1.所有分类列表（目前只限一级添加）
2.首页图片轮播（模拟添加属于 今日上新 标签的图片）
3.显示所有标签，每个标签只显示8个宝贝，上新时间排列，
  如果今日所有标签的宝贝都没有上新，则按历史最前顺序显示
4.此tab没有分页功能，只有显示更多功能
*/
// @router /index [get]
func (c *ClientController) Index() {
	var query map[string]string
	var sortby []string
	var order []string
	//分类
	query = make(map[string]string)
	query["status"] = "0"
	sortby = append(sortby, "seqno")
	category, _ := models.GetAllCouponCategory(query, nil, nil, nil, 0, 0)
	//图片轮播
	query = make(map[string]string)
	query["ishome"] = "0"
	query["status"] = "0"
	sortby = append(sortby, "seqno")
	piclungun, _ := models.GetAllCouponPiclungun(query, nil, nil, nil, 0, 0, true)

	//所有一级标签
	query = make(map[string]string)
	query["status"] = "0"
	query["parent"] = "0"
	sortby = append(sortby, "seqno")
	order = append(order, "asc")
	tags, _ := models.GetAllCouponTag(query, nil, nil, nil, 0, 0)
	//券关联标签信息
	var refTags [][]models.CouponRefTags
	for i, _ := range tags {
		var v models.CouponTag
		v = tags[i].(models.CouponTag)
		//指定标签及优惠券信息，限定30条
		query = make(map[string]string)
		var sortby []string
		var order []string
		query["tagid"] = strconv.FormatInt(v.Id, 10)
		sortby = append(sortby, "couponid")
		order = append(order, "desc")
		rts, _ := models.GetAllCouponRefTags(query, nil, sortby, order, 0, 8, true)

		var temp []models.CouponRefTags
		for _, v := range rts {
			temp = append(temp, v.(models.CouponRefTags))
		}
		refTags = append(refTags, temp)
		// copy(refTags, temp)
	}

	c.Data["Category"] = category
	c.Data["Piclungun"] = piclungun
	c.Data["Tags"] = tags
	c.Data["RefTags"] = refTags
	c.Data["tagid"] = "0"
	c.Data["stagid"] = "0"

	c.ClientShow("public/home.html")
}

/*
其它tab

1.根据一级标签，获取图片轮播资源
2.获取所有分类列表
3.加载二级标签
4.根据此一级标签、二级标签加载所有宝贝
5.如果有三级标签，则在宝贝图片块上贴上
6.每个tab只加载30个宝贝，超过以分页加载
*/
// @router /index/:tagid/:stagid/:categoryid/:page [get]
func (c *ClientController) Tags() {
	tagid, _ := c.GetInt64(":tagid")
	stagid, _ := c.GetInt64(":stagid")
	categoryid, _ := c.GetInt64(":categoryid")
	page, _ := c.GetInt64(":page")
	if tagid == 0 { //首页
		c.Redirect("/index", 302)
		return
	}

	//根据标签id(含子标签)得到所有下级的记录，并返回字符串
	var tempTagid int64
	if stagid != 0 {
		tempTagid = stagid
	} else {
		tempTagid = tagid
	}
	tempTids := models.GetTreeChlidByCouponTag(tempTagid)
	tagids := strings.Split(tempTids[0]["tagids"].(string), ",")

	var query map[string]string
	var sortby []string
	var order []string
	rows := int64(50)
	offset := abs(page-1) * rows
	//图片轮播
	query = make(map[string]string)
	query["tagsid"] = c.GetString(":tagid")
	query["status"] = "0"
	sortby = append(sortby, "seqno")
	piclungun, _ := models.GetAllCouponPiclungun(query, nil, nil, nil, 0, 0, true)
	//所有一级标签
	query = make(map[string]string)
	query["status"] = "0"
	query["parent"] = "0"
	sortby = append(sortby, "seqno")
	order = append(order, "asc")
	tags, _ := models.GetAllCouponTag(query, nil, nil, nil, 0, 0)
	//某标签id（含子标签）下的所有分类
	// query = make(map[string]string)
	// query["status"] = "0"
	// sortby = append(sortby, "seqno")
	// category, _ := models.GetAllCouponCategory(query, nil, nil, nil, 0, 0)
	categorys := models.QueryCouponCategoryByTagids(tempTids[0]["tagids"].(string))
	//根据分类及标签id（含子标签），查询每一分类宝贝数量
	var categoryTotal int64
	var categoryInfo []CategoryInfo
	for _, cs := range categorys {
		if count := models.GetCouponCountByCategoryAndTagid(cs.Id, tempTids[0]["tagids"].(string)); count > 0 {
			var temp CategoryInfo
			temp.Category = cs
			temp.Count = count
			categoryInfo = append(categoryInfo, temp)
			categoryTotal += count
		}

	}
	if categoryTotal > 0 {
		tempSlice := append([]CategoryInfo{}, CategoryInfo{Category: models.CouponCategory{Name: "全部优惠"}, Count: categoryTotal})
		categoryInfo = append(tempSlice, categoryInfo...)
	}
	//查询一级标签下对应的二级标签
	query = make(map[string]string)
	query["status"] = "0"
	query["parent"] = strconv.FormatInt(tagid, 10)
	sortby = append(sortby, "seqno")
	order = append(order, "asc")
	secondTag, _ := models.GetAllCouponTag(query, nil, nil, nil, 0, 0)

	//根据标签id，查询所有（含子标签）券关联标签信息
	var couponInfos [][]models.CouponInfo
	total := models.GetCouponInfoCountByTagids(categoryid, tempTids[0]["tagids"].(string))
	for _, tid := range tagids {
		//指定标签及优惠券信息
		query = make(map[string]string)
		var sortby []string
		var order []string
		query["tagid"] = tid
		sortby = append(sortby, "couponid")
		order = append(order, "desc")
		temp := models.GetAllCouponInfoByCAndT(convertor.StringToInt64(tid), categoryid, offset, rows)
		couponInfos = append(couponInfos, temp)
	}
	c.Data["page"] = page
	c.Data["totalPages"] = total / rows
	c.Data["rows"] = rows
	c.Data["tagid"] = strconv.FormatInt(tagid, 10)
	c.Data["stagid"] = strconv.FormatInt(stagid, 10)
	c.Data["categoryid"] = strconv.FormatInt(categoryid, 10)
	c.Data["Piclungun"] = piclungun
	c.Data["Tags"] = tags
	c.Data["CategoryInfo"] = categoryInfo
	c.Data["SecondTag"] = secondTag
	c.Data["CouponInfos"] = couponInfos
	c.ClientShow("public/tags.html")
}

//	Abs(NaN) = NaN
func abs(x int64) int64 {
	// TODO: once golang.org/issue/13095 is fixed, change this to:
	// return Float64frombits(Float64bits(x) &^ (1 << 63))
	// But for now, this generates better code and can also be inlined:
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}
