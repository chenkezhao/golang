package administrator

import (
	models "coupons/models/admin"
	"coupons/utils/convertor"
	"encoding/json"
	"errors"
	. "fmt"
	"strings"
)

// SysUnitController oprations for SysUnit
type SysUnitController struct {
	AdminController
}

// URLMapping ...
func (c *SysUnitController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("AddForm", c.AddForm)
	c.Mapping("QueryNotThreeMenu", c.QueryNotThreeMenu)
}

// Post ...
// @Title Post
// @Description create SysUnit
// @Param	body		body 	models.SysUnit	true		"body for SysUnit content"
// @Success 201 {int} models.SysUnit
// @Failure 403 body is empty
// @router / [post]
func (c *SysUnitController) Post() {

	v := models.SysUnit{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		if coun := models.HasChildCountByParentNameSysUnit(v.Parent, v.Name); coun > 0 {
			message = "名字不能与父类相同"
		} else {
			v.Creater = c.SysUserInfo.Id
			v.Creatername = c.SysUserInfo.Loginname
			if _, err := models.AddSysUnit(&v); err == nil {
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

	// var v models.SysUnit{}
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if _, err := models.AddSysUnit(&v); err == nil {
	// 		c.Ctx.Output.SetStatus(201)
	// 		c.Data["json"] = v
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SysUnit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SysUnit
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SysUnitController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := convertor.ToInt64(idStr)
	v, err := models.GetSysUnitById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SysUnit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SysUnit
// @Failure 403
// @router / [get]
func (c *SysUnitController) GetAll() {
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

	l, err := models.GetAllSysUnit(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

//查询树形数据网格列表数据
// @router /list [post]
func (c *SysUnitController) GetAllList() {
	page, _ := c.GetInt64("page")
	rows, _ := c.GetInt64("rows")
	id := c.GetString("id")
	offset := (page - 1) * rows
	Println("...========================进来可以了")
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
		total := models.HasChildCountBySysUnit(0)
		result["total"] = interface{}(total)
		//分页查询数据

		// query := make(map[string]string)
		query["parent"] = "0"
		// o.QueryTable(new(models.SysUnit)).Filter("parent", 0).Limit(rows, offset).Values(&maps)

		sysUnits, _ := models.GetAllSysUnit(query, nil, nil, nil, offset, rows)

		var maps []map[string]interface{}
		for _, r := range sysUnits {
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
			if models.HasChildCountBySysUnit(int64(m["Id"].(float64))) > 0 {
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

		// o.QueryTable(new(models.SysUnit)).Filter("parent", id).Values(&maps)

		sysUnits, _ := models.GetAllSysUnit(query, nil, nil, nil, 0, 0)

		var maps []map[string]interface{}
		for _, r := range sysUnits {
			//struct 到json str
			mJson, _ := json.Marshal(r)
			//json str 转map
			var m map[string]interface{}
			json.Unmarshal([]byte(string(mJson)), &m)
			maps = append(maps, m)
		}

		for _, m := range maps {
			if models.HasChildCountBySysUnit(int64(m["Id"].(float64))) > 0 {
				m["state"] = "closed"
			} else {
				m["state"] = "open"
			}
			items = append(items, m)
		}
		c.Data["json"] = items
	}
	Println("...------------", c)
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the SysUnit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SysUnit	true		"body for SysUnit content"
// @Success 200 {object} models.SysUnit
// @Failure 403 :id is not int
// @router /update [post]
func (c *SysUnitController) Put() {
	v := models.SysUnit{}
	success := false
	message := "操作失败！"
	if err := c.ParseForm(&v); err == nil {
		old, _ := models.GetSysUnitById(v.Id)
		v.Creater = old.Creater
		v.Creatername = old.Creatername
		v.Createrdate = old.Createrdate
		v.Change = c.SysUserInfo.Id
		v.Changename = c.SysUserInfo.Loginname
		if err := models.UpdateSysUnitById(&v); err == nil {
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
}

// Delete ...
// @Title Delete
// @Description delete the SysUnit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /del/:id [get]
func (c *SysUnitController) Delete() {

	success := false
	message := "操作失败！"
	id, _ := c.GetInt64(":id")

	SysUnit := models.HasChildCountByParentSysUnit(id)
	if SysUnit > 0 {
		message = "只能从最底子类删"
	} else {
		if err := models.DeleteSysUnit(id); err == nil {
			message = "删除成功！"
			success = true
		} else {
			message = err.Error()
		}
	}

	c.Data["json"] = c.ResponseData(success, message)
	c.ServeJSON()
}

//查询所有有效菜单列表
// @router /queryThreeMenu [get]
func (c *SysUnitController) QueryNotThreeMenu() {
	v, err := models.QueryThreeSysUnitMenu()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router /editFrom/:id [get]
func (c *SysUnitController) EditForm() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	SysUnit, _ := models.GetSysUnitById(id)
	c.Data["SysUnit"] = SysUnit
	c.Show("admin/unit/sys_unit_edit.html")
}

// @router /detailFrom/:id [get]
func (c *SysUnitController) DetailFrom() {
	id, _ := c.GetInt64(":id")
	// id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	SysUnit, _ := models.GetSysUnitById(id)
	c.Data["SysUnit"] = SysUnit
	c.Show("admin/unit/sys_unit_detail.html")
}

// @router /addFrom [get]
func (c *SysUnitController) AddForm() {
	c.Show("admin/unit/sys_unit_add.html")
}
