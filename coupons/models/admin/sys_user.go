package admin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id          int64     `orm:"column(id);auto"`
	Username    string    `orm:"column(username);size(100)"`
	Loginname   string    `orm:"column(loginname);size(100)"`
	Password    string    `orm:"column(password);size(100)"`
	Unitid      *SysUnit  `orm:"column(unitid);rel(fk)"`
	Deptid      *SysDept  `orm:"column(deptid);rel(fk)"`
	Roleid      *SysRole  `orm:"column(roleid);rel(fk)"`
	Creater     int64     `orm:"column(creater)"`
	Creatername string    `orm:"column(creatername);size(100);null"`
	Createrdate time.Time `orm:"column(createrdate);auto_now_add;type(datetime)"`
	Change      int64     `orm:"column(change);null"`
	Changename  string    `orm:"column(changename);size(100);null"`
	Changedate  time.Time `orm:"column(changedate);auto_now;type(datetime);null"`
	Remark      string    `orm:"column(remark);size(100);null"`
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(SysUser))
}

// AddSysUser insert a new SysUser into database and returns
// last inserted Id on success.
func AddSysUser(m *SysUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysUserById retrieves SysUser by Id. Returns error if
// Id doesn't exist
func GetSysUserById(id int64) (v *SysUser, err error) {
	o := orm.NewOrm()
	v = &SysUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysUser retrieves all SysUser matches certain condition. Returns empty list if
// no records exist
func GetAllSysUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysUser))
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

	var l []SysUser
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

// UpdateSysUser updates SysUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysUserById(m *SysUser) (err error) {
	o := orm.NewOrm()
	v := SysUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysUser deletes SysUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysUser(id int64) (err error) {
	o := orm.NewOrm()
	v := SysUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
