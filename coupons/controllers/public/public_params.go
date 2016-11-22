package public

import (
	"compress/gzip"
	models "coupons/models/public"
	"coupons/utils/convertor"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/httplib"
	"io"
	"io/ioutil"
	"strings"
)

// PublicParamsController oprations for PublicParams
type PublicParamsController struct {
	PublicController
}

// URLMapping ...
func (c *PublicParamsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create PublicParams
// @Param	body		body 	models.PublicParams	true		"body for PublicParams content"
// @Success 201 {int} models.PublicParams
// @Failure 403 body is empty
// @router / [post]
func (c *PublicParamsController) Post() {
	var v models.PublicParams
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPublicParams(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get PublicParams by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PublicParams
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PublicParamsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v, err := models.GetPublicParamsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get PublicParams
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.PublicParams
// @Failure 403
// @router / [get]
func (c *PublicParamsController) GetAll() {
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

	l, err := models.GetAllPublicParams(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the PublicParams
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.PublicParams	true		"body for PublicParams content"
// @Success 200 {object} models.PublicParams
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PublicParamsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v := models.PublicParams{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdatePublicParamsById(&v); err == nil {
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
// @Description delete the PublicParams
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PublicParamsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	if err := models.DeletePublicParams(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router /xx [get]
func (c *PublicParamsController) Xx() {
	//优惠券页面
	// req := httplib.Get("https://shop.m.taobao.com/shop/coupon.htm?seller_id=2048320167&activityId=52e1cbfaf09d4c299c4f915908287770")
	//下单页面
	// req := httplib.Get("https://detail.tmall.com/item.htm?id=541525424957&ali_trackid=2:mm_114203826_12892178_48986794:1479300782_255_709447294&pvid=10_111.2.132.109_49812_1479276820265&sku_properties=20509:782040598")
	req := httplib.Get("https://item.taobao.com/item.htm?id=528249033903&ali_trackid=2:mm_114203826_12892178_48986794:1479304491_213_1490521465&pvid=10_111.2.132.109_466_1479269115118")
	// req.Debug(true).Response()

	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //https请求

	req.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header("Accept-Encoding", "gzip, deflate, sdch, br")
	req.Header("accept-language", "zh-CN,zh;q=0.8")
	req.Header("Host", "shop.m.taobao.com")
	req.Header("Host", "shop.m.taobao.com")
	req.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.87 Safari/537.36")

	resp, _ := req.Response()
	var reader io.ReadCloser
	var err error
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
	} else {
		reader = resp.Body
	}

	bt, _ := ioutil.ReadAll(reader)
	c.Data["json"] = string(bt)
	// c.Data["json"] = resp.Header

	// c.Data["json"] = resp.Header
	c.ServeJSON()
}
