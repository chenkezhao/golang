package public

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CouponCategory struct {
	Id          int64     `orm:"column(id);auto"`
	Parent      int64     `orm:"column(parent);null"`
	Name        string    `orm:"column(name);size(50)"`
	Seqno       int       `orm:"column(seqno);null"`
	Status      string    `orm:"column(status);size(1)"`
	Type        string    `orm:"column(type);size(1);null"`
	Iconpath    string    `orm:"column(iconpath);size(200);null"`
	Detailinfo  string    `orm:"column(detailinfo);size(300);null"`
	Createrdate time.Time `orm:"column(createrdate);auto_now_add;type(datetime)"`
	Remark      string    `orm:"column(remark);size(100);null"`
}

func (t *CouponCategory) TableName() string {
	return "coupon_category"
}

func init() {
	orm.RegisterModel(new(CouponCategory))
}

// AddCouponCategory insert a new CouponCategory into database and returns
// last inserted Id on success.
func AddCouponCategory(m *CouponCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponCategoryById retrieves CouponCategory by Id. Returns error if
// Id doesn't exist
func GetCouponCategoryById(id int64) (v *CouponCategory, err error) {
	o := orm.NewOrm()
	v = &CouponCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponCategory retrieves all CouponCategory matches certain condition. Returns empty list if
// no records exist
func GetAllCouponCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponCategory))
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

	var l []CouponCategory
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

// UpdateCouponCategory updates CouponCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponCategoryById(m *CouponCategory) (err error) {
	o := orm.NewOrm()
	v := CouponCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponCategory deletes CouponCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponCategory(id int64) (err error) {
	o := orm.NewOrm()
	v := CouponCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据id查询子行数
func HasChildCount(id int64) int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(CouponCategory)).Filter("parent", id).Count()
	return count
}

// 修改指定的字段
func UpdateByField(id int64, params orm.Params) (err error) {
	o := orm.NewOrm()
	v := CouponCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.QueryTable(new(CouponCategory)).Filter("id", id).Update(params); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//根据标签id集合，查询标签下的所有分类
func QueryCouponCategoryByTagids(tagids string) (categorys []CouponCategory) {
	o := orm.NewOrm()
	sql := "select distinct c.* from coupon_ref_tags rt inner join coupon_info i on i.id = rt.couponid inner join coupon_category c on c.id = i.categoryid where rt.tagid in (" + tagids + ") order by c.seqno"
	if _, err := o.Raw(sql).QueryRows(&categorys); err != nil {
		return nil
	}
	return
}
