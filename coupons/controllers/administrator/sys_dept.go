package administrator

import (
	models "coupons/models/admin"
	"coupons/utils/convertor"
	"encoding/json"
	"errors"
	. "fmt"
	"strings"
)

// SysDeptController oprations for SysDept
type SysDeptController struct {
	AdminController
}

// URLMapping ...
func (c *SysDeptController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllList", c.GetAllList)
	c.Mapping("Correlate", c.Correlate)
}

// Post ...
// @Title Post
// @Description create SysDept
// @Param	body		body 	models.SysDept	true		"body for SysDept content"
// @Success 201 {int} models.SysDept
// @Failure 403 body is empty
// @router / [post]
func (c *SysDeptController) Post() {

	v := models.SysDept{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		var unitid int64
		c.Ctx.Input.Bind(&unitid, "Unitid")
		unit := &models.SysUnit{}
		unit.Id = unitid
		v.Unitid = unit
		Println("...======================", v.Name, unitid)
		if coun := models.HasChildCountByParentNameSysDept(v.Parent, v.Name); coun > 0 {
			message = "名字不能与父类相同"
		} else {
			v.Creater = c.SysUserInfo.Id
			v.Creatername = c.SysUserInfo.Loginname
			if _, err := models.AddSysDept(&v); err == nil {
				message = "操作成功！"
				success = true
			} else {
				message = err.Error()
			}
		}
	} else {
		message = err.Error()
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()

	// v := models.SysDept{}
	// Println("=====================", v)
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if _, err := models.AddSysDept(&v); err == nil {
	// 		c.Ctx.Output.SetStatus(201)
	// 		c.Data["json"] = v
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
	// Println("...==================对错", c)
}

// GetOne ...
// @Title Get One
// @Description get SysDept by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SysDept
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SysDeptController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v, err := models.GetSysDeptById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SysDept
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SysDept
// @Failure 403
// @router / [get]
func (c *SysDeptController) GetAll() {
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

	l, err := models.GetAllSysDept(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

//查询树形数据网格列表数据
// @router /list [post]
func (c *SysDeptController) GetAllList() {
	page, _ := c.GetInt64("page")
	rows, _ := c.GetInt64("rows")
	id := c.GetString("id")
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

	// var maps []orm.Params
	result := make(map[string]interface{})
	if id == "0" { //父节点查询，并组装json
		//获取子行数
		total := models.HasChildCountBySysDept(0)
		result["total"] = interface{}(total)
		//分页查询数据

		// query := make(map[string]string)
		query["parent"] = "0"
		// o.QueryTable(new(models.SysDept)).Filter("parent", 0).Limit(rows, offset).Values(&maps)

		sysDepts, _ := models.GetAllSysDept(query, nil, nil, nil, offset, rows)

		var maps []map[string]interface{}
		for _, r := range sysDepts {
			//struct 到json str
			mJson, _ := json.Marshal(r)
			//json str 转map
			var m map[string]interface{}
			json.Unmarshal([]byte(string(mJson)), &m)
			maps = append(maps, m)
		}
		//添加状态
		for _, m := range maps {
			// fmt.Printf("______________------dd------>>>>>>>>>>>%T(%v)\n", m["Id"], m["Id"])
			if models.HasChildCountBySysDept(int64(m["Id"].(float64))) > 0 {
				m["state"] = "closed"
			} else {
				m["state"] = "open"
			}
		}
		result["rows"] = maps
		c.Data["json"] = result
	} else { //子节点查询，并组装json
		var items []map[string]interface{}

		// query := make(map[string]string)
		query["parent"] = id

		// o.QueryTable(new(models.SysDept)).Filter("parent", id).Values(&maps)

		sysDepts, _ := models.GetAllSysDept(query, nil, nil, nil, 0, 0)

		var maps []map[string]interface{}
		for _, r := range sysDepts {
			//struct 到json str
			mJson, _ := json.Marshal(r)
			//json str 转map
			var m map[string]interface{}
			json.Unmarshal([]byte(string(mJson)), &m)
			maps = append(maps, m)
		}

		for _, m := range maps {
			if models.HasChildCountBySysDept(int64(m["Id"].(float64))) > 0 {
				m["state"] = "closed"
			} else {
				m["state"] = "open"
			}
			items = append(items, m)
		}
		c.Data["json"] = items
	}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the SysDept
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SysDept	true		"body for SysDept content"
// @Success 200 {object} models.SysDept
// @Failure 403 :id is not int
// @router /update [post]
func (c *SysDeptController) Put() {

	v := models.SysDept{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		old, _ := models.GetSysDeptById(v.Id)
		var unitid int64
		c.Ctx.Input.Bind(&unitid, "Unitid")
		unit := &models.SysUnit{}
		unit.Id = unitid
		v.Unitid = unit
		v.Creater = old.Creater
		v.Creatername = old.Creatername
		v.Createrdate = old.Createrdate
		v.Change = c.SysUserInfo.Id
		v.Changename = c.SysUserInfo.Loginname
		if err := models.UpdateSysDeptById(&v); err == nil {
			message = "操作成功！"
			success = true
		} else {
			message = err.Error()
		}

	} else {
		message = err.Error()
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()

	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := convertor.ToInt64(idStr)
	// v := models.SysDept{Id: id}
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if err := models.UpdateSysDeptById(&v); err == nil {
	// 		c.Data["json"] = "OK"
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// @router /deptCorrelate [post]
func (c *SysDeptController) Correlate() {
	v, err := models.QueryThreeSysUnitMenu()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SysDept
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /del/:id [get]
func (c *SysDeptController) Delete() {

	success := false
	message := "操作失败！"
	id, _ := c.GetInt64(":id")

	SysDept := models.HasChildCountByParentSysDept(id)
	if SysDept > 0 {
		message = "只能从最底子类删"
	} else {
		if err := models.DeleteSysDept(id); err == nil {
			message = "删除成功！"
			success = true
		} else {
			message = err.Error()
		}
	}
	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()

	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := convertor.ToInt64(idStr)
	// if err := models.DeleteSysDept(id); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

//查询所有有效菜单列表
// @router /queryThreeMenu [get]
func (c *SysDeptController) QueryNotThreeMenu() {
	v, err := models.QueryThreeSysDeptMenu()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router /editFrom/:id [get]
func (c *SysDeptController) EditForm() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	sysDept, _ := models.GetSysDeptById(id)
	c.Data["SysDept"] = sysDept
	c.Show("admin/dept/sys_dept_edit.html")
}

// @router /detailFrom/:id [get]
func (c *SysDeptController) DetailFrom() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	sysDept, _ := models.GetSysDeptById(id)
	c.Data["SysDept"] = sysDept
	c.Show("admin/dept/sys_dept_detail.html")
}

// @router /addFrom [get]
func (c *SysDeptController) AddForm() {
	c.Show("admin/dept/sys_dept_add.html")
}
