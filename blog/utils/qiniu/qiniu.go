package qiniu

import (
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v6/io"
	"github.com/qiniu/api.v6/rs"
)

var (
	UseQiniu       bool
	Cfg            = beego.AppConfig
	QiniuAccessKey string
	QiniuSecretKey string
	QiniuScope     string
)

func UploadFile(localFile string, destName string) (addr string, err error) {
	policy := new(rs.PutPolicy)
	policy.Scope = QiniuScope
	uptoken := policy.Token(nil)

	var ret io.PutRet
	var extra = new(io.PutExtra)
	err = io.PutFile(nil, &ret, uptoken, destName, localFile, extra)
	if err != nil {
		return
	}
	addr = "http://" + QiniuScope + ".qiniudn.com/" + destName
	return
}

func init() {
	QiniuAccessKey = Cfg.String("qiniu_access_key")
	QiniuSecretKey = Cfg.String("qiniu_secret_key")
	QiniuScope = Cfg.String("qiniu_scope")
	UseQiniu = QiniuAccessKey != "" && QiniuSecretKey != "" && QiniuScope != ""
}
