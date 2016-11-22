package admin

import (
	// "database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysRes struct {
	Id          int64     `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(100)"`
	Parent      int64     `orm:"column(parent);default(0)"`
	Respath     string    `orm:"column(respath);size(100);null"`
	Seqno       int       `orm:"column(seqno);null"`
	Creater     int64     `orm:"column(creater)"`
	Creatername string    `orm:"column(creatername);size(100);null"`
	Createrdate time.Time `orm:"column(createrdate);auto_now_add;type(datetime)"`
	Change      int64     `orm:"column(change);null"`
	Changename  string    `orm:"column(changename);size(100);null"`
	Changedate  time.Time `orm:"column(changedate);auto_now;type(datetime);null"`
	Remark      string    `orm:"column(remark);size(100);null"`
}

func (t *SysRes) TableName() string {
	return "sys_res"
}

func init() {
	orm.RegisterModel(new(SysRes))
}

// AddSysRes insert a new SysRes into database and returns
// last inserted Id on success.
func AddSysRes(m *SysRes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysResById retrieves SysRes by Id. Returns error if
// Id doesn't exist
func GetSysResById(id int64) (v *SysRes, err error) {
	o := orm.NewOrm()
	v = &SysRes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysRes retrieves all SysRes matches certain condition. Returns empty list if
// no records exist
func GetAllSysRes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysRes))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SysRes
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSysRes updates SysRes by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysResById(m *SysRes) (err error) {
	o := orm.NewOrm()
	v := SysRes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysRes deletes SysRes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysRes(id int64) (err error) {
	o := orm.NewOrm()
	v := SysRes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysRes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//除了三级，查询所有资源
func QueryNotThreeSysResMenu() (v []SysRes, err error) {
	o := orm.NewOrm()
	var ress []SysRes
	if _, err := o.Raw("select * from v_query_tree_noThreee").QueryRows(&ress); err == nil {
		return ress, nil
	}
	return nil, err
}

//查询所有有效资源，即含一、二、三级的完整树
func QueryAllSysResMenu() (v []SysRes, err error) {
	o := orm.NewOrm()
	var ress []SysRes
	if _, err := o.Raw("select * from v_query_all_menu").QueryRows(&ress); err == nil {
		return ress, nil
	}
	return nil, err
}

//根据id查询子行数
func HasChildCountBySysRes(id int64) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(SysRes)).Filter("parent", id).Count()
	return count
}
