package models

//注：这个包名下，只存放的是类的结构体

import ()

// Created time.Time `orm:"auto_now_add;type(datetime)"`
// Updated time.Time `orm:"auto_now;type(datetime)"`
// auto_now 每次 model 保存时都会对时间自动更新
// auto_now_add 第一次保存时才设置时间
// 对于批量的 update 此设置是不生效的
