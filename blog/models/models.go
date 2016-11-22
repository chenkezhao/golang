package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Created time.Time `orm:"auto_now_add;type(datetime)"`
// Updated time.Time `orm:"auto_now;type(datetime)"`
// auto_now 每次 model 保存时都会对时间自动更新
// auto_now_add 第一次保存时才设置时间
// 对于批量的 update 此设置是不生效的

/*文章分类*/
type Category struct {
	Id       int64 `PK`
	Name     string
	ParentId int64
	Article  []*Article `orm:"reverse(many)"`
}

/*文章*/
type Article struct {
	Id       int64 `PK`
	Title    string
	Summary  string    `orm:"size(500)"`
	Content  string    `orm:"type(text)"`
	Time     time.Time `orm:"auto_now_add;type(datetime)"`
	Category *Category `orm:"rel(fk)"`
}

/*系统管理员*/
type SysAdmin struct {
	Id        int64 `pk`
	UserName  string
	LoginName string
	Password  string
}

/*管理员登录日志*/
type LoginLog struct {
	Id        int64 `PK`
	LoginName string
	Ip        string
	Address   string
	Time      time.Time `orm:"auto_now_add;type(datetime)"`
	Token     string
}

func GetRom() orm.Ormer {
	return orm.NewOrm()
}
