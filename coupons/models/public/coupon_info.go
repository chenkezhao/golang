package public

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type CouponInfo struct {
	Id              int64           `orm:"column(id);auto"`
	Categoryid      *CouponCategory `orm:"column(categoryid);rel(fk)"`
	Picpaths        string          `orm:"column(picpaths);size(8000);null"`
	Title           string          `orm:"column(title);size(300)"`
	Subtitle        string          `orm:"column(subtitle);size(300);null"`
	Seqno           int             `orm:"column(seqno);null"`
	Couponslink     string          `orm:"column(couponslink);size(500)"`
	Buylink         string          `orm:"column(buylink);size(500)"`
	Status          string          `orm:"column(status);size(1)"`
	Tags            string          `orm:"column(tags);size(500);null"`
	Detailinfo      string          `orm:"column(detailinfo);size(500);null"`
	Reason          string          `orm:"column(reason);size(300);null"`
	Createrdate     time.Time       `orm:"column(createrdate);auto_now_add;type(datetime)"`
	Changedate      time.Time       `orm:"column(changedate);auto_now;type(datetime);null"`
	Remark          string          `orm:"column(remark);size(100);null"`
	Extend1         string          `orm:"column(extend1);size(100);null"`
	Extend2         string          `orm:"column(extend2);size(100);null"`
	Extend3         string          `orm:"column(extend3);size(100);null"`
	Price           float64         `orm:"column(price)"`
	Couponprice     float64         `orm:"column(couponprice)"`
	Daylimit        int             `orm:"column(daylimit);null"`
	Couponcount     int             `orm:"column(couponcount);null"`
	Couponusedcount int             `orm:"column(couponusedcount);null"`
	Buycount        int             `orm:"column(buycount);null"`
}

func (t *CouponInfo) TableName() string {
	return "coupon_info"
}

func init() {
	orm.RegisterModel(new(CouponInfo))
}

// AddCouponInfo insert a new CouponInfo into database and returns
// last inserted Id on success.
func AddCouponInfo(m *CouponInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCouponInfoById retrieves CouponInfo by Id. Returns error if
// Id doesn't exist
func GetCouponInfoById(id int64) (v *CouponInfo, err error) {
	o := orm.NewOrm()
	v = &CouponInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCouponInfo retrieves all CouponInfo matches certain condition. Returns empty list if
// no records exist
func GetAllCouponInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, related bool) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CouponInfo))
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

	var l []CouponInfo
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

// UpdateCouponInfo updates CouponInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateCouponInfoById(m *CouponInfo) (err error) {
	o := orm.NewOrm()
	v := CouponInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCouponInfo deletes CouponInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCouponInfo(id int64) (err error) {
	o := orm.NewOrm()
	v := CouponInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CouponInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据id查询子行数
func CountByCouponInfo() int64 {
	o := orm.NewOrm()
	count, _ := o.QueryTable(new(CouponInfo)).Count()
	return count
}

//修改指定字段
func UpdateCouponInfoByField(cp *CouponInfo, field string) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(cp, field)
	if err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func GetCouponCountByCategoryAndTagid(categoryid int64, tagids string) (count int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := "select count(*) count from coupon_info i inner join coupon_ref_tags rt on rt.couponid = i.id where i.categoryid = ? and rt.tagid in (" + tagids + ")"
	if num, err := o.Raw(sql, categoryid).Values(&maps); err == nil && num > 0 {
		count, _ = strconv.ParseInt(maps[0]["count"].(string), 10, 64)
		return
	}
	return 0
}

func GetAllCouponInfoByCAndT(tagid, categoryid, offset, rows int64) (ml []CouponInfo) {
	var sql bytes.Buffer
	sql.WriteString(" select i.*  from coupon_ref_tags rt  ")
	sql.WriteString(" inner join coupon_tag t on t.id=rt.tagid  ")
	sql.WriteString(" inner join coupon_info i on i.id=rt.couponid  ")
	sql.WriteString(" inner join coupon_category c on c.id=i.categoryid  ")
	sql.WriteString(" where 1=1  ")
	sql.WriteString(" and rt.tagid = ?  ")
	if categoryid != 0 {
		sql.WriteString(" and c.id =" + strconv.FormatInt(categoryid, 10))
	}
	sql.WriteString(" order by rt.couponid desc  ")
	sql.WriteString(" LIMIT ? OFFSET ?  ")
	o := orm.NewOrm()
	o.Raw(sql.String(), tagid, rows, offset).QueryRows(&ml)
	return
}

func GetCouponInfoCountByTagids(categoryid int64, tagids string) (count int64) {

	var sql bytes.Buffer
	sql.WriteString(" select count(*) count  ")
	sql.WriteString(" from coupon_ref_tags rt  ")
	sql.WriteString(" inner join coupon_info i on i.id=rt.couponid  ")
	sql.WriteString(" inner join coupon_category c on c.id=i.categoryid  ")
	sql.WriteString(" where 1=1  ")
	sql.WriteString(" and rt.tagid in(" + tagids + ")  ")
	if categoryid != 0 {
		sql.WriteString(" and c.id = " + strconv.FormatInt(categoryid, 10) + "  ")
	}

	var maps []orm.Params
	o := orm.NewOrm()
	if num, err := o.Raw(sql.String()).Values(&maps); err == nil && num > 0 {
		count, _ = strconv.ParseInt(maps[0]["count"].(string), 10, 64)
		return
	}
	return 0
}
