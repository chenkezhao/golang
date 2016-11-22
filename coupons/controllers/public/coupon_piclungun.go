package public

import (
	models "coupons/models/public"
	"coupons/utils/convertor"
	"coupons/utils/filetool"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// CouponPiclungunController oprations for CouponPiclungun
type CouponPiclungunController struct {
	PublicController
	UploadPath string
}

func (c *CouponPiclungunController) InitUploadPath() {
	c.UploadPath = c.UploadBasePath + "coupon\\piclungun\\"
}

// URLMapping ...
func (c *CouponPiclungunController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllList", c.GetAllList)
	c.Mapping("EditForm", c.EditForm)
	c.Mapping("DetailFrom", c.DetailFrom)
	c.Mapping("AddForm", c.AddForm)
}

// Post ...
// @Title Post
// @Description create CouponPiclungun
// @Param	body		body 	models.CouponPiclungun	true		"body for CouponPiclungun content"
// @Success 201 {int} models.CouponPiclungun
// @Failure 403 body is empty
// @router / [post]
func (c *CouponPiclungunController) Post() {
	v := models.CouponPiclungun{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		var tagsid int64
		c.Ctx.Input.Bind(&tagsid, "tagsid")
		tag := &models.CouponTag{}
		tag.Id = tagsid
		v.Tagsid = tag

		if id, err := models.AddCouponPiclungun(&v); err == nil {
			message = "操作成功！"
			success = true

			/*如果有图片，保存图片*/
			f, h, err := c.GetFile("file")
			if f != nil {
				defer f.Close()
				if err != nil {
					message = err.Error()
				} else {
					// 保存位置在 static/upload,没有文件夹要先创建
					if err := os.MkdirAll(c.UploadPath, 0777); err == nil {
						filename := strconv.FormatInt(id, 10) + filetool.Ext(h.Filename)
						if err := c.SaveToFile("file", c.UploadPath+filename); err != nil {
							message = "保存上传文件失败！" + err.Error()
							success = false
						} else {
							obj, _ := models.GetCouponPiclungunById(id)
							obj.Picpath = filename
							if _, err := models.UpdateCouponPiclungunByField(obj, "picpath"); err != nil {
								message = err.Error()
								success = false
							}
						}

					} else {
						message = "创建上传文件目录失败！" + err.Error()
						success = false
					}
				}
			}
		} else {
			message = err.Error()
		}
	} else {
		message = err.Error()
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get CouponPiclungun by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CouponPiclungun
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CouponPiclungunController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v, err := models.GetCouponPiclungunById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get CouponPiclungun
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CouponPiclungun
// @Failure 403
// @router / [get]
func (c *CouponPiclungunController) GetAll() {
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

	l, err := models.GetAllCouponPiclungun(query, fields, sortby, order, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the CouponPiclungun
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CouponPiclungun	true		"body for CouponPiclungun content"
// @Success 200 {object} models.CouponPiclungun
// @Failure 403 :id is not int
// @router /update [post]
func (c *CouponPiclungunController) Put() {
	v := models.CouponPiclungun{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		var tagsid int64
		c.Ctx.Input.Bind(&tagsid, "tagsid")
		tag := &models.CouponTag{}
		tag.Id = tagsid
		v.Tagsid = tag
		temp, _ := models.GetCouponPiclungunById(v.Id)
		if err := models.UpdateCouponPiclungunById(&v); err == nil {
			message = "操作成功！"
			success = true
			/*如果有图片，保存图片*/
			f, h, err := c.GetFile("file")
			if f != nil {
				defer f.Close()
				if err != nil {
					message = err.Error()
				} else {
					// 保存位置在 static/upload,没有文件夹要先创建
					if err := os.MkdirAll(c.UploadPath, 0777); err == nil {
						filename := strconv.FormatInt(v.Id, 10) + filetool.Ext(h.Filename)
						if err := c.SaveToFile("file", c.UploadPath+filename); err != nil {
							message = "保存上传文件失败！" + err.Error()
							success = false
						} else {
							obj, _ := models.GetCouponPiclungunById(v.Id)
							obj.Picpath = filename
							if _, err := models.UpdateCouponPiclungunByField(obj, "picpath"); err != nil {
								message = err.Error()
								success = false
							}
						}

					} else {
						message = "创建上传文件目录失败！" + err.Error()
						success = false
					}
				}
			} else {
				if strings.Replace(v.Picpath, " ", "", -1) != temp.Picpath {
					file := c.UploadPath + temp.Picpath
					if filetool.IsExist(file) {
						if err := filetool.Unlink(file); err != nil {
							message = "删除已有上传图片失败！"
							success = true
						}
					}
				}
			}
		} else {
			message = err.Error()
		}

	} else {
		message = err.Error()
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the CouponPiclungun
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /del/:id [get]
func (c *CouponPiclungunController) Delete() {
	success := false
	message := ""
	id, _ := c.GetInt64(":id")
	//删除上传的图片
	obj, _ := models.GetCouponPiclungunById(id)
	file := c.UploadPath + obj.Picpath
	if filetool.IsExist(file) {
		if err := filetool.Unlink(file); err != nil {
			message = "删除已上传图片失败！"
			success = true
		}
	}
	if err := models.DeleteCouponPiclungun(id); err == nil {
		message = "删除成功！" + message
		success = true
	} else {
		message = err.Error()
		success = false
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}

//查询数据网格列表数据
// @router /list [post]
func (c *CouponPiclungunController) GetAllList() {
	page, _ := c.GetInt64("page")
	rows, _ := c.GetInt64("rows")
	offset := (page - 1) * rows

	query := make(map[string]string)
	sDate := c.GetString("sDate")
	eDate := c.GetString("eDate")
	if sDate != "" {
		query["createrdate__gt"] = sDate
	}
	if eDate != "" {
		query["createrdate__lte"] = eDate
	}
	query["name__icontains"] = c.GetString("name")

	result := make(map[string]interface{})
	//获取总行数
	total := models.CountByCouponPiclungun()
	result["total"] = interface{}(total)
	//分页查询数据
	objs, _ := models.GetAllCouponPiclungun(query, nil, nil, nil, offset, rows, true)

	var maps []map[string]interface{}
	for _, r := range objs {
		//struct 到json str
		mJson, _ := json.Marshal(r)
		//json str 转map
		var m map[string]interface{}
		json.Unmarshal([]byte(string(mJson)), &m)
		maps = append(maps, m)
	}
	if maps == nil {
		result["rows"] = models.CouponPiclungun{}
	} else {
		result["rows"] = maps
	}
	c.Data["json"] = result
	c.ServeJSON()

}

// @router /editFrom/:id [get]
func (c *CouponPiclungunController) EditForm() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	v, _ := models.GetCouponPiclungunById(id)
	c.Data["Obj"] = v
	uploadpath := strings.Replace(c.UploadPath, "\\", "/", -1)
	c.Data["uploadpath"] = uploadpath + v.Picpath
	c.Show("admin/coupon/piclungun/coupon_piclungun_edit.html")
}

// @router /detailFrom/:id [get]
func (c *CouponPiclungunController) DetailFrom() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	Obj, _ := models.GetCouponPiclungunById(id)
	c.Data["Obj"] = Obj
	c.Show("admin/coupon/piclungun/coupon_piclungun_detail.html")
}

// @router /addFrom [get]
func (c *CouponPiclungunController) AddForm() {
	c.Show("admin/coupon/piclungun/coupon_piclungun_add.html")
}