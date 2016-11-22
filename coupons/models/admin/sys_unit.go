package admin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysUnit struct {
	Id          int64     `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(100)"`
	Code        string    `orm:"column(code);size(30)"`
	Parent      int64     `orm:"column(parent);null"`
	Call        string    `orm:"column(call);size(200);null"`
	Email       string    `orm:"column(email);size(50);null"`
	Address     string    `orm:"column(address);size(300);null"`
	Creater     int64     `orm:"column(creater)"`
	Creatername string    `orm:"column(creatername);size(100);null"`
	Createrdate time.Time `orm:"column(createrdate);auto_now_add;type(datetime)"`
	Change      int64     `orm:"column(change);null"`
	Changename  string    `orm:"column(changename);size(100);null"`
	Changedate  time.Time `orm:"column(changedate);auto_now;type(datetime);null"`
	Remark      string    `orm:"column(remark);size(100);null"`
}

func (t *SysUnit) TableName() string {
	return "sys_unit"
}

func init() {
	orm.RegisterModel(new(SysUnit))
}

// AddSysUnit insert a new SysUnit into database and returns
// last inserted Id on success.
func AddSysUnit(m *SysUnit) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysUnitById retrieves SysUnit by Id. Returns error if
// Id doesn't exist
func GetSysUnitById(id int64) (v *SysUnit, err error) {
	o := orm.NewOrm()
	v = &SysUnit{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//外键相关查询
func GetSysUnitParentById(id int64) (v *SysUnit, err error) {
	o := orm.NewOrm()
	v = &SysUnit{Parent: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysUnit retrieves all SysUnit matches certain condition. Returns empty list if
// no records exist
func GetAllSysUnit(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysUnit))
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

	var l []SysUnit
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

// UpdateSysUnit updates SysUnit by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysUnitById(m *SysUnit) (err error) {
	o := orm.NewOrm()
	v := SysUnit{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysUnit deletes SysUnit by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysUnit(id int64) (err error) {
	o := orm.NewOrm()
	v := SysUnit{Id: id}
	// m2m := o.QueryM2M(v, "parent")
	// tags := SysUnit{Parent: id}
	// num, errs := m2m.Remove(tags)
	// if errs == nil {
	// 	fmt.Println("...===================num", num)
	// } else {
	// 	fmt.Println("=====================err", errs)
	// }

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysUnit{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//查询数据行数
func GetSysUnitCount() int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(SysUnit)).Count()
	if err != nil {
		cnt = 1
		fmt.Println(err)
	}
	return cnt
}

//查询所有资源
func QueryThreeSysUnitMenu() (v []SysUnit, err error) {
	o := orm.NewOrm()
	var unit []SysUnit
	if _, err := o.Raw("select * from sys_unit").QueryRows(&unit); err == nil {
		return unit, nil
	}
	return v, err
}

//根据id查询子行数
func HasChildCountBySysUnit(id int64) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(SysUnit)).Filter("parent", id).Count()
	return count
}

//根据id查询行数
func HasChildCountByParentSysUnit(id int64) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(SysUnit)).Filter("parent", id).Count()
	return count
}

//根据id,name查询行数
func HasChildCountByParentNameSysUnit(id int64, name string) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(SysUnit)).Filter("id", id).Filter("name", name).Count()
	return count
}
