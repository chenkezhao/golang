package category

import (
	m "blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetCategoryByParentId(parentid int64) *m.Category {
	category := &m.Category{ParentId: parentid}
	err := m.GetRom().Read(category)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
	}
	return category
}
func GetCategoryById(id int64) *m.Category {
	category := &m.Category{Id: id}
	err := m.GetRom().Read(category)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
	}
	return category
}

/*获取所有分类，树形菜单*/
func GetCategoryAll() []*m.Category {
	var categorys []*m.Category
	var temp []*m.Category
	_, err := m.GetRom().QueryTable("t_category").OrderBy("parent_id", "id").All(&temp)
	if err == nil {
		for _, t := range temp {
			if t.ParentId != 0 {
				break
			}
			categorys = append(categorys, t)
			for _, c := range temp {
				if t.Id == c.ParentId {
					categorys = append(categorys, c)
				}
			}
		}
	}
	return categorys
}

func GetCategoryByParent(cid int64) int64 {
	cnt, _ := m.GetRom().QueryTable("t_category").Filter("parent_id", cid).Count()
	return cnt
}

func DelCategoryById(id int64) bool {
	if _, err := m.GetRom().Delete(&m.Category{Id: id}); err == nil {
		return true
	}
	return false
}
func UpdateCategory(category m.Category) {
	_, err := m.GetRom().Update(&category)
	commitChekErr(err)
}

func GetCategoryChildCount(id int64) int64 {
	cnt, _ := m.GetRom().QueryTable("t_category").Filter("parent_id", id).Count()
	return cnt
}

func commitChekErr(err error) {
	if err == nil {
		m.GetRom().Commit()
		fmt.Println("提交。。。")
	} else {
		m.GetRom().Rollback()
		fmt.Println("回滚。。。", err)
	}
}
