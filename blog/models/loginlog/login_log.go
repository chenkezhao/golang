package loginlog

import (
	m "blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//插入单个对象
func Insert(loginLog m.LoginLog) int64 {
	id, err := m.GetRom().Insert(&loginLog)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

//插入多个对象
func InsertMulti(loginLogs []m.LoginLog) int64 {
	successNums, err := m.GetRom().InsertMulti(100, loginLogs)
	if err != nil {
		fmt.Println(err)
	}
	return successNums
}

//修改
func Update(loginLog m.LoginLog) int64 {
	if m.GetRom().Read(&loginLog) == nil {
		num, err := m.GetRom().Update(&loginLog)
		if err != nil {
			fmt.Println(err)
		}
		return num
	}
	return 0
}

//删除
func Delete(id int64) int64 {
	num, err := m.GetRom().Delete(&m.LoginLog{Id: 1})
	if err != nil {
		fmt.Println(err)
	}
	return num
}

//查询
func GetById(id int64) m.LoginLog {
	loginLog := m.LoginLog{Id: id}

	err := m.GetRom().Read(&loginLog)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		//fmt.Println(user.Id, user.Name)
	}
	return loginLog
}

/*获取所有*/
func GetLoginLogAll() []*m.LoginLog {
	var loginLogs []*m.LoginLog
	_, err := m.GetRom().QueryTable(new(m.LoginLog)).OrderBy("time").All(&loginLogs)
	if err != nil {
		fmt.Println(err)
	}
	return loginLogs
}

//清除所有记录
func ClearLoginLog() {
	m.GetRom().Raw("delete from t_login_log where  id != 0").Exec()
}
