package admin

import (
	m "blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//插入单个对象
func Insert(admin m.SysAdmin) int64 {
	id, err := m.GetRom().Insert(&admin)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

//插入多个对象
func InsertMulti(admins []m.SysAdmin) int64 {
	successNums, err := m.GetRom().InsertMulti(100, admins)
	if err != nil {
		fmt.Println(err)
	}
	return successNums
}

//修改
func Update(admin m.SysAdmin) int64 {
	if m.GetRom().Read(&admin) == nil {
		num, err := m.GetRom().Update(&admin)
		if err != nil {
			fmt.Println(err)
		}
		return num
	}
	return 0
}

//删除
func Delete(id int64) int64 {
	num, err := m.GetRom().Delete(&m.SysAdmin{Id: 1})
	if err != nil {
		fmt.Println(err)
	}
	return num
}

//查询
func GetById(id int64) m.SysAdmin {
	admin := m.SysAdmin{Id: id}

	err := m.GetRom().Read(&admin)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		//fmt.Println(user.Id, user.Name)
	}
	return admin
}

//查询数据行数
func GetAdminCount() int64 {
	cnt, err := m.GetRom().QueryTable("t_sys_admin").Count()
	if err != nil {
		cnt = 1
		fmt.Println(err)
	}
	return cnt
}

//登录名和密码登录
func Login(admin m.SysAdmin) bool {
	b := false
	count, err := m.GetRom().QueryTable(new(m.SysAdmin)).Filter("LoginName", admin.LoginName).Filter("Password", admin.Password).Count()
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		//fmt.Println(user.Id, user.Name)
		if count != 0 {
			b = true
		}
	}
	return b
}
