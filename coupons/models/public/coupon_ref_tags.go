package public

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CouponRefTags struct {
	Id       int64       `orm:"column(id);auto"`
	Couponid *CouponInfo `orm:"column(couponid);rel(fk)"`
	Tagid    *CouponTag  `orm:"column(tagid);rel(fk)"`
	Tagname  string      `orm:"column(tagname);size(30);null"`
}

func (t *CouponRefTags) TableName() string {
	return "coupon_ref_tags"
}

func init() {
	orm.RegisterModel(new(CouponRefTags))
	//beego 支持用户定义模板函数，但是必须在 beego.Run() 调用之前
	beego.AddFuncMap("couponreftagscount", CouponrefTagsCount)
}

// AddCouponRefTags insert a new CouponRefTags into database and returns
// last inserted Id on success.
func AddCouponRefTags(m *CouponRefTags) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponRefTagsById retrieves CouponRefTags by Id. Returns error if
// Id doesn't exist
func GetCouponRefTagsById(id int64) (v *CouponRefTags, err error) {
	o := orm.NewOrm()
	v = &CouponRefTags{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponRefTags retrieves all CouponRefTags matches certain condition. Returns empty list if
// no records exist
func GetAllCouponRefTags(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, related bool) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponRefTags))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	if related {
		qs = qs.RelatedSel()
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

	var l []CouponRefTags
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

// UpdateCouponRefTags updates CouponRefTags by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponRefTagsById(m *CouponRefTags) (err error) {
	o := orm.NewOrm()
	v := CouponRefTags{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponRefTags deletes CouponRefTags by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponRefTags(id int64) (err error) {
	o := orm.NewOrm()
	v := CouponRefTags{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponRefTags{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func InsertMultiCouponRefTags(m []CouponRefTags) (int64, error) {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, m)
	return successNums, err
}

func DeleteCouponRefTagsByField(fields map[string]interface{}) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponRefTags))
	for k, v := range fields {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	if _, err := qs.Delete(); err != nil {
		return err
	}
	return nil
}

func CouponrefTagsCount(v [][]CouponRefTags, rows int64) (count int64) {
	var total int64
	for _, rtss := range v {
		total += int64(len(rtss))
	}
	count = total / rows
	return
}
