package initialization

import (
	// m "coupons/models"
	mAdmin "coupons/models/admin"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var (
	Cfg     = beego.AppConfig
	RunMode string
)

func init() {
	// -----------------------------------start---------------database
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"

	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	// register model
	// orm.RegisterModel( new(LoginLog), new(Article), new(Category))
	//注册表并添加表的前缀名
	// orm.RegisterModelWithPrefix("t_", new(m.LoginLog), new(m.Article), new(m.Category), new(m.SysAdmin))
	// create table，自动生成数据结构表
	// orm.RunSyncdb("default", false, true)
	RunMode = Cfg.String("runmode")
	if RunMode == "dev" {
		//控制台打印sql语句
		orm.Debug = true
	}
	// -----------------------------------end---------------database
}

func InitData() {

	// -----------------------------------start---------------初始化数据
	count := mAdmin.GetSysUnitCount()
	if count == 0 {
		initAdminData()
	} else {
		fmt.Println("无初始化数据...")
	}
	// -----------------------------------end---------------初始化数据
}

var o orm.Ormer

func initAdminData() {
	var (
		unitId, deptId, roleId, resId1, resId2, resId3 int64
		err                                            error
	)
	sysUnit := &mAdmin.SysUnit{
		Name:        "系统默认机构",
		Code:        "S001",
		Creater:     -1,
		Creatername: "初始化",
	}
	sysDept := &mAdmin.SysDept{
		Name:        "默认部门",
		Category:    -1,
		Unitid:      &mAdmin.SysUnit{},
		Creater:     -1,
		Creatername: "初始化",
	}
	sysRole := &mAdmin.SysRole{
		Name:   "超级管理员",
		Code:   "R001",
		Remark: "初始化",
	}
	sysRes1 := &mAdmin.SysRes{
		Name:        "系统管理",
		Creater:     -1,
		Creatername: "初始化",
		Seqno:       1,
	}
	sysRes2 := &mAdmin.SysRes{
		Name:        "系统资源",
		Parent:      resId1,
		Creater:     -1,
		Creatername: "初始化",
		Seqno:       1,
	}
	sysRes3 := &mAdmin.SysRes{
		Name:        "系统资源管理",
		Parent:      resId2,
		Respath:     "admin/res/sys_res.html",
		Creater:     -1,
		Creatername: "初始化",
		Seqno:       1,
	}
	sysResRole := &mAdmin.SysResRole{
		Roleid: &mAdmin.SysRole{},
		Resid:  &mAdmin.SysRes{},
	}

	sysUser := &mAdmin.SysUser{
		Username:    "超级管理员",
		Loginname:   Cfg.String("root_name"),
		Password:    Cfg.String("root_pass"),
		Unitid:      &mAdmin.SysUnit{},
		Deptid:      &mAdmin.SysDept{},
		Roleid:      &mAdmin.SysRole{},
		Creater:     -1,
		Creatername: "初始化",
	}

	o = orm.NewOrm()
	txerr := o.Begin()
	defer func() {
		if txerr != nil {
			o.Rollback()
			fmt.Println("............>>>>>>>>>>>>>>>>>始化数据失败...数据回滚。", txerr)
			return
		} else {
			o.Commit()
			fmt.Println("............>>>>>>>>>>>>>>>>>始化数据成功...数据提交。")
		}
	}()

	unitId, err = mAdmin.AddSysUnit(sysUnit)

	if err == nil {
		sysDept.Unitid.Id = unitId
		deptId, err = mAdmin.AddSysDept(sysDept)

	}

	roleId, err = mAdmin.AddSysRole(sysRole)

	resId1, err = mAdmin.AddSysRes(sysRes1)
	if err == nil {
		sysRes2.Parent = resId1
		resId2, err = mAdmin.AddSysRes(sysRes2)
		if err == nil {
			sysRes3.Parent = resId2
			resId3, err = mAdmin.AddSysRes(sysRes3)

		}
	}
	sysResRole.Resid.Id = resId3
	sysResRole.Roleid.Id = roleId
	_, err = mAdmin.AddSysResRole(sysResRole)
	sysUser.Unitid.Id = unitId
	sysUser.Deptid.Id = deptId
	sysUser.Roleid.Id = roleId
	_, err = mAdmin.AddSysUser(sysUser)
	fmt.Println("-----------------------------------准备提交数据--------------")

}
