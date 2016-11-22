package public

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CouponTag struct {
	Id         int64  `orm:"column(id);auto"`
	Name       string `orm:"column(name);size(30)"`
	Parent     int64  `orm:"column(parent);null"`
	Seqno      int    `orm:"column(seqno);null"`
	Status     string `orm:"column(status);size(1)"`
	Color      string `orm:"column(color);size(30);null"`
	Iconpath   string `orm:"column(iconpath);size(300);null"`
	Detailinfo string `orm:"column(detailinfo);size(300);null"`
	Remark     string `orm:"column(remark);size(100);null"`
}

func (t *CouponTag) TableName() string {
	return "coupon_tag"
}

func init() {
	orm.RegisterModel(new(CouponTag))
}

// AddCouponTag insert a new CouponTag into database and returns
// last inserted Id on success.
func AddCouponTag(m *CouponTag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponTagById retrieves CouponTag by Id. Returns error if
// Id doesn't exist
func GetCouponTagById(id int64) (v *CouponTag, err error) {
	o := orm.NewOrm()
	v = &CouponTag{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponTag retrieves all CouponTag matches certain condition. Returns empty list if
// no records exist
func GetAllCouponTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponTag))
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

	var l []CouponTag
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

// UpdateCouponTag updates CouponTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponTagById(m *CouponTag) (err error) {
	o := orm.NewOrm()
	v := CouponTag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponTag deletes CouponTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponTag(id int64) (err error) {
	o := orm.NewOrm()
	v := CouponTag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponTag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据id查询子行数
func HasChildCountByCouponTag(id int64) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(CouponTag)).Filter("parent", id).Count()
	return count
}

//除了三级，查询所有资源
func QueryNotThreeCouponTagMenu() (v []CouponTag, err error) {
	o := orm.NewOrm()
	var tags []CouponTag
	if _, err := o.Raw("select * from v_query_tag_noThreee").QueryRows(&tags); err == nil {
		return tags, nil
	}
	return nil, err
}

func GetTreeChlidByCouponTag(tagid int64) []orm.Params {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select getTreeChlidByCouponTag(?) as tagids", tagid).Values(&maps)
	if err == nil && num > 0 {
		return maps
	}
	return nil
}
