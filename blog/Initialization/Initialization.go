package Initialization

import (
	m "blog/models"
	mAdmin "blog/models/admin"
	"blog/utils/strtool"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var Cfg = beego.AppConfig
var RunMode string

func InitBlog() {
	// -----------------------------------start---------------database
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"

	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:@/blog?charset=utf8")

	// register model
	// orm.RegisterModel( new(LoginLog), new(Article), new(Category))
	//注册表并添加表的前缀名
	orm.RegisterModelWithPrefix("t_", new(m.LoginLog), new(m.Article), new(m.Category), new(m.SysAdmin))

	// create table
	orm.RunSyncdb("default", false, true)
	RunMode = Cfg.String("runmode")
	if RunMode == "dev" {
		//控制台打印sql语句
		orm.Debug = true
	}
	// -----------------------------------end---------------database

	// -----------------------------------start---------------初始化数据
	count := mAdmin.GetAdminCount()
	if count == 0 {
		generateData()
	} else {
		fmt.Println("无初始化数据...")
	}
	// -----------------------------------end---------------初始化数据
}

func generateData() {
	admin := m.SysAdmin{
		UserName:  "陈科肇",
		LoginName: "admin",
		Password:  strtool.Md5("000000"),
	}
	mAdmin.Insert(admin)
	defer func() {
		fmt.Println("始化数据成功...")
	}()
}
