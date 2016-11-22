package public

import (
	models "coupons/models/public"
	"coupons/utils/convertor"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// CouponRefTagsController oprations for CouponRefTags
type CouponRefTagsController struct {
	PublicController
}

// URLMapping ...
func (c *CouponRefTagsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("QueryRefTagsByiid", c.QueryRefTagsByiid)
	c.Mapping("DelBytagidAndCouponid", c.DelBytagidAndCouponid)
}

// Post ...
// @Title Post
// @Description create CouponRefTags
// @Param	body		body 	models.CouponRefTags	true		"body for CouponRefTags content"
// @Success 201 {int} models.CouponRefTags
// @Failure 403 body is empty
// @router / [post]
func (c *CouponRefTagsController) Post() {
	success := false
	message := "操作失败！"

	var couponid, tagid int64
	c.Ctx.Input.Bind(&couponid, "couponid")
	c.Ctx.Input.Bind(&tagid, "tagid")
	coupon, _ := models.GetCouponInfoById(couponid)
	tag, _ := models.GetCouponTagById(tagid)
	v := models.CouponRefTags{Couponid: coupon, Tagid: tag, Tagname: tag.Name}

	query := make(map[string]string)
	query["couponid"] = strconv.FormatInt(couponid, 10)
	query["tagid"] = strconv.FormatInt(tagid, 10)
	if temp, _ := models.GetAllCouponRefTags(query, nil, nil, nil, 0, 0, false); temp == nil {
		if _, err := models.AddCouponRefTags(&v); err == nil {
			message = "操作成功！"
			success = true
		} else {
			message = err.Error()
		}
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get CouponRefTags by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CouponRefTags
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CouponRefTagsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v, err := models.GetCouponRefTagsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get CouponRefTags
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CouponRefTags
// @Failure 403
// @router / [get]
func (c *CouponRefTagsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCouponRefTags(query, fields, sortby, order, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the CouponRefTags
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CouponRefTags	true		"body for CouponRefTags content"
// @Success 200 {object} models.CouponRefTags
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CouponRefTagsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v := models.CouponRefTags{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCouponRefTagsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the CouponRefTags
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CouponRefTagsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	if err := models.DeleteCouponRefTags(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//根据优惠券信息id查询数据
// @router /queryRefTagsByiid/:iid [get]
func (c *CouponRefTagsController) QueryRefTagsByiid() {
	iid := c.GetString(":iid")
	query := make(map[string]string)
	query["couponid"] = iid
	l, err := models.GetAllCouponRefTags(query, nil, nil, nil, 0, 0, true)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @router /delBytagidAndCouponid/:tagid/:couponid [get]
func (c *CouponRefTagsController) DelBytagidAndCouponid() {
	success := false
	message := "操作失败！"
	tagid, _ := c.GetInt64(":tagid")
	couponid, _ := c.GetInt64(":couponid")
	field := make(map[string]interface{})
	field["tagid"] = tagid
	field["couponid"] = couponid
	if err := models.DeleteCouponRefTagsByField(field); err == nil {
		message = "删除成功！"
		success = true
	} else {
		message = err.Error()
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}
