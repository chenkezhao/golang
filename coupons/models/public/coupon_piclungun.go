package public

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CouponPiclungun struct {
	Id         int64      `orm:"column(id);auto"`
	Tagsid     *CouponTag `orm:"column(tagsid);rel(fk)"`
	Name       string     `orm:"column(name);size(100)"`
	Title      string     `orm:"column(title);size(200);null"`
	Seqno      int        `orm:"column(seqno);null"`
	Ishome     string     `orm:"column(ishome);size(10);default(1)"`
	Status     string     `orm:"column(status);size(10)"`
	Picpath    string     `orm:"column(picpath);size(300);null"`
	Detailinfo string     `orm:"column(detailinfo);size(500);null"`
	Remark     string     `orm:"column(remark);size(100);null"`
	Extend1    string     `orm:"column(extend1);size(100);null"`
	Extend2    string     `orm:"column(extend2);size(100);null"`
	Extend3    string     `orm:"column(extend3);size(100);null"`
}

func (t *CouponPiclungun) TableName() string {
	return "coupon_piclungun"
}

func init() {
	orm.RegisterModel(new(CouponPiclungun))
}

// AddCouponPiclungun insert a new CouponPiclungun into database and returns
// last inserted Id on success.
func AddCouponPiclungun(m *CouponPiclungun) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponPiclungunById retrieves CouponPiclungun by Id. Returns error if
// Id doesn't exist
func GetCouponPiclungunById(id int64) (v *CouponPiclungun, err error) {
	o := orm.NewOrm()
	v = &CouponPiclungun{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponPiclungun retrieves all CouponPiclungun matches certain condition. Returns empty list if
// no records exist
func GetAllCouponPiclungun(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, related bool) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponPiclungun))
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

	var l []CouponPiclungun
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

// UpdateCouponPiclungun updates CouponPiclungun by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponPiclungunById(m *CouponPiclungun) (err error) {
	o := orm.NewOrm()
	v := CouponPiclungun{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponPiclungun deletes CouponPiclungun by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponPiclungun(id int64) (err error) {
	o := orm.NewOrm()
	v := CouponPiclungun{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponPiclungun{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据id查询子行数
func CountByCouponPiclungun() int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(CouponPiclungun)).Count()
	return count
}

//修改指定字段
func UpdateCouponPiclungunByField(cp *CouponPiclungun, field string) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(cp, field)
	if err == nil {
		return num, nil
	} else {
		return 0, err
	}
}
