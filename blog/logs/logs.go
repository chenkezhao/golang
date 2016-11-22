package logs

import (
	"github.com/astaxie/beego/logs"
)

/*配置多文件日志 multifile*/
// filename 保存的文件名
// maxlines 每个文件保存的最大行数，默认值 1000000
// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
// daily 是否按照每天 logrotate，默认是 true
// maxdays 文件最多保存多少天，默认保存 7 天
// rotate 是否开启 logrotate，默认是 true
// level 日志保存的时候的级别，默认是 Trace 级别
// perm 日志文件权限
// separate 需要单独写入文件的日志级别,设置后命名类似test.error.log
func init() {
	log := logs.NewLogger()
	logs.Async()
	//设置输出
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"blog.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	//输出文件名和行号
	beego.SetLogFuncCall(true)

	/*示例：*/
	// beego.Emergency("this is emergency")
	// beego.Alert("this is alert")
	// beego.Critical("this is critical")
	// beego.Error("this is error")
	// beego.Warning("this is warning")
	// beego.Notice("this is notice")
	// beego.Informational("this is informational")
	// beego.Debug("this is debug")
}
